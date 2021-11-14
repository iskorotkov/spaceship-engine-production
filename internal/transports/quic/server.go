package quic

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/gob"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"

	"github.com/lucas-clemente/quic-go"
)

type Handler func(req interface{}) (Response, error)

type Server struct {
	handlers map[Type]Handler
}

func (s Server) Start(addr string) error {
	tlsConfig, err := s.generateTLSConfig()
	if err != nil {
		return err
	}

	listener, err := quic.ListenAddr(addr, tlsConfig, nil)
	if err != nil {
		return err
	}
	defer listener.Close()

	log.Printf("listener created")

	for {
		session, err := listener.Accept(context.Background())
		if err != nil {
			return err
		}

		log.Printf("session created")

		go func() {
			stream, err := session.AcceptStream(context.Background())
			if err != nil {
				log.Print(err)
			}

			log.Printf("stream created")

			if err := s.handleConnection(stream); err != nil {
				log.Print(err)
			}
		}()
	}
}

func (s *Server) Handle(t Type, h Handler) {
	if s.handlers == nil {
		s.handlers = make(map[string]Handler)
	}

	s.handlers[t] = h
}

func (s Server) handleConnection(stream quic.Stream) error {
	encoder := gob.NewEncoder(stream)
	decoder := gob.NewDecoder(stream)

	for {
		var req Request
		if err := decoder.Decode(&req); err != nil {
			return err
		}

		log.Printf("request decoded")

		h, ok := s.handlers[req.Type]
		if !ok {
			return fmt.Errorf("no handler set for %q", req.Type)
		}

		resp, err := h(req.Message)
		if err != nil {
			return err
		}

		if err := encoder.Encode(resp); err != nil {
			return err
		}

		log.Printf("response sent")
	}
}

func (s Server) generateTLSConfig() (*tls.Config, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{SerialNumber: big.NewInt(1)}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		return nil, err
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"x-quic"},
	}, nil
}

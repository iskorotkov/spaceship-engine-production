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
	"math/big"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/lucas-clemente/quic-go"
)

var _ = (transport.Server)(&Server{})

type Server struct {
	handlers map[transport.Type]transport.Handler
}

func (s *Server) Start(addr string) error {
	tlsConfig, err := s.generateTLSConfig()
	if err != nil {
		return fmt.Errorf("error generating tls config: %w", err)
	}

	listener, err := quic.ListenAddr(addr, tlsConfig, nil)
	if err != nil {
		return fmt.Errorf("error starting listener: %w", err)
	}

	log.Printf("listener created")

	go func() {
		defer listener.Close()

		for {
			session, err := listener.Accept(context.Background())
			if err != nil {
				log.Printf("error accepting session: %v", err)
				return
			}

			log.Printf("session created")

			go func() {
				stream, err := session.AcceptStream(context.Background())
				if err != nil {
					log.Printf("error accepting stream: %v", err)
					return
				}

				log.Printf("stream created")

				if err := s.handleConnection(stream); err != nil {
					log.Printf("error handling connection: %v", err)
					session.CloseWithError(1, err.Error())
					return
				}
			}()
		}
	}()

	return nil
}

func (s *Server) Handle(t transport.Type, h transport.Handler) error {
	if s.handlers == nil {
		s.handlers = make(map[string]transport.Handler)
	}

	s.handlers[t] = h

	log.Printf("added quic handler for %v", t)

	return nil
}

func (s *Server) Close() error {
	// Closing currently not supported.
	return nil
}

func (s *Server) handleConnection(stream quic.Stream) error {
	defer stream.Close()

	encoder := gob.NewEncoder(stream)
	decoder := gob.NewDecoder(stream)

	for {
		var req transport.Request
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

		log.Printf("sent response %v", resp.Message)
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

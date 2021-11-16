package quic

import (
	"context"
	"crypto/tls"
	"encoding/gob"
	"fmt"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/lucas-clemente/quic-go"
)

var _ = (transport.Server)(&Server{})

type Server struct {
	handlers map[transport.Type]transport.Handler
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(addr string, config *tls.Config) error {
	listener, err := quic.ListenAddr(addr, config, nil)
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

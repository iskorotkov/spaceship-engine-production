package tcp

import (
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
)

var _ = (transport.Server)(&Server{})

type Server struct {
	handlers map[transport.Type]transport.Handler
	listener net.Listener
}

func (s *Server) Start(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error starting listener: %w", err)
	}

	s.listener = listener

	log.Printf("listener created")

	go func() {
		defer listener.Close()

		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("error accepting connection: %v", err)
				return
			}

			log.Printf("connection created")

			go func() {
				if err := s.handleConnection(conn); err != nil {
					log.Printf("error handling connection: %v", err)
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

	log.Printf("added tcp handler for %v", t)

	return nil
}

func (s *Server) Close() error {
	return s.listener.Close()
}

func (s *Server) handleConnection(conn net.Conn) error {
	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	for {
		var req transport.Request
		if err := decoder.Decode(&req); errors.Is(err, io.EOF) {
			return nil
		} else if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("request decoded")

		h, ok := s.handlers[req.Type]
		if !ok {
			return fmt.Errorf("no handler set for %q", req.Type)
		}

		log.Printf("used handler for %q", req.Type)

		resp, err := h(req.Message)
		if err != nil {
			if err := encoder.Encode(transport.Response{Message: err.Error()}); err != nil {
				return fmt.Errorf("error sending error response: %w", err)
			}

			return err
		}

		if err := encoder.Encode(resp); err != nil && strings.Contains(err.Error(), "connection reset by peer") {
			log.Printf("client closed connection")
			return nil
		} else if err != nil {
			log.Print(err)
			continue
		}

		log.Printf("sent response %v", resp.Message)
	}
}

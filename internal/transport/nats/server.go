package nats

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/iskorotkov/spaceship-engine-production/internal/transport"
	"github.com/nats-io/nats.go"
)

var _ = (transport.Server)(&Server{})

type Server struct {
	nats *nats.Conn
}

func (s *Server) Start(addr string) error {
	nc, err := nats.Connect(addr)
	if err != nil {
		return fmt.Errorf("error creating nats connection: %w", err)
	}

	log.Printf("nats connection created")

	s.nats = nc

	return nil
}

func (s *Server) Handle(t transport.Type, h transport.Handler) error {
	_, err := s.nats.Subscribe(t, func(msg *nats.Msg) {
		log.Printf("received nats event from subject %q", msg.Subject)

		decoder := gob.NewDecoder(bytes.NewReader(msg.Data))

		var req transport.Request
		if err := decoder.Decode(&req); err != nil {
			log.Printf("error decoding request: %v", err)
			return
		}

		resp, err := h(req.Message)
		if err != nil {
			log.Printf("handler returned error: %v", err)
			return
		}

		var b bytes.Buffer
		encoder := gob.NewEncoder(&b)

		if err := encoder.Encode(resp); err != nil {
			log.Printf("error encoding response: %v", err)
			return
		}

		if err := msg.Respond(b.Bytes()); err != nil {
			log.Printf("error sending response: %v", err)
			return
		}

		log.Printf("sent response %v", resp.Message)
	})
	if err != nil {
		return fmt.Errorf("error subscribing to subject: %w", err)
	}

	if err := s.nats.Flush(); err != nil {
		return fmt.Errorf("error performing nats flush: %w", err)
	}

	log.Printf("nats server subscribed to %q", t)

	return nil
}

func (s *Server) Close() error {
	if err := s.nats.Drain(); err != nil {
		return fmt.Errorf("error draining nats connection: %w", err)
	}

	log.Printf("nats server stopped")

	return nil
}

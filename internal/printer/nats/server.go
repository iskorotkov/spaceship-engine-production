package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/iskorotkov/spaceship-engine-production/api"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer"
	"github.com/nats-io/nats.go"
)

var ErrNoSubscription = fmt.Errorf("error stopping printer server: subscription isn't created")

type Server struct {
	client       *nats.Conn
	subscription *nats.Subscription
	service      *printer.Printer
}

func NewServer(url string) (Server, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return Server{}, fmt.Errorf("error creating nats client: %w", err)
	}

	service := printer.New()

	return Server{
		client:       nc,
		subscription: nil,
		service:      service,
	}, nil
}

func (s *Server) Subscribe(ctx context.Context) error {
	subscription, err := s.client.Subscribe(SubjectPrintRequests, func(msg *nats.Msg) {
		var req api.PrintRequest

		if err := json.Unmarshal(msg.Data, &req); err != nil {
			log.Printf("error unmarshaling message data: %v", err)

			return
		}

		if _, err := s.service.Print(ctx, &req); err != nil {
			log.Printf("error printing request: %v", err)

			return
		}
	})
	if err != nil {
		return fmt.Errorf("error subscribing to subject: %w", err)
	}

	s.subscription = subscription

	if err := s.client.FlushWithContext(ctx); err != nil {
		return fmt.Errorf("error performing nats flush: %w", err)
	}

	return nil
}

func (s Server) Unsubscribe() error {
	if s.subscription == nil {
		return ErrNoSubscription
	}

	if err := s.subscription.Drain(); err != nil {
		return fmt.Errorf("error draining subscription: %w", err)
	}

	return nil
}

func (s Server) Close() error {
	if err := s.client.Drain(); err != nil {
		return fmt.Errorf("error draining printer server: %w", err)
	}

	return nil
}

package nats

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/iskorotkov/spaceship-engine-production/internal/printer"
	"github.com/nats-io/nats.go"
)

var _ = (printer.Client)(Client{}) //nolint:exhaustivestruct

type Client struct {
	client *nats.Conn
}

func NewClient(url string) (Client, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return Client{}, fmt.Errorf("error creating nats client: %w", err)
	}

	log.Printf("nats client created")

	return Client{
		client: nc,
	}, nil
}

func (c Client) Print(ctx context.Context, req printer.Request) error {
	b, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("error marshaling request to json: %w", err)
	}

	if err := c.client.Publish(SubjectPrintRequests, b); err != nil {
		return fmt.Errorf("error publishing request: %w", err)
	}

	if err := c.client.FlushWithContext(ctx); err != nil {
		return fmt.Errorf("error performing nats flush: %w", err)
	}

	log.Printf("nats client sent request")

	return nil
}

func (c Client) Close() error {
	c.client.Close()

	log.Printf("nats client closed")

	return nil
}

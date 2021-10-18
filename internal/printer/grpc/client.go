package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/iskorotkov/spaceship-engine-production/api"
	"github.com/iskorotkov/spaceship-engine-production/internal/printer"
	"google.golang.org/grpc"
)

var _ = (printer.Client)(Client{}) //nolint:exhaustivestruct

type Client struct {
	conn   *grpc.ClientConn
	client api.ReportPrinterClient
}

func NewClient(url string) (Client, error) {
	conn, err := grpc.Dial(url,
		grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error creating grpc connection: %v", err)
	}

	client := api.NewReportPrinterClient(conn)

	return Client{
		conn:   conn,
		client: client,
	}, nil
}

func (c Client) Print(ctx context.Context, req printer.Request) error {
	if _, err := c.client.Print(ctx, req); err != nil {
		return fmt.Errorf("error processing print request: %w", err)
	}

	return nil
}

func (c Client) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("error closing grpc connection: %w", err)
	}

	return nil
}

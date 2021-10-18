package printer

import (
	"context"

	"github.com/iskorotkov/spaceship-engine-production/api"
)

type Request *api.PrintRequest

type Client interface {
	Print(ctx context.Context, req Request) error
	Close() error
}

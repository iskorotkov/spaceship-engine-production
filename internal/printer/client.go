package printer

import (
	"context"

	"github.com/iskorotkov/spaceship-engine-production/api/report-printer"
)

type Request *report_printer.PrintRequest

type Client interface {
	Print(ctx context.Context, req Request) error
	Close() error
}

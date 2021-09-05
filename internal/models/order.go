package models

import (
	"time"
)

type Order struct {
	ID          uint
	Amount      int
	CreatedAt   time.Time
	CompletedAt *time.Time

	Client   *Client
	ClientID uint

	Engine   *Engine
	EngineID uint
}

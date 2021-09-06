package models

import (
	"time"
)

type Row struct {
	ID uint
	RowClient
	RowComponent
	RowEngine
	RowOrder
	RowProvider

	// Requirement: M:M, no additional fields.
	// Supply: M:M, no additional fields.
}

type RowClient struct {
	ClientID   uint
	ClientName string
}

type RowComponent struct {
	ComponentID   uint
	ComponentName string
}

type RowEngine struct {
	EngineID    uint
	EngineName  string
	EnginePower float64
}

type RowOrder struct {
	OrderID          uint
	OrderAmount      int
	OrderCreatedAt   time.Time
	OrderCompletedAt *time.Time
}

type RowProvider struct {
	ProviderID   uint
	ProviderName string
}

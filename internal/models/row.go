package models

import "time"

type Row struct {
	ID uint

	// Client.
	ClientID   uint
	ClientName string

	// Component.
	ComponentID   uint
	ComponentName string

	// Engine.
	EngineID    uint
	EngineName  string
	EnginePower float64

	// Order.
	OrderID          uint
	OrderAmount      int
	OrderCreatedAt   time.Time
	OrderCompletedAt *time.Time

	// Provider.
	ProviderID   uint
	ProviderName string

	// Requirement: M:M, no additional fields.
	// Supply: M:M, no additional fields.
}

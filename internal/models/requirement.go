package models

type Requirement struct {
	ID uint

	Engine   *Engine
	EngineID uint

	Component   *Component
	ComponentID uint
}

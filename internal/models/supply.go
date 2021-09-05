package models

type Supply struct {
	ID uint

	Component   *Component
	ComponentID uint

	Provider   *Provider
	ProviderID uint
}

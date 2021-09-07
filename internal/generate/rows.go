package generate

import (
	"math/rand"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
)

const maxOrderAmount = 20

func Rows(clients []models.Client, engines []models.Engine, components []models.Component, providers []models.Provider, ordersPerClient, componentsPerEngine, providersPerComponent Range) []models.Row {
	engineComponents := generateEngineComponents(engines, components, componentsPerEngine)
	componentProviders := generateComponentProviders(components, providers, providersPerComponent)

	var (
		rows    []models.Row
		orderID uint
	)

	for _, client := range clients {
		for i := 0; i < ordersPerClient.Random(); i++ {
			createdAt, completedAt := timeRange()
			engine := engines[rand.Intn(len(engines))] //nolint:gosec
			amount := rand.Intn(maxOrderAmount)        //nolint:gosec
			orderID++

			for _, component := range engineComponents[engine] {
				for _, provider := range componentProviders[component] {
					item := models.Row{
						ID: uint(len(rows) + 1),
						RowClient: models.RowClient{
							ClientID:   client.ID,
							ClientName: client.Name,
						},
						RowComponent: models.RowComponent{
							ComponentID:   component.ID,
							ComponentName: component.Name,
						},
						RowEngine: models.RowEngine{
							EngineID:    engine.ID,
							EngineName:  engine.Name,
							EnginePower: engine.Power,
						},
						RowOrder: models.RowOrder{
							OrderID:          orderID,
							OrderAmount:      amount,
							OrderCreatedAt:   createdAt,
							OrderCompletedAt: &completedAt,
						},
						RowProvider: models.RowProvider{
							ProviderID:   provider.ID,
							ProviderName: provider.Name,
						},
					}

					rows = append(rows, item)
				}
			}
		}
	}

	return rows
}

func generateComponentProviders(components []models.Component, providers []models.Provider, providersPerComponent Range) map[models.Component][]models.Provider {
	componentProviders := make(map[models.Component][]models.Provider)

	for _, component := range components {
		for i := 0; i < providersPerComponent.Random(); i++ {
			provider := providers[rand.Intn(len(providers))] //nolint:gosec
			componentProviders[component] = append(componentProviders[component], provider)
		}
	}

	return componentProviders
}

func generateEngineComponents(engines []models.Engine, components []models.Component, componentsPerEngine Range) map[models.Engine][]models.Component {
	engineComponents := make(map[models.Engine][]models.Component)

	for _, engine := range engines {
		for i := 0; i < componentsPerEngine.Random(); i++ {
			component := components[rand.Intn(len(components))] //nolint:gosec
			engineComponents[engine] = append(engineComponents[engine], component)
		}
	}

	return engineComponents
}

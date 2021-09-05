package generate

import (
	"math/rand"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
)

const maxOrderAmount = 20

func Rows(
	clients []models.Client,
	engines []models.Engine,
	components []models.Component,
	providers []models.Provider,
	clientOrders, engineComponents, componentProviders int,
) []models.Row {
	var rows []models.Row

	for _, client := range clients {
		for i := 0; i < clientOrders; i++ {
			engine := engines[rand.Intn(len(engines))] //nolint:gosec

			for j := 0; j < engineComponents; j++ {
				component := components[rand.Intn(len(components))] //nolint:gosec

				for k := 0; k < componentProviders; k++ {
					provider := providers[rand.Intn(len(providers))] //nolint:gosec

					createdAt, completedAt := timeRange()

					item := models.Row{
						ID: uint(len(rows) + 1),

						// Client.
						ClientID:   client.ID,
						ClientName: client.Name,

						// Component.
						ComponentID:   component.ID,
						ComponentName: component.Name,

						// Engine.
						EngineID:    engine.ID,
						EngineName:  engine.Name,
						EnginePower: engine.Power,

						// Order.
						OrderID:          uint(len(rows) + 1),
						OrderAmount:      rand.Intn(maxOrderAmount), //nolint:gosec
						OrderCreatedAt:   createdAt,
						OrderCompletedAt: &completedAt,

						// Provider.
						ProviderID:   provider.ID,
						ProviderName: provider.Name,
					}

					rows = append(rows, item)
				}
			}
		}
	}

	return rows
}

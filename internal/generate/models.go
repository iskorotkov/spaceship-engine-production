package generate

import (
	"fmt"
	"math/rand"

	"github.com/iskorotkov/spaceship-engine-production/internal/models"
)

const (
	enginePowerMultiplier = 10
	maxEnginePowerFactor  = 100
)

func Components(adjectives, nouns []string, count int) []models.Component {
	items := make([]models.Component, 0, count)

	for i, name := range names([][]string{adjectives, nouns}, count) {
		items = append(items, models.Component{
			ID:   uint(i + 1),
			Name: name,
		})
	}

	return items
}

func Engines(types, generations, series []string, count int) []models.Engine {
	items := make([]models.Engine, 0, count)

	for i := 0; i < count; i++ {
		t := types[rand.Intn(len(types))]               //nolint:gosec
		gen := generations[rand.Intn(len(generations))] //nolint:gosec
		series := series[rand.Intn(len(series))]        //nolint:gosec

		items = append(items, models.Engine{
			ID:    uint(i + 1),
			Name:  fmt.Sprintf("%s Engine Mk %s%s", t, gen, series),
			Power: float64(enginePowerMultiplier * rand.Intn(maxEnginePowerFactor)), //nolint:gosec
		})
	}

	return items
}

func Clients(adjectives, nouns []string, count int) []models.Client {
	items := make([]models.Client, 0, count)

	for i, name := range names([][]string{adjectives, adjectives, nouns}, count) {
		items = append(items, models.Client{
			ID:   uint(i + 1),
			Name: name,
		})
	}

	return items
}

func Providers(adjectives, nouns []string, count int) []models.Provider {
	items := make([]models.Provider, 0, count)

	for i, name := range names([][]string{adjectives, adjectives, nouns}, count) {
		items = append(items, models.Provider{
			ID:   uint(i + 1),
			Name: name,
		})
	}

	return items
}

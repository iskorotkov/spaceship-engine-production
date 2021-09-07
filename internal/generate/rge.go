package generate

import (
	"math/rand"
)

type Range struct {
	Min, Max int
}

func (r Range) Random() int {
	min, max := r.Min, r.Max
	return min + rand.Intn(max-min) //nolint:gosec
}

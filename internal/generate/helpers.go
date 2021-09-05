package generate

import (
	"math/rand"
	"strings"
	"time"
)

const fiveYears = 365 * 5

func timeRange() (time.Time, time.Time) {
	start := time.Now().AddDate(0, 0, -rand.Intn(fiveYears)) //nolint:gosec
	end := start.AddDate(0, 0, rand.Intn(fiveYears))         //nolint:gosec

	return start, end
}

func names(parts [][]string, count int) []string {
	names := make(map[string]struct{})

	for len(names) != count {
		var words []string

		for _, part := range parts {
			word := part[rand.Intn(len(part))] //nolint:gosec

			for len(words) != 0 && words[len(words)-1] == word {
				word = part[rand.Intn(len(part))] //nolint:gosec
			}

			words = append(words, word)
		}

		name := strings.Join(words, " ")
		names[name] = struct{}{}
	}

	items := make([]string, 0, len(names))
	for name := range names {
		items = append(items, name)
	}

	return items
}

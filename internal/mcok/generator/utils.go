package generator

import (
	"math/rand"
	"time"

	"github.com/mroth/weightedrand/v2"
)

// RandomIntBetween [min, max)
func RandomIntBetween(min int, max int) int {
	randomSeed := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return randomSeed.Intn(max-min) + min
}

// Something return different option from [availables] other than [currentOption]
//
// This is use to diversify the output by returning an unexpected output with (100-percentage)% chance
func MixWithBias[T comparable](currentOption T, availables []T, percentage int) T {
	currentOptionWeight := 10_000 - percentage*100
	otherOptionWeight := (10_000 - currentOptionWeight) / (len(availables) - 1)

	choices := make([]weightedrand.Choice[T, int], 0)
	for _, item := range availables {
		if item == currentOption {
			choices = append(choices, weightedrand.NewChoice(item, currentOptionWeight))
		} else {
			choices = append(choices, weightedrand.NewChoice(item, otherOptionWeight))
		}
	}

	chooser, _ := weightedrand.NewChooser(choices...)
	return chooser.Pick()
}

// Pick one random item from source
func RandomFromSource[T any](source []T) T {
	return source[RandomIntBetween(0, len(source))]
}

// Pick multiple DISTINCT item from source
func RandomMultipleFromSource(source []string, max int) []string {
	randomCount := RandomIntBetween(1, max+1)
	if randomCount >= len(source) {
		return source
	}

	pickedItems := make([]string, 0)
	picked := make(map[int]bool)
	for len(pickedItems) < randomCount {
		randomIdx := RandomIntBetween(0, len(source))
		if picked[randomIdx] {
			continue
		}

		picked[randomIdx] = true
		pickedItems = append(pickedItems, source[randomIdx])
	}

	return pickedItems
}

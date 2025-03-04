package mcok

import (
	"givemesomething/mcok/config"
	"math/rand"
	"time"

	"github.com/mroth/weightedrand/v2"
)

func randomPassenger() config.Passenger {
	age := randomIntBetween(1, 80)
	ageRange := config.NewAgeRange(age)

	country := randomFromSource(config.AvailableCountries)
	name := randomFromSource(config.PassengerNames)

	foodCacheKey := config.GenerateCacheKey(config.UPrefFood, mixAgeRange(ageRange), mixCountry(country))
	foods := randomMultipleFromSource(config.ConfigCache[foodCacheKey], 3)

	movieCacheKey := config.GenerateCacheKey(config.UPrefMovie, mixAgeRange(ageRange), mixCountry(country))
	movies := randomMultipleFromSource(config.ConfigCache[movieCacheKey], 3)

	magazineCacheKey := config.GenerateCacheKey(config.UPrefMagazine, mixAgeRange(ageRange), mixCountry(country))
	magazines := randomMultipleFromSource(config.ConfigCache[magazineCacheKey], 3)

	return config.Passenger{
		Name:      name,
		Age:       age,
		Country:   country,
		Foods:     foods,
		Movies:    movies,
		Magazines: magazines,
	}
}

func randomMultipleFromSource(source []string, max int) []string {
	randomCount := randomIntBetween(1, max+1)
	if randomCount >= len(source) {
		return source
	}

	pickedItems := make([]string, 0)
	picked := make(map[int]bool)
	for len(pickedItems) < randomCount {
		randomIdx := randomIntBetween(0, len(source))
		if picked[randomIdx] {
			continue
		}

		picked[randomIdx] = true
		pickedItems = append(pickedItems, source[randomIdx])
	}

	return pickedItems
}

func mixAgeRange(ageRange config.AgeRange) config.AgeRange {
	choices := make([]weightedrand.Choice[config.AgeRange, int], 0)
	for _, currentAgeRange := range config.AvailableAgeRanges {
		if ageRange == currentAgeRange {
			choices = append(choices, weightedrand.NewChoice(currentAgeRange, 2))
		} else {
			choices = append(choices, weightedrand.NewChoice(currentAgeRange, 1))
		}
	}

	chooser, _ := weightedrand.NewChooser(choices...)
	result := chooser.Pick()

	if result == ageRange {
		return result
	}

	return ageRange.Previous()
}

func mixCountry(country config.Country) config.Country {
	choices := make([]weightedrand.Choice[config.Country, int], 0)
	for _, currentCountry := range config.AvailableCountries {
		if country == currentCountry {
			choices = append(choices, weightedrand.NewChoice(currentCountry, 2))
		} else {
			choices = append(choices, weightedrand.NewChoice(currentCountry, 1))
		}
	}

	chooser, _ := weightedrand.NewChooser(choices...)
	return chooser.Pick()
}

func randomFromSource[T any](source []T) T {
	return source[randomIntBetween(0, len(source))]
}

// randomIntBetween [min, max)
func randomIntBetween(min int, max int) int {
	randomSeed := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	return randomSeed.Intn(max-min) + min
}

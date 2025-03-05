package generator

import (
	"givemesomething/mcok/config"
)

func RandomPassenger() config.Passenger {
	age := RandomIntBetween(1, 80)
	ageRange := config.NewAgeRange(age)

	country := RandomFromSource(config.AvailableCountries)
	name := RandomFromSource(config.PassengerNames)

	foodCacheKey := config.GenerateCacheKey(config.UPrefFood, mixAgeRange(ageRange), mixCountry(country))
	foods := RandomMultipleFromSource(config.ConfigCache[foodCacheKey], 3)

	movieCacheKey := config.GenerateCacheKey(config.UPrefMovie, mixAgeRange(ageRange), mixCountry(country))
	movies := RandomMultipleFromSource(config.ConfigCache[movieCacheKey], 3)

	magazineCacheKey := config.GenerateCacheKey(config.UPrefMagazine, mixAgeRange(ageRange), mixCountry(country))
	magazines := RandomMultipleFromSource(config.ConfigCache[magazineCacheKey], 3)

	return config.Passenger{
		Name:      name,
		Age:       age,
		Country:   country,
		Foods:     foods,
		Movies:    movies,
		Magazines: magazines,
	}
}

func mixAgeRange(ageRange config.AgeRange) config.AgeRange {
	result := MixWithBias(ageRange, config.AvailableAgeRanges, 90)
	if result == ageRange {
		return result
	}
	return ageRange.Previous()
}

func mixCountry(country config.Country) config.Country {
	return MixWithBias(country, config.AvailableCountries, 90)
}

package config

import (
	"fmt"
	"strings"
)

type Country string

const (
	Vietnam Country = "Vietnam"
	India   Country = "India"
	USA     Country = "USA"
)

func (country Country) IsValid() bool {
	return country == Vietnam || country == India || country == USA
}

type AgeRange string

const (
	Children AgeRange = "children"
	Teenager AgeRange = "teenager"
	Adult    AgeRange = "adult"
)

func (ageRange AgeRange) IsValid() bool {
	return ageRange == Children || ageRange == Teenager || ageRange == Adult
}

func NewAgeRange(age int) AgeRange {
	if age <= 12 {
		return Children
	}

	if age <= 18 {
		return Teenager
	}

	return Adult
}

func (ageRange AgeRange) Previous() AgeRange {
	if ageRange == Adult {
		return Teenager
	}

	if ageRange == Teenager {
		return Children
	}

	return Adult
}

type Passenger struct {
	Name      string
	Age       int
	Country   Country
	Foods     []string
	Movies    []string
	Magazines []string
}

func (passenger Passenger) ToString() string {
	return fmt.Sprintf(
		"%s,%d,%s,\"%s\",\"%s\",\"%s\"",
		passenger.Name,
		passenger.Age,
		passenger.Country,
		strings.Join(passenger.Foods, ","),
		strings.Join(passenger.Movies, ","),
		strings.Join(passenger.Magazines, ","),
	)
}

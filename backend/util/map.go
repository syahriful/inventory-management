package util

import (
	"errors"
	"math"
)

type UnitOfMass struct {
	NameAcronym     string
	NameDescription string
	Position        int
}

var (
	unitOfMass = map[string]UnitOfMass{
		"ton": {
			NameAcronym:     "ton",
			NameDescription: "Ton",
			Position:        10,
		},
		"kg": {
			NameAcronym:     "kg",
			NameDescription: "Kilogram",
			Position:        7,
		},
		"hg": {
			NameAcronym:     "hg",
			NameDescription: "Hectogram",
			Position:        6,
		},
		"dag": {
			NameAcronym:     "dag",
			NameDescription: "Decagram",
			Position:        5,
		},
		"g": {
			NameAcronym:     "g",
			NameDescription: "Gram",
			Position:        4,
		},
		"dg": {
			NameAcronym:     "dg",
			NameDescription: "Decigram",
			Position:        3,
		},
		"cg": {
			NameAcronym:     "cg",
			NameDescription: "Centigram",
			Position:        2,
		},
		"mg": {
			NameAcronym:     "mg",
			NameDescription: "Milligram",
			Position:        1,
		},
	}
)

func CalculateUnitOfMass(currentUnitOfMass string, newUnitOfMassName string, newValueOfMass float64) (float64, error) {
	currentMass, ok := unitOfMass[currentUnitOfMass]
	if !ok {
		return 0, errors.New("current unit of mass not found")
	}
	newMass, ok := unitOfMass[newUnitOfMassName]
	if !ok {
		return 0, errors.New("new unit of mass not found")
	}

	currentPosition := currentMass.Position
	newPosition := newMass.Position

	if currentPosition == newPosition {
		return newValueOfMass, nil
	}

	count := newPosition - currentPosition
	if count < 0 {
		count *= -1
	}

	scaleFactor := math.Pow10(count)
	if newPosition > currentPosition {
		newValueOfMass *= scaleFactor
	} else if newPosition < currentPosition {
		newValueOfMass /= scaleFactor
	}

	return newValueOfMass, nil
}

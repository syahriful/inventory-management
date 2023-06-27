package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateUnitOfMass(t *testing.T) {
	t.Run("g to g", func(t *testing.T) {
		newUnitOfMass := "g"
		newValueOfMass := float64(4)

		currentUnitOfMass := "g"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 4, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("kg to g", func(t *testing.T) {
		newUnitOfMass := "kg"
		newValueOfMass := float64(4)

		currentUnitOfMass := "g"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 4000, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("g to kg", func(t *testing.T) {
		newUnitOfMass := "g"
		newValueOfMass := float64(257)

		currentUnitOfMass := "kg"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 0.257, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("kg to mg", func(t *testing.T) {
		newUnitOfMass := "kg"
		newValueOfMass := float64(4)

		currentUnitOfMass := "mg"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 4000000, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("hg to kg", func(t *testing.T) {
		newUnitOfMass := "hg"
		newValueOfMass := float64(1)

		currentUnitOfMass := "kg"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 0.1, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("mg to dag", func(t *testing.T) {
		newUnitOfMass := "mg"
		newValueOfMass := float64(500)

		currentUnitOfMass := "dag"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 0.05, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("cg to mg", func(t *testing.T) {
		newUnitOfMass := "cg"
		newValueOfMass := float64(250)

		currentUnitOfMass := "mg"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 2500, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("kg to ton", func(t *testing.T) {
		newUnitOfMass := "kg"
		newValueOfMass := float64(3000)

		currentUnitOfMass := "ton"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 3, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("ton to kg", func(t *testing.T) {
		newUnitOfMass := "ton"
		newValueOfMass := float64(5)

		currentUnitOfMass := "kg"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 5000, total, 0.001, "The decimal values are not within the allowed delta.")
	})

	t.Run("g to cg", func(t *testing.T) {
		newUnitOfMass := "g"
		newValueOfMass := float64(50)

		currentUnitOfMass := "cg"
		total, err := CalculateUnitOfMass(currentUnitOfMass, newUnitOfMass, newValueOfMass)
		assert.Nil(t, err)

		assert.InDelta(t, 5000, total, 0.001, "The decimal values are not within the allowed delta.")
	})
}

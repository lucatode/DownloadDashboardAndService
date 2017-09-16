package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLocation(t *testing.T) {
	t.Run("Get country from coordinates", func(t *testing.T) {

		country := NewLocation("45", "45")

		expectedCountry := "ru"

		assert.Equal(t, expectedCountry, string(country))

	})
}

func TestNewLocation2(t *testing.T) {
	t.Run("Get country from coordinates", func(t *testing.T) {

		country := NewLocation("45.467491", "9.138257")

		expectedCountry := "it"

		assert.Equal(t, expectedCountry, string(country))

	})
}

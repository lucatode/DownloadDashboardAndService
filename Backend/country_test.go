package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLocation(t *testing.T) {
	t.Run("Get country from coordinates", func(t *testing.T) {
		country := NewLocation(45, 45)

		expectedCountry := ""

		assert.Equal(t, expectedCountry, country)

	})
}

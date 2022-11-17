package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanTurnOnLight(t *testing.T) {
	t.Run("Turn On light", func(t *testing.T) {
		light := NewLight(false)
		assert.True(t, light.TurnOn().IsOn(), "Turn On light")
	})
}

func TestCanTurnOffLight(t *testing.T) {
	t.Run("Turn off light", func(t *testing.T) {
		light := NewLight(true)
		assert.False(t, light.TurnOff().IsOn(), "Turn off light")
	})
}

func TestCanToggleLight(t *testing.T) {
	t.Run("Toggle light", func(t *testing.T) {
		light := NewLight(false)
		assert.True(t, light.Toggle().IsOn(), "Turn On light")
		assert.False(t, light.Toggle().IsOn(), "Turn off light")
	})
}

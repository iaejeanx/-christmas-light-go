package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCannotCreateCoordinateWithNegativeX(t *testing.T) {
	t.Run("Fail creating mesh with negative X length", func(t *testing.T) {
		_, coordinateException := NewCoordinate(-1, 1)
		assert.Error(t, coordinateException, "Fail creating mesh with negative X length")
	})
}

func TestCannotCreateCoordinateWithNegativeY(t *testing.T) {
	t.Run("Fail creating mesh with negative Y length", func(t *testing.T) {
		_, coordinateException := NewCoordinate(1, -1)
		assert.Error(t, coordinateException, "Fail creating mesh with negative Y length")
	})
}

func TestCanGetX(t *testing.T) {
	t.Run("Get right X point", func(t *testing.T) {
		coordinate, _ := NewCoordinate(10, 0)
		assert.Equal(t, 10, coordinate.GetX(), "Get right X point")
	})
}

func TestCanGetY(t *testing.T) {
	t.Run("Get right Y point", func(t *testing.T) {
		coordinate, _ := NewCoordinate(0, 10)
		assert.Equal(t, 10, coordinate.GetY(), "Get right Y point")
	})
}

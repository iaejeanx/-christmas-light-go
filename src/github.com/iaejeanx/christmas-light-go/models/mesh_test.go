package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCannotCreateMeshWithNegativeX(t *testing.T) {
	t.Run("Fail creating mesh with negative X length", func(t *testing.T) {
		_, meshException := NewMeshLight(-1000, 1000)
		assert.Error(t, meshException, "Fail creating mesh with negative X length")
	})
}

func TestCannotCreateMeshWithNegativeY(t *testing.T) {
	t.Run("Fail creating mesh with negative Y length", func(t *testing.T) {
		_, meshException := NewMeshLight(1000, -1000)
		assert.Error(t, meshException, "Fail creating mesh with negative Y length")
	})
}

func TestCannotMoveMeshWithCoordinateStartXInvalid(t *testing.T) {
	t.Run("Test cannot move on mesh with coordinate start X invalid", func(t *testing.T) {
		coordinateStart, _ := NewCoordinate(1001, 1000)
		coordinateEnd, _ := NewCoordinate(1000, 1000)
		mesh, _ := NewMeshLight(1000, 1000)
		_, turnOnMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, turnOffMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, toggleMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)

		assert.Error(t, turnOnMeshException, "Test cannot turn on mesh with coordinate start X invalid")
		assert.Error(t, turnOffMeshException, "Test cannot turn off mesh with coordinate start X invalid")
		assert.Error(t, toggleMeshException, "Test cannot toggle mesh with coordinate start X invalid")
	})
}

func TestCannotTurnOnMeshWithCoordinateStartYInvalid(t *testing.T) {
	t.Run("Test cannot move on mesh with coordinate start Y invalid", func(t *testing.T) {
		coordinateStart, _ := NewCoordinate(1000, 1001)
		coordinateEnd, _ := NewCoordinate(1000, 1000)
		mesh, _ := NewMeshLight(1000, 1000)
		_, turnOnMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, turnOffMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, toggleMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)

		assert.Error(t, turnOnMeshException, "Test cannot turn on mesh with coordinate start y invalid")
		assert.Error(t, turnOffMeshException, "Test cannot turn off mesh with coordinate start y invalid")
		assert.Error(t, toggleMeshException, "Test cannot toggle mesh with coordinate start y invalid")
	})
}

func TestCannotTurnOnMeshWithCoordinateEndXInvalid(t *testing.T) {
	t.Run("Test cannot move on mesh with coordinate end X invalid", func(t *testing.T) {
		coordinateStart, _ := NewCoordinate(1000, 1000)
		coordinateEnd, _ := NewCoordinate(1001, 1000)
		mesh, _ := NewMeshLight(1000, 1000)
		_, turnOnMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, turnOffMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, toggleMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)

		assert.Error(t, turnOnMeshException, "Test cannot turn on mesh with coordinate end x invalid")
		assert.Error(t, turnOffMeshException, "Test cannot turn off mesh with coordinate end X invalid")
		assert.Error(t, toggleMeshException, "Test cannot toggle mesh with coordinate end X invalid")
	})
}

func TestCannotTurnOnMeshWithCoordinateEndYInvalid(t *testing.T) {
	t.Run("Test cannot move on mesh with coordinate end Y invalid", func(t *testing.T) {
		coordinateStart, _ := NewCoordinate(1000, 1000)
		coordinateEnd, _ := NewCoordinate(1000, 1001)
		mesh, _ := NewMeshLight(1000, 1000)
		_, turnOnMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, turnOffMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)
		_, toggleMeshException := mesh.TurnOnRange(coordinateStart, coordinateEnd)

		assert.Error(t, turnOnMeshException, "Test cannot turn on mesh with coordinate end Y invalid")
		assert.Error(t, turnOffMeshException, "Test cannot turn off mesh with coordinate end Y invalid")
		assert.Error(t, toggleMeshException, "Test cannot toggle mesh with coordinate end Y invalid")
	})
}

func TestCanTurnOnMesh(t *testing.T) {
	t.Run("Mesh is turn on", func(t *testing.T) {
		mesh, _ := NewMeshLight(1000, 1000)
		assert.True(t, mesh.TurnOnMesh().IsOn(), "Mesh is turn on")
	})
}

func TestCanTurnOffMesh(t *testing.T) {
	t.Run("Mesh is turn off", func(t *testing.T) {
		mesh, _ := NewMeshLight(1000, 1000)
		assert.False(t, mesh.TurnOffMesh().IsOn(), "Mesh is turn off")
	})
}

func TestCanGetLightsTurnOnMesh(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mesh, _ := NewMeshLight(1000, 1000)

		assert.Equal(t, 0, mesh.GetLightsTurnOn(), "All mesh is turn off")

		mesh.TurnOnMesh()
		assert.Equal(t, 1000*1000, mesh.GetLightsTurnOn(), "All mesh is turn on")

		mesh.TurnOffMesh()
		assert.Equal(t, 0, mesh.GetLightsTurnOn(), "All mesh is turn off")
	})
}

func TestCanTurnOnRangeMesh(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mesh, _ := NewMeshLight(1000, 1000)

		assert.Equal(t, 0, mesh.GetLightsTurnOn(), "All mesh is turn off")

		startPoint, _ := NewCoordinate(0, 0)
		endPoint, _ := NewCoordinate(0, 4)
		meshWithRangeOn, _ := mesh.TurnOnRange(startPoint, endPoint)
		assert.Equal(t, 1*5, meshWithRangeOn.GetLightsTurnOn(), "Mesh is turn on in 0, 5 & 0, 5")
	})
}

func TestCanTurnOffRangeMesh(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mesh, _ := NewMeshLight(1000, 1000)

		assert.Equal(t, 1000*1000, mesh.TurnOnMesh().GetLightsTurnOn(), "All mesh is turn on")

		startPoint, _ := NewCoordinate(0, 0)
		endPoint, _ := NewCoordinate(0, 4)
		meshWithRangeOff, _ := mesh.TurnOffRange(startPoint, endPoint)
		assert.Equal(t, 1000*1000-1*5, meshWithRangeOff.GetLightsTurnOn(), "Mesh is turn off in 0, 5 & 0, 5")
	})
}

func TestCanToggleRangeMesh(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mesh, _ := NewMeshLight(1000, 1000)

		assert.Equal(t, 0, mesh.GetLightsTurnOn(), "All mesh is turn off")

		startPointA, _ := NewCoordinate(0, 0)
		endPointA, _ := NewCoordinate(0, 4)
		meshWithToggleA, _ := mesh.ToggleRange(startPointA, endPointA)
		assert.Equal(t, 1*5, meshWithToggleA.GetLightsTurnOn(), "Mesh is turn on in 0, 5 and  0, 5")

		startPointB, _ := NewCoordinate(0, 0)
		endPointB, _ := NewCoordinate(0, 4)
		meshWithToggleB, _ := mesh.ToggleRange(startPointB, endPointB)
		assert.Equal(t, 0, meshWithToggleB.GetLightsTurnOn(), "All mesh is turn off")
	})
}

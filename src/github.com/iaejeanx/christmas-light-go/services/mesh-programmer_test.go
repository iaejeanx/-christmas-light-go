package services

import (
	"github.com/iaejeanx/christmas-light-go/github.com/iaejeanx/christmas-light-go/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanTurnOffRange(t *testing.T) {
	t.Run("", func(t *testing.T) {
		meshProgrammer := MeshProgrammer{}
		mesh, _ := models.NewMeshLight(1000, 1000)
		mesh.TurnOnMesh()

		assert.Equal(t, 1000*1000, mesh.GetLightsTurnOn(), "Mesh is turn on")

		startPoint, _ := models.NewCoordinate(0, 0)
		endPoint, _ := models.NewCoordinate(999, 0)
		meshWithRangeOff, _ := meshProgrammer.TurnOffRange(mesh, startPoint, endPoint)
		assert.Equal(t, 999*1000, meshWithRangeOff.GetLightsTurnOn(), "Mesh is turn off at 0, 0 and 999, 999")
	})
}

func TestCanTurnOnRange(t *testing.T) {
	t.Run("", func(t *testing.T) {
		meshProgrammer := MeshProgrammer{}
		mesh, _ := models.NewMeshLight(1000, 1000)

		assert.Equal(t, 0, mesh.GetLightsTurnOn(), "Mesh is turn off")

		startPoint, _ := models.NewCoordinate(0, 0)
		endPoint, _ := models.NewCoordinate(999, 0)
		meshWithRangeOn, _ := meshProgrammer.TurnOnRange(mesh, startPoint, endPoint)
		assert.Equal(t, 1*1000, meshWithRangeOn.GetLightsTurnOn(), "Mesh is turn on at 0, 0 and 999, 999")
	})
}

func TestCanToggleRange(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mesh, _ := models.NewMeshLight(1000, 1000)
		meshProgrammer := MeshProgrammer{}

		assert.Equal(t, 0, mesh.GetLightsTurnOn(), "Mesh is turn off")

		startPointA, _ := models.NewCoordinate(0, 0)
		endPointA, _ := models.NewCoordinate(999, 0)
		meshWithToggleA, _ := meshProgrammer.ToggleRange(mesh, startPointA, endPointA)
		assert.Equal(t, 1*1000, meshWithToggleA.GetLightsTurnOn(), "Mesh is turn on at 0, 0 and 999, 999")

		startPointB, _ := models.NewCoordinate(0, 0)
		endPointB, _ := models.NewCoordinate(999, 0)
		meshWithToggleB, _ := meshProgrammer.ToggleRange(mesh, startPointB, endPointB)
		assert.Equal(t, 0, meshWithToggleB.GetLightsTurnOn(), "Mesh is turn off")
	})
}

func TestCanTestExamples(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mesh, _ := models.NewMeshLight(1000, 1000)
		meshProgrammer := MeshProgrammer{}

		startPointA, _ := models.NewCoordinate(0, 0)
		endPointA, _ := models.NewCoordinate(999, 999)
		meshWithRangeA, _ := meshProgrammer.TurnOffRange(mesh, startPointA, endPointA)
		assert.Equal(t, 0, meshWithRangeA.GetLightsTurnOn(), "Mesh is turn off")

		startPointB, _ := models.NewCoordinate(0, 0)
		endPointB, _ := models.NewCoordinate(999, 999)
		meshWithRangeB, _ := meshProgrammer.TurnOnRange(mesh, startPointB, endPointB)
		assert.Equal(t, 1000*1000, meshWithRangeB.GetLightsTurnOn(), "Mesh is turn on")

		startPointC, _ := models.NewCoordinate(0, 0)
		endPointC, _ := models.NewCoordinate(999, 0)
		meshWithRangeC, _ := meshProgrammer.ToggleRange(mesh, startPointC, endPointC)
		assert.Equal(t, 999*1000, meshWithRangeC.GetLightsTurnOn(), "First row is toggle")

		startPointD, _ := models.NewCoordinate(499, 499)
		endPointD, _ := models.NewCoordinate(500, 500)
		meshWithRangeD, _ := meshProgrammer.TurnOffRange(mesh, startPointD, endPointD)
		assert.Equal(t, (999*1000)-4, meshWithRangeD.GetLightsTurnOn(), "Turn off 499, 499 and 500, 500")
	})
}

func TestCanTestWinPattern(t *testing.T) {
	t.Run("", func(t *testing.T) {
		mesh, _ := models.NewMeshLight(1000, 1000)
		meshProgrammer := MeshProgrammer{}

		meshWinner, _ := meshProgrammer.RunWinPattern(mesh)
		assert.Equal(t, 230022, meshWinner.GetLightsTurnOn(), "Checking lights on")
		//assert.Equal(t, 513805, meshProgrammer.RunWinPattern(&mesh).GetLightsTurnOn(), "Checking lights on")
	})
}

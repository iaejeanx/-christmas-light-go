package models

import (
	"github.com/iaejeanx/christmas-light-go/github.com/iaejeanx/christmas-light-go/exceptions"
)

type Mesh struct {
	matrix [][]Light
	on     bool
}

func NewMeshLight(xAxisLength int, yAxisLength int) (Mesh, error) {
	if xAxisLength < 0 {
		return Mesh{}, exceptions.MeshException{}.InvalidSizeForXAxis()
	}

	if yAxisLength < 0 {
		return Mesh{}, exceptions.CoordinateException{}.InvalidArgumentForYAxis()
	}

	var matrix = make([][]Light, xAxisLength)
	for i := range matrix {
		matrix[i] = make([]Light, yAxisLength)
	}

	for x := 0; x < xAxisLength; x++ {
		for y := 0; y < yAxisLength; y++ {
			matrix[x][y] = NewLight(false)
		}
	}

	return Mesh{matrix, false}, nil
}

func (mesh *Mesh) IsOn() bool {
	return mesh.on
}

func (mesh *Mesh) TurnOnMesh() *Mesh {
	for i, row := range mesh.matrix {
		mesh.matrix[i] = TurnOnRow(row)
	}
	mesh.on = true
	return mesh
}

func (mesh *Mesh) TurnOffMesh() *Mesh {
	for i, row := range mesh.matrix {
		mesh.matrix[i] = TurnOffRow(row)
	}
	mesh.on = false
	return mesh
}

func (mesh *Mesh) GetLightsTurnOn() int {
	count := 0
	for _, row := range mesh.matrix {
		count += getLightsTurnOnInRow(row)
	}
	return count
}

func (mesh *Mesh) TurnOnRange(startCoordinate Coordinate, endCoordinate Coordinate) (*Mesh, error) {
	meshException := checkCoordinatesInMesh(mesh, startCoordinate, endCoordinate)
	if meshException != nil {
		return mesh, meshException
	}

	for xIndex, row := range mesh.matrix {
		for yIndex := range row {
			if xIndex >= startCoordinate.GetX() &&
				yIndex >= startCoordinate.GetY() &&
				xIndex <= endCoordinate.GetX() &&
				yIndex <= endCoordinate.GetY() {
				mesh.matrix[xIndex][yIndex].TurnOn()
			}
		}
	}
	return mesh, nil
}

func (mesh *Mesh) TurnOffRange(startCoordinate Coordinate, endCoordinate Coordinate) (*Mesh, error) {
	meshException := checkCoordinatesInMesh(mesh, startCoordinate, endCoordinate)
	if meshException != nil {
		return mesh, meshException
	}

	for xIndex, row := range mesh.matrix {
		for yIndex := range row {
			if xIndex >= startCoordinate.GetX() &&
				yIndex >= startCoordinate.GetY() &&
				xIndex <= endCoordinate.GetX() &&
				yIndex <= endCoordinate.GetY() {
				mesh.matrix[xIndex][yIndex].TurnOff()
			}
		}
	}
	return mesh, nil
}

func (mesh *Mesh) ToggleRange(startCoordinate Coordinate, endCoordinate Coordinate) (*Mesh, error) {
	meshException := checkCoordinatesInMesh(mesh, startCoordinate, endCoordinate)
	if meshException != nil {
		return mesh, meshException
	}

	for xIndex, row := range mesh.matrix {
		for yIndex := range row {
			if xIndex >= startCoordinate.GetX() &&
				yIndex >= startCoordinate.GetY() &&
				xIndex <= endCoordinate.GetX() &&
				yIndex <= endCoordinate.GetY() {
				mesh.matrix[xIndex][yIndex].Toggle()
			}
		}
	}
	return mesh, nil
}

func checkCoordinatesInMesh(mesh *Mesh, startCoordinate Coordinate, endCoordinate Coordinate) error {
	if startCoordinate.GetX() > len(mesh.matrix) {
		return exceptions.CoordinateException{}.CoordinateStartXIsBiggerThanMesh()
	}

	if startCoordinate.GetY() > len(mesh.matrix[0]) {
		return exceptions.CoordinateException{}.CoordinateStartYIsBiggerThanMesh()
	}

	if endCoordinate.GetX() > len(mesh.matrix) {
		return exceptions.CoordinateException{}.CoordinateEndXIsBiggerThanMesh()
	}

	if endCoordinate.GetY() > len(mesh.matrix[0]) {
		return exceptions.CoordinateException{}.CoordinateEndYIsBiggerThanMesh()
	}

	return nil
}

func getLightsTurnOnInRow(row []Light) int {
	count := 0
	for _, light := range row {
		if light.IsOn() {
			count++
		}
	}
	return count
}

func TurnOnRow(row []Light) []Light {
	newRow := make([]Light, len(row))
	for i, light := range row {
		light.TurnOn()
		newRow[i] = light

	}
	return newRow
}

func TurnOffRow(row []Light) []Light {
	newRow := make([]Light, len(row))
	for i, light := range row {
		light.TurnOff()
		newRow[i] = light

	}
	return newRow
}

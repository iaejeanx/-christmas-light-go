package contracts

import "github.com/iaejeanx/christmas-light-go/github.com/iaejeanx/christmas-light-go/models"

type MeshProgrammerInterface interface {
	TurnOnRange(mesh models.Mesh, startCoordinate models.Coordinate, endCoordinate models.Coordinate) (models.Mesh, error)

	TurnOffRange(mesh models.Mesh, startCoordinate models.Coordinate, endCoordinate models.Coordinate) (models.Mesh, error)

	ToggleRange(mesh models.Mesh, startCoordinate models.Coordinate, endCoordinate models.Coordinate) (models.Mesh, error)

	RunWinPattern(mesh models.Mesh) (models.Mesh, error)
}

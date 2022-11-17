package services

import "github.com/iaejeanx/christmas-light-go/github.com/iaejeanx/christmas-light-go/models"

type MeshProgrammer struct {
}

func (meshProgrammer MeshProgrammer) TurnOnRange(mesh models.Mesh, startCoordinate models.Coordinate, endCoordinate models.Coordinate) (*models.Mesh, error) {
	return mesh.TurnOnRange(startCoordinate, endCoordinate)
}

func (meshProgrammer MeshProgrammer) TurnOffRange(mesh models.Mesh, startCoordinate models.Coordinate, endCoordinate models.Coordinate) (*models.Mesh, error) {
	return mesh.TurnOffRange(startCoordinate, endCoordinate)
}

func (meshProgrammer MeshProgrammer) ToggleRange(mesh models.Mesh, startCoordinate models.Coordinate, endCoordinate models.Coordinate) (*models.Mesh, error) {
	return mesh.ToggleRange(startCoordinate, endCoordinate)
}

func (meshProgrammer MeshProgrammer) RunWinPattern(mesh models.Mesh) (*models.Mesh, error) {
	startPointA, _ := models.NewCoordinate(887, 9)
	endPointA, _ := models.NewCoordinate(959, 629)
	startPointB, _ := models.NewCoordinate(454, 398)
	endPointB, _ := models.NewCoordinate(844, 448)
	startPointC, _ := models.NewCoordinate(539, 243)
	endPointC, _ := models.NewCoordinate(559, 965)
	startPointD, _ := models.NewCoordinate(370, 819)
	endPointD, _ := models.NewCoordinate(676, 868)
	startPointE, _ := models.NewCoordinate(145, 40)
	endPointE, _ := models.NewCoordinate(370, 997)
	startPointF, _ := models.NewCoordinate(301, 3)
	endPointF, _ := models.NewCoordinate(808, 453)
	startPointG, _ := models.NewCoordinate(351, 678)
	endPointG, _ := models.NewCoordinate(951, 908)
	startPointH, _ := models.NewCoordinate(720, 196)
	endPointH, _ := models.NewCoordinate(897, 994)
	startPointI, _ := models.NewCoordinate(831, 394)
	endPointI, _ := models.NewCoordinate(904, 860)
	//mesh.TurnOnMesh().
	meshRangeA, _ := mesh.TurnOnRange(startPointA, endPointA)
	meshRangeB, _ := meshRangeA.TurnOnRange(startPointB, endPointB)
	meshRangeC, _ := meshRangeB.TurnOffRange(startPointC, endPointC)
	meshRangeD, _ := meshRangeC.TurnOffRange(startPointD, endPointD)
	meshRangeE, _ := meshRangeD.TurnOffRange(startPointE, endPointE)
	meshRangeF, _ := meshRangeE.TurnOffRange(startPointF, endPointF)
	meshRangeG, _ := meshRangeF.TurnOnRange(startPointG, endPointG)
	meshRangeH, _ := meshRangeG.ToggleRange(startPointH, endPointH)
	meshRangeI, _ := meshRangeH.ToggleRange(startPointI, endPointI)
	return meshRangeI, nil
}

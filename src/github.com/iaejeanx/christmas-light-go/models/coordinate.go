package models

import "github.com/iaejeanx/christmas-light-go/github.com/iaejeanx/christmas-light-go/exceptions"

type Coordinate struct {
	X int
	y int
}

func NewCoordinate(x int, y int) (Coordinate, error) {
	if x < 0 {
		return Coordinate{}, exceptions.CoordinateException{}.InvalidArgumentForXAxis()
	}

	if y < 0 {
		return Coordinate{}, exceptions.CoordinateException{}.InvalidArgumentForYAxis()
	}

	return Coordinate{x, y}, nil
}

func (coordinate Coordinate) GetX() int {
	return coordinate.X
}

func (coordinate Coordinate) GetY() int {
	return coordinate.y
}

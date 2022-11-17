package exceptions

type CoordinateException struct {
	Message string
	Code    string
}

func (coordinateException CoordinateException) Error() string {
	return coordinateException.Message
}

func (coordinateException CoordinateException) InvalidArgumentForXAxis() CoordinateException {
	return CoordinateException{"invalid argument for X axis", "100"}
}

func (coordinateException CoordinateException) InvalidArgumentForYAxis() CoordinateException {
	return CoordinateException{"invalid argument for Y axis", "101"}
}

func (coordinateException CoordinateException) CoordinateStartXIsBiggerThanMesh() CoordinateException {
	return CoordinateException{"Coordinate start X s bigger than mesh", "102"}
}

func (coordinateException CoordinateException) CoordinateStartYIsBiggerThanMesh() CoordinateException {
	return CoordinateException{"Coordinate start / s bigger than mesh", "103"}
}

func (coordinateException CoordinateException) CoordinateEndXIsBiggerThanMesh() CoordinateException {
	return CoordinateException{"Coordinate end X s bigger than mesh", "104"}
}

func (coordinateException CoordinateException) CoordinateEndYIsBiggerThanMesh() CoordinateException {
	return CoordinateException{"Coordinate end Y s bigger than mesh", "105"}
}

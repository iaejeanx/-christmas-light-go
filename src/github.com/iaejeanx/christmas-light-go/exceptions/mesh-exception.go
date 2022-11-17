package exceptions

type MeshException struct {
	Message string
	Code    string
}

func (meshException MeshException) Error() string {
	return meshException.Message
}

func (meshException MeshException) InvalidSizeForXAxis() MeshException {
	return MeshException{"Invalid size for X axis", "200"}
}

func (meshException MeshException) InvalidSizeForYAxis() MeshException {
	return MeshException{"Invalid size for Y axis", "201"}
}

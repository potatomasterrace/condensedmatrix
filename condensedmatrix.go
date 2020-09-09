package condensedmatrix

type CondensedMatrix interface {
	Position(offset int) (i, j int)
	Offset(i, j int) int
	ForEachPosition(callback func(i, j int))
	ForEachPositionMultiThreaded(callback func(i, j int))
	Size() int
	RealSize() int
}

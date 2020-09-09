package condensedmatrix

// CondensedMatrix is the interface used for both CachedMatrix and Matrix
type CondensedMatrix interface {
	Position(offset int) (i, j int)
	Offset(i, j int) int
	ForEachPosition(callback func(i, j int))
	ForEachPositionMultiThreaded(callback func(i, j int))
	Size() int
	RealSize() int
}

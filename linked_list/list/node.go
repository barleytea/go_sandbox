package list

type node[T any] struct {
	value T
	prev *node[T]
	next *node[T]
}
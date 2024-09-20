package datautils

import "C"
import "runtime"

type Vector[T any] struct {
	Size     int
	Capacity int
	Data     T
	pinner   *runtime.Pinner
}

func NewVectorFromC[T any](size, capacity C.int, data T) Vector[T] {
	return Vector[T]{Size: int(size), Capacity: int(capacity), Data: data, pinner: &runtime.Pinner{}}
}

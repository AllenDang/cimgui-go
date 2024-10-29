package vectors

import "C"
import "runtime"

type Vector[T any] struct {
	Size     int
	Capacity int
	Data     *T
	pinner   *runtime.Pinner
}

func NewVectorFromC[T any, CINT ~int32](size, capacity CINT, data *T) Vector[T] {
	return Vector[T]{Size: int(size), Capacity: int(capacity), Data: data, pinner: &runtime.Pinner{}}
}

func (v *Vector[T]) Pinner() *runtime.Pinner {
	return v.pinner
}

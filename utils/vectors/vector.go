package vectors

import "C"
import (
	"runtime"
	"unsafe"
)

// Vector is a representation of ImVector from Dear ImGui.
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

// Slice converts a Vector to a slice []T.
func (v Vector[T]) Slice() []T {
	// We consider v.Data a pointer to 1st element of a slice []T
	// of size v.Size and capacity v.Capacity
	return unsafe.Slice(v.Data, v.Size)
}

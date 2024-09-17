package imgui

// #include <memory.h>
// #include <stdlib.h>
// #include <stdbool.h>
import "C"

import (
	"unsafe"
)

// unrealisticLargePointer is used to cast an arbitrary native pointer to a slice.
// Its value is chosen to fit into a 32bit architecture, and still be large
// enough to cover "any" data blob. Note that this value is in bytes.
// Should an array of larger primitives be addressed, be sure to divide the value
// by the size of the elements.
const unrealisticLargePointer = 1 << 30

func PtrToByteSlice(p unsafe.Pointer) []byte {
	return (*[unrealisticLargePointer]byte)(p)[:]
}

func PtrToUint16Slice(p unsafe.Pointer) []uint16 {
	return (*[unrealisticLargePointer / 2]uint16)(p)[:]
}

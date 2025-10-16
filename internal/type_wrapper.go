package internal

// #include <memory.h>
// #include <stdlib.h>
import "C"

import (
	"unsafe"
)

// Char represents C.char. On amd64 it is equivalent to int8, on arm64 it is uint8.
// See https://github.com/AllenDang/cimgui-go/issues/140
type Char interface {
	~int8 | ~uint8
}

// ReinterpretCast acts exactly like C++'s reinterpret_cast.
// Intendedd use is to convert packageA.C.MyType to packageB.C.MyType.
// make sure your types are identical C types before using it.
// THIS IS HIGHLY UNSAFE (even more than Go's unsafe package) AND NOT RECOMMENDED TO USE OUTSIDE CIMGUI.
// It just forces pointer/type reinterpretation with unsafe.Pointer.
func ReinterpretCast[RET, SRC any](src SRC) RET {
	return *(*RET)(unsafe.Pointer(&src))
}

// WrapString converts Go string to C char*
// Default value of RET is C.char
func WrapString[RET Char](value string) (wrapped *RET, finisher func()) {
	wrapped = ReinterpretCast[*RET](C.CString(value))
	finisher = func() { C.free(unsafe.Pointer(wrapped)) } // nolint: gas
	return
}

// WrapStringList converts Go string slice to C char**
// Default value of RET is C.char
func WrapStringList[RET Char](value []string) (wrapped **RET, finisher func()) {
	if len(value) == 0 {
		return nil, func() {}
	}

	wrappedList := make([]*C.char, len(value))
	for i, v := range value {
		wrappedList[i] = C.CString(v)
	}

	wrapped = (**RET)(unsafe.Pointer(&wrappedList[0]))

	finisher = func() {
		for _, v := range wrappedList {
			C.free(unsafe.Pointer(v))
		}
	}

	return
}

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

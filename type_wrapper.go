package imgui

// #include <memory.h>
// #include <stdlib.h>
// #include <stdbool.h>
import "C"
import (
	"runtime"
	"unsafe"
)

func CastBool(value bool) (cast int) {
	if value {
		cast = 1
	}
	return
}

func WrapBool(goValue *bool) (wrapped *C.bool, finisher func()) {
	if goValue != nil {
		var cValue C.bool
		if *goValue {
			cValue = C.bool(true)
		}
		wrapped = &cValue
		finisher = func() {
			*goValue = cValue == C.bool(true)
		}
	} else {
		finisher = func() {}
	}
	return
}

// Number is a generic type for Go/C types that can be used as a number.
// It could be anything that you can convert to that type (e.g. C.int is a Number,
// because it can be directly converted to int)
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// WrapNumberPtr is a generic method to convert GOTYPE (int32/float32 e.t.c.) into CTYPE (c_int/c_float e.t.c.)
func WrapNumberPtr[CTYPE Number, GOTYPE Number](goValue *GOTYPE) (wrapped *CTYPE, finisher func()) {
	if goValue != nil {
		cValue := CTYPE(*goValue)
		wrapped = &cValue
		finisher = func() {
			*goValue = GOTYPE(cValue)
		}
	} else {
		finisher = func() {}
	}

	return
}

func WrapString(value string) (wrapped *C.char, finisher func()) {
	wrapped = C.CString(value)
	finisher = func() { C.free(unsafe.Pointer(wrapped)) } // nolint: gas
	return
}

func WrapStringList(value []string) (wrapped **C.char, finisher func()) {
	if len(value) == 0 {
		return nil, func() {}
	}

	wrappedList := make([]*C.char, len(value))
	for i, v := range value {
		wrappedList[i] = C.CString(v)
	}

	wrapped = (**C.char)(unsafe.Pointer(&wrappedList[0]))

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

type StringBuffer struct {
	ptr  unsafe.Pointer
	size int
}

func NewStringBuffer(initialValue string) *StringBuffer {
	rawText := []byte(initialValue)
	bufSize := len(rawText) + 1
	newPtr := C.malloc(C.size_t(bufSize))
	zeroOffset := bufSize - 1
	buf := PtrToByteSlice(newPtr)
	copy(buf[:zeroOffset], rawText)
	buf[zeroOffset] = 0

	return &StringBuffer{ptr: newPtr, size: bufSize}
}

func (buf *StringBuffer) Free() {
	C.free(buf.ptr)
	buf.size = 0
}

func (buf *StringBuffer) ResizeTo(requestedSize int) {
	bufSize := requestedSize
	if bufSize < 1 {
		bufSize = 1
	}
	newPtr := C.malloc(C.size_t(bufSize))
	copySize := bufSize
	if copySize > buf.size {
		copySize = buf.size
	}
	if copySize > 0 {
		C.memcpy(newPtr, buf.ptr, C.size_t(copySize))
	}
	PtrToByteSlice(newPtr)[bufSize-1] = 0
	C.free(buf.ptr)
	buf.ptr = newPtr
	buf.size = bufSize
}

func (buf *StringBuffer) ToGo() string {
	if (buf.ptr == nil) || (buf.size < 1) {
		return ""
	}
	PtrToByteSlice(buf.ptr)[buf.size-1] = 0
	return C.GoString((*C.char)(buf.ptr))
}

// WrapVoidPtr uses runtime.Pinner to pin value
func WrapVoidPtr(value unsafe.Pointer) (wrapped unsafe.Pointer, finisher func()) {
	p := &runtime.Pinner{}
	tryPin(value, p)
	return value, func() {
		p.Unpin()
	}
}

// TODO: this is workaround because of bug/feature request in GO.
// It might be changed after 1.22 release
func tryPin(value any, pinner *runtime.Pinner) {
	defer func() {
		if r := recover(); r != nil {
			// nothing to do here hehe
		}
	}()

	pinner.Pin(value)
}

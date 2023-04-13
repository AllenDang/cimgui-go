package imgui

// #include <memory.h>
// #include <stdlib.h>
// #include <stdbool.h>
import "C"
import "unsafe"

func castBool(value bool) (cast int) {
	if value {
		cast = 1
	}
	return
}

func wrapBool(goValue *bool) (wrapped *C.bool, finisher func()) {
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
	~int8 | ~int16 | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// wrapPtrCType is a generic method to convert GOTYPE (int32/float32 e.t.c.) into CTYPE (c_int/c_float e.t.c.)
func wrapNumberPtr[CTYPE Number, GOTYPE Number](goValue *GOTYPE) (wrapped *CTYPE, finisher func()) {
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

func wrapString(value string) (wrapped *C.char, finisher func()) {
	wrapped = C.CString(value)
	finisher = func() { C.free(unsafe.Pointer(wrapped)) } // nolint: gas
	return
}

func wrapStringList(value []string) (wrapped **C.char, finisher func()) {
	if len(value) == 0 {
		return nil, func() {}
	}

	wrappedList := make([]*C.char, len(value))
	for i, v := range value {
		wrappedList[i] = C.CString(v)
	}

	wrapped = (**C.char)(unsafe.Pointer(&wrappedList[0]))

	finisher = func() {
		C.free(unsafe.Pointer(wrapped))
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

func ptrToByteSlice(p unsafe.Pointer) []byte {
	return (*[unrealisticLargePointer]byte)(p)[:]
}

func ptrToUint16Slice(p unsafe.Pointer) []uint16 {
	return (*[unrealisticLargePointer / 2]uint16)(p)[:]
}

type stringBuffer struct {
	ptr  unsafe.Pointer
	size int
}

func newStringBuffer(initialValue string) *stringBuffer {
	rawText := []byte(initialValue)
	bufSize := len(rawText) + 1
	newPtr := C.malloc(C.size_t(bufSize))
	zeroOffset := bufSize - 1
	buf := ptrToByteSlice(newPtr)
	copy(buf[:zeroOffset], rawText)
	buf[zeroOffset] = 0

	return &stringBuffer{ptr: newPtr, size: bufSize}
}

func (buf *stringBuffer) free() {
	C.free(buf.ptr)
	buf.size = 0
}

func (buf *stringBuffer) resizeTo(requestedSize int) {
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
	ptrToByteSlice(newPtr)[bufSize-1] = 0
	C.free(buf.ptr)
	buf.ptr = newPtr
	buf.size = bufSize
}

func (buf stringBuffer) toGo() string {
	if (buf.ptr == nil) || (buf.size < 1) {
		return ""
	}
	ptrToByteSlice(buf.ptr)[buf.size-1] = 0
	return C.GoString((*C.char)(buf.ptr))
}

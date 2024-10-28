package internal

// #include <memory.h>
// #include <stdlib.h>
// #include <stdbool.h>
import "C"

import (
	"unsafe"
)

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

func (buf *StringBuffer) Ptr() unsafe.Pointer {
	return buf.ptr
}

func (buf *StringBuffer) Size() int {
	return buf.size
}

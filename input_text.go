package cimgui

// #include <memory.h>
// #include <stdlib.h>
// #include "cimgui_wrapper.h"
// #include "cimgui_structs_accessor.h"
// extern void generalInputTextCallback(ImGuiInputTextCallbackData* data);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

//export generalInputTextCallback
func generalInputTextCallback(cb *C.ImGuiInputTextCallbackData) {
	data := ImGuiInputTextCallbackData(unsafe.Pointer(cb))
	if data.GetEventFlag() == ImGuiInputTextFlags_CallbackResize {
		bufHandle := (*cgo.Handle)(data.c().UserData)

		buf := bufHandle.Value().(*stringBuffer)
		buf.resizeTo(data.GetBufSize())
		C.ImGuiInputTextCallbackData_SetBuf(cb, (*C.char)(buf.ptr))
		data.SetBufSize(int32(buf.size))
		data.SetBufTextLen(int32(data.GetBufTextLen()))
		data.SetBufDirty(true)
	}
}

func InputTextWithHint(label, hint string, buf *string, flags ImGuiInputTextFlags) {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	hintArg, hintFin := wrapString(hint)
	defer hintFin()

	sBuf := newStringBuffer(*buf)
	defer func() {
		*buf = sBuf.toGo()
		sBuf.free()
	}()

	sBufHandle := cgo.NewHandle(sBuf)
	defer sBufHandle.Delete()

	flags |= ImGuiInputTextFlags_CallbackResize

	C.InputTextWithHint(
		labelArg,
		hintArg,
		(*C.char)(sBuf.ptr),
		C.ulong(len(*buf)),
		C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.generalInputTextCallback),
		unsafe.Pointer(&sBufHandle),
	)
}

package cimgui

// #include <memory.h>
// #include <stdlib.h>
// #include "cimgui_wrapper.h"
// #include "cimgui_structs_accessor.h"
// extern int generalInputTextCallback(ImGuiInputTextCallbackData* data);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

type ImGuiInputTextCallback func(data ImGuiInputTextCallbackData) int

type inputTextInternalState struct {
	buf      *stringBuffer
	callback ImGuiInputTextCallback
}

func (state *inputTextInternalState) release() {
	state.buf.free()
}

//export generalInputTextCallback
func generalInputTextCallback(cbData *C.ImGuiInputTextCallbackData) C.int {
	data := ImGuiInputTextCallbackData(unsafe.Pointer(cbData))

	bufHandle := (*cgo.Handle)(data.c().UserData)
	statePtr := bufHandle.Value().(*inputTextInternalState)

	if data.GetEventFlag() == ImGuiInputTextFlags_CallbackResize {
		statePtr.buf.resizeTo(data.GetBufSize())
		C.ImGuiInputTextCallbackData_SetBuf(cbData, (*C.char)(statePtr.buf.ptr))
		data.SetBufSize(int32(statePtr.buf.size))
		data.SetBufTextLen(int32(data.GetBufTextLen()))
		data.SetBufDirty(true)
	}

	if statePtr.callback != nil {
		return C.int(statePtr.callback(data))
	}

	return 0
}

func InputTextWithHint(label, hint string, buf *string, flags ImGuiInputTextFlags, callback ImGuiInputTextCallback) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	hintArg, hintFin := wrapString(hint)
	defer hintFin()

	state := &inputTextInternalState{
		buf:      newStringBuffer(*buf),
		callback: callback,
	}

	defer func() {
		*buf = state.buf.toGo()
		state.release()
	}()

	stateHandle := cgo.NewHandle(state)
	defer stateHandle.Delete()

	flags |= ImGuiInputTextFlags_CallbackResize

	return C.InputTextWithHint(
		labelArg,
		hintArg,
		(*C.char)(state.buf.ptr),
		C.ulong(len(*buf)+1),
		C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.generalInputTextCallback),
		unsafe.Pointer(&stateHandle),
	) == C.bool(true)
}

func InputTextMultiline(label string, buf *string, size ImVec2, flags ImGuiInputTextFlags, callback ImGuiInputTextCallback) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	state := &inputTextInternalState{
		buf:      newStringBuffer(*buf),
		callback: callback,
	}

	defer func() {
		*buf = state.buf.toGo()
		state.release()
	}()

	stateHandle := cgo.NewHandle(state)
	defer stateHandle.Delete()

	flags |= ImGuiInputTextFlags_CallbackResize

	return C.InputTextMultiline(
		labelArg,
		(*C.char)(state.buf.ptr),
		C.ulong(len(*buf)+1),
		size.toC(),
		C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.generalInputTextCallback),
		unsafe.Pointer(&stateHandle),
	) == C.bool(true)
}

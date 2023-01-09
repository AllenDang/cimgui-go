package cimgui

// #include <memory.h>
// #include <stdlib.h>
// #include "extra_types.h"
// #include "cimgui_wrapper.h"
// #include "cimgui_structs_accessor.h"
// extern int generalInputTextCallback(ImGuiInputTextCallbackData* data);
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

type InputTextCallback func(data InputTextCallbackData) int

type inputTextInternalState struct {
	buf		*stringBuffer
	callback	InputTextCallback
}

func (state *inputTextInternalState) release() {
	state.buf.free()
}

//export generalInputTextCallback
func generalInputTextCallback(cbData *C.ImGuiInputTextCallbackData) C.int {
	data := InputTextCallbackData(unsafe.Pointer(cbData))

	bufHandle := (*cgo.Handle)(data.c().UserData)
	statePtr := bufHandle.Value().(*inputTextInternalState)

	if data.EventFlag() == InputTextFlagsCallbackResize {
		statePtr.buf.resizeTo(data.BufSize())
		C.ImGuiInputTextCallbackData_SetBuf(cbData, (*C.char)(statePtr.buf.ptr))
		data.SetBufSize(int32(statePtr.buf.size))
		data.SetBufTextLen(int32(data.BufTextLen()))
		data.SetBufDirty(true)
	}

	if statePtr.callback != nil {
		return C.int(statePtr.callback(data))
	}

	return 0
}

func InputTextWithHint(label, hint string, buf *string, flags InputTextFlags, callback InputTextCallback) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	hintArg, hintFin := wrapString(hint)
	defer hintFin()

	state := &inputTextInternalState{
		buf:		newStringBuffer(*buf),
		callback:	callback,
	}

	defer func() {
		*buf = state.buf.toGo()
		state.release()
	}()

	stateHandle := cgo.NewHandle(state)
	defer stateHandle.Delete()

	flags |= InputTextFlagsCallbackResize

	return C.InputTextWithHintV(
		labelArg,
		hintArg,
		(*C.char)(state.buf.ptr),
		C.xlong(len(*buf)+1),
		C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.generalInputTextCallback),
		unsafe.Pointer(&stateHandle),
	) == C.bool(true)
}

func InputTextMultiline(label string, buf *string, size Vec2, flags InputTextFlags, callback InputTextCallback) bool {
	labelArg, labelFin := wrapString(label)
	defer labelFin()

	state := &inputTextInternalState{
		buf:		newStringBuffer(*buf),
		callback:	callback,
	}

	defer func() {
		*buf = state.buf.toGo()
		state.release()
	}()

	stateHandle := cgo.NewHandle(state)
	defer stateHandle.Delete()

	flags |= InputTextFlagsCallbackResize

	return C.InputTextMultilineV(
		labelArg,
		(*C.char)(state.buf.ptr),
		C.xlong(len(*buf)+1),
		size.toC(),
		C.ImGuiInputTextFlags(flags),
		C.ImGuiInputTextCallback(C.generalInputTextCallback),
		unsafe.Pointer(&stateHandle),
	) == C.bool(true)
}

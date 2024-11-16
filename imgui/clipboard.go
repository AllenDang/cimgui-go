package imgui

// #include <stdlib.h>
// #include "wrapper.h"
// #include "typedefs.h"
// typedef char* (*get_clipboard_cb)(ImGuiContext* );
// typedef void (*set_clipboard_cb)(ImGuiContext*, char* );
// extern char* get_clipboard_callback(ImGuiContext* user_data);
// extern void set_clipboard_callback(ImGuiContext* user_data, char *text);
import "C"

import (
	"runtime/cgo"
	"unsafe"
)

//export get_clipboard_callback
func get_clipboard_callback(ctx *C.ImGuiContext) *C.char {
	h := (cgo.Handle)(ctx.PlatformIO.Platform_ClipboardUserData)
	helper := h.Value().(*clipboardHelper)
	handler := helper.handler
	str := handler.GetClipboard()

	buf := C.CString(str)
	if helper.buffer != nil {
		// imgui doesn't free the string, so we release the previous string every time we create a new one.
		// This does mean that we always keep the last clipboard contents alive.

		C.free(unsafe.Pointer(helper.buffer))
	}
	helper.buffer = buf

	return buf
}

//export set_clipboard_callback
func set_clipboard_callback(ctx *C.ImGuiContext, text *C.char) {
	h := (cgo.Handle)(ctx.PlatformIO.Platform_ClipboardUserData)
	helper := h.Value().(*clipboardHelper)
	handler := helper.handler

	str := C.GoString(text)
	handler.SetClipboard(str)
}

// clipboardHelper keeps state needed for the clipboard callbacks
type clipboardHelper struct {
	handler ClipboardHandler
	buffer  *C.char
}

// ClipboardHandler interfaces between imgui and the platforms clipboard
type ClipboardHandler interface {
	// GetClipboard should return the current contents of the platform clipboard
	GetClipboard() string
	// SetClipboard should replace the contents of the platform clipboard
	SetClipboard(s string)
}

func (io PlatformIO) SetClipboardHandler(handler ClipboardHandler) {
	rawIO, rawIOFin := io.Handle()
	defer rawIOFin()

	rawIO.Platform_GetClipboardTextFn = (C.get_clipboard_cb)(C.get_clipboard_callback)
	rawIO.Platform_SetClipboardTextFn = (C.set_clipboard_cb)(C.set_clipboard_callback)

	/*
		TODO: Turns out the ClipboardUserData isn't guaranteed to be nil for a fresh IO struct,
		so there's no good way to know if it's a handle we can delete, or just random data

		if rawIO.ClipboardUserData != nil {
			// cgo.NewHandle will keep the value alive until delete is called,
			// so if we're replacing an existing handler we delete it.
			h := (cgo.Handle)(rawIO.ClipboardUserData)
			h.Delete()
		}
	*/

	helper := &clipboardHelper{
		handler: handler,
	}

	userData := cgo.NewHandle(helper)
	rawIO.Platform_ClipboardUserData = unsafe.Pointer(userData)
}

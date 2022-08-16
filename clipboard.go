package cimgui

// #include "cimgui_wrapper.h"
import "C"

// Clipboard describes the access to the text clipboard of the window manager.
type Clipboard interface {
	// Text returns the current text from the clipboard, if available.
	Text() (string, error)
	// SetText sets the text as the current text on the clipboard.
	SetText(value string)
}

var clipboards = map[*C.ImGuiIO]Clipboard{}
var dropLastClipboardText = func() {}

//export ImGuiIO_GetClipboardText
func ImGuiIO_GetClipboardText(io *C.ImGuiIO) *C.char {
	dropLastClipboardText()

	board := clipboards[io]
	if board == nil {
		return nil
	}

	text, err := board.Text()
	if err != nil {
		return nil
	}
	textPtr, textFin := wrapString(text)
	dropLastClipboardText = func() {
		dropLastClipboardText = func() {}
		textFin()
	}
	return textPtr
}

//export ImGuiIO_SetClipboardText
func ImGuiIO_SetClipboardText(io *C.ImGuiIO, text *C.char) {
	dropLastClipboardText()

	board := clipboards[io]
	if board == nil {
		return
	}
	board.SetText(C.GoString(text))
}

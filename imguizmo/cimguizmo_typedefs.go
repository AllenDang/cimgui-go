// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imguizmo

import (
	"github.com/AllenDang/cimgui-go/datautils"
)

// #include <stdlib.h>
// #include <memory.h>
// #include "../imgui/extra_types.h"
// #include "cimguizmo_wrapper.h"
// #include "cimguizmo_typedefs.h"
import "C"

type Style struct {
	CData *C.Style
}

// Handle returns C version of Style and its finalizer func.
func (self *Style) Handle() (result *C.Style, fin func()) {
	return self.CData, func() {}
}

// C is like Handle but returns plain type instead of pointer.
func (self Style) C() (C.Style, func()) {
	result, fn := self.Handle()
	return *result, fn
}

// NewStyleFromC creates Style from its C pointer.
// SRC ~= *C.Style
func NewStyleFromC[SRC any](cvalue SRC) *Style {
	return &Style{CData: datautils.ConvertCTypes[*C.Style](cvalue)}
}
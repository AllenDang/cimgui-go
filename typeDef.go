package cimgui

// #include "cimgui_wrapper.h"
import "C"

type ImVec2 struct {
	X float32
	Y float32
}

func (v ImVec2) ToC() C.ImVec2 {
	return C.ImVec2{x: C.float(v.X), y: C.float(v.Y)}
}

type ImVec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func (v ImVec4) ToC() C.ImVec4 {
	return C.ImVec4{
		x: C.float(v.X),
		y: C.float(v.Y),
		z: C.float(v.Z),
		w: C.float(v.W),
	}
}

package cimgui

// #include "cimgui_wrapper.h"
import "C"

type ImWchar C.uint
type ImGuiID C.ImGuiID
type ImTextureID uintptr
type ImDrawIdx C.ImDrawIdx

type ImVec2 struct {
	X float32
	Y float32
}

func (i ImVec2) ToC() C.ImVec2 {
	return C.ImVec2{x: C.float(i.X), y: C.float(i.Y)}
}

func (vec *ImVec2) wrap() (out *C.ImVec2, finisher func()) {
	if vec != nil {
		out = &C.ImVec2{
			x: C.float(vec.X),
			y: C.float(vec.Y),
		}
		finisher = func() {
			vec.X = float32(out.x) // nolint: gotype
			vec.Y = float32(out.y) // nolint: gotype
		}
	} else {
		finisher = func() {}
	}
	return
}

type ImVec4 struct {
	X float32
	Y float32
	W float32
	Z float32
}

func (i ImVec4) ToC() C.ImVec4 {
	return C.ImVec4{x: C.float(i.X), y: C.float(i.Y), w: C.float(i.W), z: C.float(i.Z)}
}

func (vec *ImVec4) wrap() (out *C.ImVec4, finisher func()) {
	if vec != nil {
		out = &C.ImVec4{
			x: C.float(vec.X),
			y: C.float(vec.Y),
			w: C.float(vec.W),
			z: C.float(vec.Z),
		}
		finisher = func() {
			vec.X = float32(out.x) // nolint: gotype
			vec.Y = float32(out.y) // nolint: gotype
			vec.W = float32(out.w)
			vec.Z = float32(out.z)
		}
	} else {
		finisher = func() {}
	}
	return
}

type ImColor struct {
	Value ImVec4
}

func (i ImColor) ToC() C.ImColor {
	return C.ImColor{Value: i.Value.ToC()}
}

func (i *ImColor) wrap() (out *C.ImColor, finisher func()) {
	if i != nil {
		out = &C.ImColor{
			Value: i.Value.ToC(),
		}
		finisher = func() {
			i.Value.X = float32(out.Value.x)
			i.Value.Y = float32(out.Value.y)
			i.Value.W = float32(out.Value.w)
			i.Value.Z = float32(out.Value.z)
		}
	} else {
		finisher = func() {}
	}
	return
}

type ImGuiInputTextCallback func(*ImGuiInputTextCallbackData) int
type ImGuiSizeCallback func(*ImGuiSizeCallbackData)

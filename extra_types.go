package cimgui

// #include "cimgui_wrapper.h"
// #include "cimplot_wrapper.h"
import "C"

type (
	ImWchar                  C.uint
	ImGuiID                  C.ImGuiID
	ImTextureID              uintptr
	ImDrawIdx                C.ImDrawIdx
	ImGuiTableColumnIdx      C.ImGuiTableColumnIdx
	ImGuiTableDrawChannelIdx C.ImGuiTableDrawChannelIdx
)

var _ wrappableType[C.ImVec2, *Vec2] = &Vec2{}

type Vec2 struct {
	X float32
	Y float32
}

func NewVec2(x, y float32) Vec2 {
	return Vec2{X: x, Y: y}
}

func (i *Vec2) fromC(vec2 C.ImVec2) *Vec2 {
	*i = NewVec2(float32(vec2.x), float32(vec2.y))
	return i
}

func (i Vec2) toC() C.ImVec2 {
	return C.ImVec2{x: C.float(i.X), y: C.float(i.Y)}
}

var _ wrappableType[C.ImVec4, *ImVec4] = &ImVec4{}

type ImVec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func NewImVec4(r, g, b, a float32) ImVec4 {
	return ImVec4{
		X: r,
		Y: g,
		Z: b,
		W: a,
	}
}

func (i *ImVec4) fromC(vec4 C.ImVec4) *ImVec4 {
	*i = NewImVec4(float32(vec4.x), float32(vec4.y), float32(vec4.z), float32(vec4.w))
	return i
}

func (i ImVec4) toC() C.ImVec4 {
	return C.ImVec4{x: C.float(i.X), y: C.float(i.Y), z: C.float(i.Z), w: C.float(i.W)}
}

var _ wrappableType[C.ImColor, *ImColor] = &ImColor{}

type ImColor struct {
	Value ImVec4
}

func NewImColor(x, y, z, w float32) ImColor {
	return ImColor{
		Value: ImVec4{
			X: x,
			Y: y,
			Z: z,
			W: w,
		},
	}
}

func (i *ImColor) fromC(col C.ImColor) *ImColor {
	*i = NewImColor(float32(col.Value.x), float32(col.Value.y), float32(col.Value.z), float32(col.Value.w))
	return i
}

func (i ImColor) toC() C.ImColor {
	return C.ImColor{Value: i.Value.toC()}
}

var _ wrappableType[C.ImRect, *ImRect] = &ImRect{}

type ImRect struct {
	Min Vec2
	Max Vec2
}

func (i *ImRect) fromC(rect C.ImRect) *ImRect {
	out := &Vec2{}
	out.fromC(rect.Min)
	i.Min = *out

	out = &Vec2{}
	out.fromC(rect.Max)
	i.Max = *out

	return i
}

func (r *ImRect) toC() C.ImRect {
	return C.ImRect{Min: r.Min.toC(), Max: r.Max.toC()}
}

var _ wrappableType[C.ImPlotPoint, *ImPlotPoint] = &ImPlotPoint{}

type ImPlotPoint struct {
	X float64
	Y float64
}

func NewImPlotPoint(x, y float64) ImPlotPoint {
	return ImPlotPoint{X: x, Y: y}
}

func (i *ImPlotPoint) fromC(p C.ImPlotPoint) *ImPlotPoint {
	*i = NewImPlotPoint(float64(p.x), float64(p.y))
	return i
}

func (p ImPlotPoint) toC() C.ImPlotPoint {
	return C.ImPlotPoint{x: C.double(p.X), y: C.double(p.Y)}
}

// wrappableType represents a GO type that can be converted into a C value
// and back - into a GO value.
// CTYPE represents the equivalent C type of self.
// self is the type wrappableType applies to - TODO - figure out if it can be ommited :-)
// intentional values:
// - CTYPE is e.g. C.ImVec2, C.ImColor e.t.c.
// - self is a pointer type (e.g. *Vec2, ImColor)
type wrappableType[CTYPE any, self any] interface {
	// toC converts self into CTYPE
	toC() CTYPE
	// fromC converts takes CTYPE, converts it into self,
	// applies to receiver and returns it.
	fromC(CTYPE) self
}

// wrap takes a variable of one of the types defined in this file
// and returns a pointer to its C equivalend as well as a "finisher" func.
// This finisher func should be called after doing any operations
// on C pointer in order to apply the changes back to the original GO variable.
func wrap[CTYPE any, self any](in wrappableType[CTYPE, self]) (cPtr *CTYPE, finisher func()) {
	if in != nil {
		c := in.toC()
		cPtr = &c

		finisher = func() {
			in.fromC(*cPtr)
		}
	} else {
		finisher = func() {}
	}
	return cPtr, finisher
}

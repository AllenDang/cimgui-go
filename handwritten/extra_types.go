package cimgui

// #include "cimgui_wrapper.h"
// #include "cimplot_wrapper.h"
import "C"

import "image/color"

type ImWchar C.uint
type ImGuiID C.ImGuiID
type ImTextureID uintptr
type ImDrawIdx C.ImDrawIdx
type ImGuiTableColumnIdx C.ImGuiTableColumnIdx
type ImGuiTableDrawChannelIdx C.ImGuiTableDrawChannelIdx

type ImVec2 struct {
	X float32
	Y float32
}

func NewImVec2(x, y float32) ImVec2 {
	return ImVec2{X: x, Y: y}
}

func newImVec2FromC(vec2 C.ImVec2) ImVec2 {
	return NewImVec2(float32(vec2.x), float32(vec2.y))
}

func (i ImVec2) toC() C.ImVec2 {
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

func (vec ImVec2) Plus(other ImVec2) ImVec2 {
	return ImVec2{X: vec.X + other.X, Y: vec.Y + other.Y}
}

func (vec ImVec2) Minus(other ImVec2) ImVec2 {
	return ImVec2{X: vec.X - other.X, Y: vec.Y - other.Y}
}

func (vec ImVec2) Times(t float32) ImVec2 {
	return ImVec2{X: vec.X * t, Y: vec.Y * t}
}

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

func newImVec4FromCPtr(vec4 *C.ImVec4) ImVec4 {
	return NewImVec4(float32(vec4.x), float32(vec4.y), float32(vec4.z), float32(vec4.w))
}

func newImVec4FromC(vec4 C.ImVec4) ImVec4 {
	return NewImVec4(float32(vec4.x), float32(vec4.y), float32(vec4.z), float32(vec4.w))
}

func (i ImVec4) toC() C.ImVec4 {
	return C.ImVec4{x: C.float(i.X), y: C.float(i.Y), z: C.float(i.Z), w: C.float(i.W)}
}

func (vec *ImVec4) wrap() (out *C.ImVec4, finisher func()) {
	if vec != nil {
		out = &C.ImVec4{
			x: C.float(vec.X),
			y: C.float(vec.Y),
			z: C.float(vec.Z),
			w: C.float(vec.W),
		}
		finisher = func() {
			vec.X = float32(out.x) // nolint: gotype
			vec.Y = float32(out.y) // nolint: gotype
			vec.Z = float32(out.z)
			vec.W = float32(out.w)
		}
	} else {
		finisher = func() {}
	}
	return
}

func (vec ImVec4) Plus(other ImVec4) ImVec4 {
	return ImVec4{
		X: vec.X + other.X,
		Y: vec.Y + other.Y,
		Z: vec.Z + other.Z,
		W: vec.W + other.W,
	}
}

func (vec ImVec4) Minus(other ImVec4) ImVec4 {
	return ImVec4{
		X: vec.X - other.X,
		Y: vec.Y - other.Y,
		Z: vec.Z - other.Z,
		W: vec.W - other.W,
	}
}

func (vec ImVec4) Times(t float32) ImVec4 {
	return ImVec4{
		X: vec.X * t,
		Y: vec.Y * t,
		Z: vec.Z * t,
		W: vec.W * t,
	}
}

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

func NewImColorFromPacked(v uint32) ImColor {
	return NewImColor(
		float32(0xff&(v>>0))/255,
		float32(0xff&(v>>8))/255,
		float32(0xff&(v>>16))/255,
		float32(0xff&(v>>24))/255,
	)
}

func NewImColorFromColor(c color.Color) ImColor {
	nc := color.NRGBAModel.Convert(c).(color.NRGBA)
	return NewImColor(
		float32(nc.R)/0xffff,
		float32(nc.G)/0xffff,
		float32(nc.B)/0xffff,
		float32(nc.A)/0xffff,
	)
}

func (i ImColor) toC() C.ImColor {
	return C.ImColor{Value: i.Value.toC()}
}

func (i *ImColor) wrap() (out *C.ImColor, finisher func()) {
	if i != nil {
		out = &C.ImColor{
			Value: i.Value.toC(),
		}
		finisher = func() {
			i.Value.X = float32(out.Value.x)
			i.Value.Y = float32(out.Value.y)
			i.Value.Z = float32(out.Value.z)
			i.Value.W = float32(out.Value.w)
		}
	} else {
		finisher = func() {}
	}
	return
}

func colorComponent(v float32) uint8 {
	r := int(v*255 + 0.5)
	if r < 0 {
		return 0
	}
	if r > 255 {
		return 255
	}
	return uint8(r)
}

func (i ImColor) Pack() uint32 {
	return uint32(colorComponent(i.Value.X))<<0 |
		uint32(colorComponent(i.Value.Y))<<8 |
		uint32(colorComponent(i.Value.Z))<<16 |
		uint32(colorComponent(i.Value.W))<<24
}

func (i ImColor) Color() color.Color {
	return color.NRGBA{
		R: colorComponent(i.Value.X),
		G: colorComponent(i.Value.Y),
		B: colorComponent(i.Value.Z),
		A: colorComponent(i.Value.W),
	}
}

type ImRect struct {
	Min ImVec2
	Max ImVec2
}

func newImRectFromC(rect C.ImRect) ImRect {
	return ImRect{
		Min: newImVec2FromC(rect.Min),
		Max: newImVec2FromC(rect.Max),
	}
}

func (r *ImRect) toC() C.ImRect {
	return C.ImRect{Min: r.Min.toC(), Max: r.Max.toC()}
}

type ImPlotPoint struct {
	X float64
	Y float64
}

func NewImPlotPoint(x, y float64) ImPlotPoint {
	return ImPlotPoint{X: x, Y: y}
}

func newImPlotPointFromC(p C.ImPlotPoint) ImPlotPoint {
	return NewImPlotPoint(float64(p.x), float64(p.y))
}

func (p ImPlotPoint) toC() C.ImPlotPoint {
	return C.ImPlotPoint{x: C.double(p.X), y: C.double(p.Y)}
}

func (p *ImPlotPoint) wrap() (out *C.ImPlotPoint, finisher func()) {
	if p != nil {
		out = &C.ImPlotPoint{
			x: C.double(p.X),
			y: C.double(p.Y),
		}

		finisher = func() {
			p.X = float64(out.x)
			p.Y = float64(out.y)
		}
	} else {
		finisher = func() {}
	}

	return
}

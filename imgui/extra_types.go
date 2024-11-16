package imgui

// #include "wrapper.h"
// #include "extra_types.h"
import "C"

import (
	"image/color"
	"unsafe"

	"github.com/AllenDang/cimgui-go/internal"
)

type (
	Wchar C.uint
)

var _ internal.WrappableType[C.ImVec2, *Vec2] = &Vec2{}

type Vec2 struct {
	X float32
	Y float32
}

func NewVec2(x, y float32) Vec2 {
	return Vec2{X: x, Y: y}
}

// vec2Any is ~C.ImVec2 and will be force coerted!
func (i *Vec2) FromC(vec2Any unsafe.Pointer) *Vec2 {
	vec2 := *(*C.ImVec2)(vec2Any)
	*i = NewVec2(float32(vec2.x), float32(vec2.y))
	return i
}

func (i Vec2) ToC() C.ImVec2 {
	return C.ImVec2{x: C.float(i.X), y: C.float(i.Y)}
}

var _ internal.WrappableType[C.ImVec4, *Vec4] = &Vec4{}

type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

func NewVec4(r, g, b, a float32) Vec4 {
	return Vec4{
		X: r,
		Y: g,
		Z: b,
		W: a,
	}
}

// vec4Any is ~C.ImVec4 and will be force coerted!
func (i *Vec4) FromC(vec4Any unsafe.Pointer) *Vec4 {
	vec4 := *(*C.ImVec4)(vec4Any)
	*i = NewVec4(float32(vec4.x), float32(vec4.y), float32(vec4.z), float32(vec4.w))
	return i
}

func (i Vec4) ToC() C.ImVec4 {
	return C.ImVec4{x: C.float(i.X), y: C.float(i.Y), z: C.float(i.Z), w: C.float(i.W)}
}

var _ internal.WrappableType[C.ImColor, *Color] = &Color{}

type Color struct {
	FieldValue Vec4
}

func NewColor(r, g, b, a float32) Color {
	return Color{
		FieldValue: Vec4{
			X: r,
			Y: g,
			Z: b,
			W: a,
		},
	}
}

func NewColorFromPacked(v uint32) Color {
	return NewColor(
		float32(0xff&(v>>0))/255,
		float32(0xff&(v>>8))/255,
		float32(0xff&(v>>16))/255,
		float32(0xff&(v>>24))/255,
	)
}

func NewColorFromColor(c color.Color) Color {
	nc := color.NRGBAModel.Convert(c).(color.NRGBA)
	return NewColor(
		float32(nc.R)/0xffff,
		float32(nc.G)/0xffff,
		float32(nc.B)/0xffff,
		float32(nc.A)/0xffff,
	)
}

// colAny i ~C.ImColor and will be force coerted!
func (i *Color) FromC(colAny unsafe.Pointer) *Color {
	col := *(*C.ImColor)(colAny)
	*i = NewColor(float32(col.Value.x), float32(col.Value.y), float32(col.Value.z), float32(col.Value.w))
	return i
}

func (i Color) ToC() C.ImColor {
	return C.ImColor{Value: i.FieldValue.ToC()}
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

func (i Color) Pack() uint32 {
	return uint32(colorComponent(i.FieldValue.X))<<0 |
		uint32(colorComponent(i.FieldValue.Y))<<8 |
		uint32(colorComponent(i.FieldValue.Z))<<16 |
		uint32(colorComponent(i.FieldValue.W))<<24
}

func (i Color) Color() color.Color {
	return color.NRGBA{
		R: colorComponent(i.FieldValue.X),
		G: colorComponent(i.FieldValue.Y),
		B: colorComponent(i.FieldValue.Z),
		A: colorComponent(i.FieldValue.W),
	}
}

var _ internal.WrappableType[C.ImRect, *Rect] = &Rect{}

type Rect struct {
	Min Vec2
	Max Vec2
}

// rectAny is ~C.ImRect and will be force coerted!
func (i *Rect) FromC(rectAny unsafe.Pointer) *Rect {
	rect := *(*C.ImRect)(rectAny)
	out := &Vec2{}
	out.FromC(unsafe.Pointer(&rect.Min))
	i.Min = *out

	out = &Vec2{}
	out.FromC(unsafe.Pointer(&rect.Max))
	i.Max = *out

	return i
}

func (r *Rect) ToC() C.ImRect {
	return C.ImRect{Min: r.Min.ToC(), Max: r.Max.ToC()}
}

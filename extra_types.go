package imgui

// #include "cimgui_wrapper.h"
// #include "cimplot_wrapper.h"
// #include "extra_types.h"
import "C"

import (
	"image/color"
	"runtime"
	"time"
)

type (
	Wchar C.uint
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

var _ wrappableType[C.ImVec4, *Vec4] = &Vec4{}

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

func (i *Vec4) fromC(vec4 C.ImVec4) *Vec4 {
	*i = NewVec4(float32(vec4.x), float32(vec4.y), float32(vec4.z), float32(vec4.w))
	return i
}

func (i Vec4) toC() C.ImVec4 {
	return C.ImVec4{x: C.float(i.X), y: C.float(i.Y), z: C.float(i.Z), w: C.float(i.W)}
}

var _ wrappableType[C.ImColor, *Color] = &Color{}

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

func (i *Color) fromC(col C.ImColor) *Color {
	*i = NewColor(float32(col.Value.x), float32(col.Value.y), float32(col.Value.z), float32(col.Value.w))
	return i
}

func (i Color) toC() C.ImColor {
	return C.ImColor{Value: i.FieldValue.toC()}
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

var _ wrappableType[C.ImRect, *Rect] = &Rect{}

type Rect struct {
	Min Vec2
	Max Vec2
}

func (i *Rect) fromC(rect C.ImRect) *Rect {
	out := &Vec2{}
	out.fromC(rect.Min)
	i.Min = *out

	out = &Vec2{}
	out.fromC(rect.Max)
	i.Max = *out

	return i
}

func (r *Rect) toC() C.ImRect {
	return C.ImRect{Min: r.Min.toC(), Max: r.Max.toC()}
}

var _ wrappableType[C.ImPlotPoint, *PlotPoint] = &PlotPoint{}

type PlotPoint struct {
	X float64
	Y float64
}

func NewPlotPoint(x, y float64) PlotPoint {
	return PlotPoint{X: x, Y: y}
}

func (i *PlotPoint) fromC(p C.ImPlotPoint) *PlotPoint {
	*i = NewPlotPoint(float64(p.x), float64(p.y))
	return i
}

func (p PlotPoint) toC() C.ImPlotPoint {
	return C.ImPlotPoint{x: C.double(p.X), y: C.double(p.Y)}
}

type PlotTime struct {
	S       int // second part
	FieldUs int // microsecond part
}

func NewPlotTime(t time.Time) PlotTime {
	ts := t.UnixMicro()
	return PlotTime{
		S:       int(ts / 1e6),
		FieldUs: int(ts % 1e6),
	}
}

func (i PlotTime) Time() time.Time {
	return time.Unix(int64(i.S), int64(i.FieldUs)*int64(time.Microsecond))
}

func (i *PlotTime) fromC(p C.ImPlotTime) *PlotTime {
	*i = PlotTime{int(p.S), int(p.Us)}
	return i
}

func (p PlotTime) toC() C.ImPlotTime {
	return C.ImPlotTime{S: C.xlong(p.S), Us: C.int(p.FieldUs)}
}

type Vector[T any] struct {
	Size     int
	Capacity int
	Data     T
	pinner   *runtime.Pinner
}

func newVectorFromC[T any](size, capacity C.int, data T) Vector[T] {
	return Vector[T]{Size: int(size), Capacity: int(capacity), Data: data, pinner: &runtime.Pinner{}}
}

// wrappableType represents a GO type that can be converted into a C value
// and back - into a GO value.
// CTYPE represents the equivalent C type of self.
// self is the type wrappableType applies to - TODO - figure out if it can be ommited :-)
// intentional values:
// - CTYPE is e.g. C.ImVec2, C.ImColor e.t.c.
// - self is a pointer type (e.g. *Vec2, Color)
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

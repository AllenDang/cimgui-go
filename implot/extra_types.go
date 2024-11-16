package implot

// #include "wrapper.h"
// #include "../imgui/extra_types.h"
import "C"

import (
	"time"
	"unsafe"

	"github.com/AllenDang/cimgui-go/internal"
)

var _ internal.WrappableType[C.ImPlotPoint, *PlotPoint] = &PlotPoint{}

type PlotPoint struct {
	X float64
	Y float64
}

func NewPlotPoint(x, y float64) PlotPoint {
	return PlotPoint{X: x, Y: y}
}

// pAny is ~C.ImPlotPoint and will be free converted!
func (i *PlotPoint) FromC(pAny unsafe.Pointer) *PlotPoint {
	p := *(*C.ImPlotPoint)(pAny)
	*i = NewPlotPoint(float64(p.x), float64(p.y))
	return i
}

func (p PlotPoint) ToC() C.ImPlotPoint {
	return C.ImPlotPoint{x: C.double(p.X), y: C.double(p.Y)}
}

type PlotTime struct {
	Seconds int // second part
	FieldUs int // microsecond part
}

func NewPlotTime(t time.Time) PlotTime {
	ts := t.UnixMicro()
	return PlotTime{
		Seconds: int(ts / 1e6),
		FieldUs: int(ts % 1e6),
	}
}

func (i PlotTime) Time() time.Time {
	return time.Unix(int64(i.Seconds), int64(i.FieldUs)*int64(time.Microsecond))
}

// pAny is ~C.ImPlotTime and will be free converted!
func (i *PlotTime) FromC(pAny unsafe.Pointer) *PlotTime {
	p := *(*C.ImPlotTime)(pAny)
	*i = PlotTime{int(p.S), int(p.Us)}
	return i
}

func (p PlotTime) ToC() C.ImPlotTime {
	return C.ImPlotTime{S: C.xlong(p.Seconds), Us: C.int(p.FieldUs)}
}

var _ internal.WrappableType[C.struct_tm, *Tm] = &Tm{}

// Tm is an implemenation of tm C type
// ref: https://en.cppreference.com/w/c/chrono/tm
type Tm struct {
	Seconds int32
	Minutes int32
	Hours   int32
	MDay    int32
	Month   int32
	Year    int32
	WDay    int32
	YDay    int32
	IsDST   int32
}

func NewTm(S, M, H, MD, MO, Y, WD, YD, DST int32) Tm {
	return Tm{
		Seconds: S,
		Minutes: M,
		Hours:   H,
		MDay:    MD,
		Month:   MO,
		Year:    Y,
		WDay:    WD,
		YDay:    YD,
		IsDST:   DST,
	}
}

func (i Tm) ToC() C.struct_tm {
	return C.struct_tm{
		tm_sec:   C.int(i.Seconds),
		tm_min:   C.int(i.Minutes),
		tm_hour:  C.int(i.Hours),
		tm_mday:  C.int(i.MDay),
		tm_mon:   C.int(i.Month),
		tm_year:  C.int(i.Year),
		tm_wday:  C.int(i.WDay),
		tm_yday:  C.int(i.YDay),
		tm_isdst: C.int(i.IsDST),
	}
}

func (i *Tm) FromC(pAny unsafe.Pointer) *Tm {
	p := *(*C.struct_tm)(pAny)
	*i = NewTm(int32(p.tm_sec), int32(p.tm_min), int32(p.tm_hour), int32(p.tm_mday), int32(p.tm_mon), int32(p.tm_year), int32(p.tm_wday), int32(p.tm_yday), int32(p.tm_isdst))
	return i
}

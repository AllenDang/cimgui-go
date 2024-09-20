package implot

// #include "cimplot_wrapper.h"
// #include "../imgui/extra_types.h"
import "C"

import (
	"image/color"
	"runtime"
	"time"
)

var _ imgui.WrappableType[C.ImPlotPoint, *PlotPoint] = &PlotPoint{}

type PlotPoint struct {
	X float64
	Y float64
}

func NewPlotPoint(x, y float64) PlotPoint {
	return PlotPoint{X: x, Y: y}
}

func (i *PlotPoint) FromC(p C.ImPlotPoint) *PlotPoint {
	*i = NewPlotPoint(float64(p.x), float64(p.y))
	return i
}

func (p PlotPoint) ToC() C.ImPlotPoint {
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

func (i *PlotTime) FromC(p C.ImPlotTime) *PlotTime {
	*i = PlotTime{int(p.S), int(p.Us)}
	return i
}

func (p PlotTime) ToC() C.ImPlotTime {
	return C.ImPlotTime{S: C.xlong(p.S), Us: C.int(p.FieldUs)}
}

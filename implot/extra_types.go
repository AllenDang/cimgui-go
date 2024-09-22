package implot

// #include "cimplot_wrapper.h"
// #include "../imgui/extra_types.h"
import "C"

import (
	"time"
	"unsafe"

	"github.com/AllenDang/cimgui-go/datautils"
)

var _ datautils.WrappableType[C.ImPlotPoint, *PlotPoint] = &PlotPoint{}

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

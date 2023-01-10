package main

import (
	"fmt"
	"image"

	"github.com/AllenDang/cimgui-go"
)

var (
	showDemoWindow bool
	value1         int32
	value2         int32
	value3         int32
	values         [2]*int32 = [2]*int32{&value1, &value2}
	content        string    = "Let me try"
	r              float32
	g              float32
	b              float32
	a              float32
	color4         [4]*float32 = [4]*float32{&r, &g, &b, &a}
	selected       bool
	window         cimgui.GLFWwindow
	img            *image.RGBA
	texture        *cimgui.Texture
	barValues      []int64
)

func callback(data cimgui.InputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

func showWidgetsDemo() {
	if showDemoWindow {
		cimgui.ShowDemoWindowV(&showDemoWindow)
	}

	cimgui.SetNextWindowSizeV(cimgui.NewVec2(300, 300), cimgui.CondOnce)
	cimgui.Begin("Window 1")
	if cimgui.ButtonV("Click Me", cimgui.NewVec2(80, 20)) {
		w, h := window.DisplaySize()
		fmt.Println(w, h)
	}
	cimgui.TextUnformatted("Unformatted text")
	cimgui.Checkbox("Show demo window", &showDemoWindow)
	if cimgui.BeginCombo("Combo", "Combo preview") {
		cimgui.SelectableBoolPtr("Item 1", &selected)
		cimgui.SelectableBool("Item 2")
		cimgui.SelectableBool("Item 3")
		cimgui.EndCombo()
	}

	if cimgui.RadioButtonBool("Radio button1", selected) {
		selected = true
	}

	cimgui.SameLine()

	if cimgui.RadioButtonBool("Radio button2", !selected) {
		selected = false
	}

	cimgui.InputTextWithHint("Name", "write your name here", &content, 0, callback)
	cimgui.Text(content)
	cimgui.SliderInt("Slider int", &value3, 0, 100)
	cimgui.DragInt("Drag int", &value1)
	cimgui.DragInt2("Drag int2", values)
	cimgui.ColorEdit4("Color Edit3", color4)
	cimgui.End()
}

func showPictureLoadingDemo() {
	// demo of showing a picture
	basePos := cimgui.MainViewport().Pos()
	cimgui.SetNextWindowPosV(cimgui.NewVec2(basePos.X+60, 600), cimgui.CondOnce, cimgui.NewVec2(0, 0))
	cimgui.Begin("Image")
	cimgui.Text(fmt.Sprintf("pointer = %v", texture.ID()))
	cimgui.ImageV(texture.ID(), cimgui.NewVec2(float32(texture.Width), float32(texture.Height)), cimgui.NewVec2(0, 0), cimgui.NewVec2(1, 1), cimgui.NewVec4(1, 1, 1, 1), cimgui.NewVec4(0, 0, 0, 0))
	cimgui.End()
}

func showImPlotDemo() {
	basePos := cimgui.MainViewport().Pos()
	cimgui.SetNextWindowPosV(cimgui.NewVec2(basePos.X+400, basePos.Y+60), cimgui.CondOnce, cimgui.NewVec2(0, 0))
	cimgui.SetNextWindowSizeV(cimgui.NewVec2(500, 300), cimgui.CondOnce)
	cimgui.Begin("Plot window")
	if cimgui.PlotBeginPlotV("Plot", cimgui.NewVec2(-1, -1), 0) {
		cimgui.PlotPlotBarsS64PtrInt("Bar", barValues, int32(len(barValues)))
		cimgui.PlotPlotLineS64PtrInt("Line", barValues, int32(len(barValues)))
		cimgui.PlotEndPlot()
	}
	cimgui.End()
}

func afterCreateContext() {
	texture = cimgui.NewTextureFromRgba(img)
	cimgui.PlotCreateContext()
}

func loop() {
	showWidgetsDemo()
	showPictureLoadingDemo()
	showImPlotDemo()
}

func beforeDestroyContext() {
	cimgui.PlotDestroyContext()
}

func main() {
	var err error
	img, err = cimgui.LoadImage("./test.jpeg")
	if err != nil {
		panic("Failed to load test.jpeg")
	}

	for i := 0; i < 10; i++ {
		barValues = append(barValues, int64(i+1))
	}

	cimgui.SetAfterCreateContextHook(afterCreateContext)
	cimgui.SetBeforeDestroyContextHook(beforeDestroyContext)

	cimgui.SetBgColor(cimgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	window = cimgui.CreateGlfwWindow("Hello from cimgui-go", 1200, 900, 0)

	window.Run(loop)
}

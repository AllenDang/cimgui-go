package main

import (
	"fmt"
	"image"
	"runtime"

	imgui "github.com/AllenDang/cimgui-go"
)

var (
	showDemoWindow bool
	value1         int32
	value2         int32
	value3         int32
	values         [2]int32 = [2]int32{value1, value2}
	content        string   = "Let me try"
	r              float32
	g              float32
	b              float32
	a              float32
	color4         [4]float32 = [4]float32{r, g, b, a}
	selected       bool
	backend        imgui.Backend[imgui.GLFWWindowFlags]
	img            *image.RGBA
	texture        *imgui.Texture
	barValues      []int64
)

func callback(data imgui.InputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

func showWidgetsDemo() {
	if showDemoWindow {
		imgui.ShowDemoWindowV(&showDemoWindow)
	}

	imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)
	imgui.Begin("Window 1")
	if imgui.ButtonV("Click Me", imgui.NewVec2(80, 20)) {
		w, h := backend.DisplaySize()
		fmt.Println(w, h)
	}
	imgui.TextUnformatted("Unformatted text")
	imgui.Checkbox("Show demo window", &showDemoWindow)
	if imgui.BeginCombo("Combo", "Combo preview") {
		imgui.SelectableBoolPtr("Item 1", &selected)
		imgui.SelectableBool("Item 2")
		imgui.SelectableBool("Item 3")
		imgui.EndCombo()
	}

	if imgui.RadioButtonBool("Radio button1", selected) {
		selected = true
	}

	imgui.SameLine()

	if imgui.RadioButtonBool("Radio button2", !selected) {
		selected = false
	}

	imgui.InputTextWithHint("Name", "write your name here", &content, 0, callback)
	imgui.Text(content)
	imgui.SliderInt("Slider int", &value3, 0, 100)
	imgui.DragInt("Drag int", &value1)
	imgui.DragInt2("Drag int2", &values)
	value1 = values[0]
	imgui.ColorEdit4("Color Edit3", &color4)
	imgui.End()
}

func showPictureLoadingDemo() {
	// demo of showing a picture
	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+60, 600), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.Begin("Image")
	imgui.Text(fmt.Sprintf("pointer = %v", texture.ID))
	imgui.ImageV(texture.ID, imgui.NewVec2(float32(texture.Width), float32(texture.Height)), imgui.NewVec2(0, 0), imgui.NewVec2(1, 1), imgui.NewVec4(1, 1, 1, 1), imgui.NewVec4(0, 0, 0, 0))
	imgui.End()
}

func showImPlotDemo() {
	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+400, basePos.Y+60), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.SetNextWindowSizeV(imgui.NewVec2(500, 300), imgui.CondOnce)
	imgui.Begin("Plot window")
	if imgui.PlotBeginPlotV("Plot", imgui.NewVec2(-1, -1), 0) {
		imgui.PlotPlotBarsS64PtrInt("Bar", barValues, int32(len(barValues)))
		imgui.PlotPlotLineS64PtrInt("Line", barValues, int32(len(barValues)))
		imgui.PlotEndPlot()
	}
	imgui.End()
}

func afterCreateContext() {
	texture = imgui.NewTextureFromRgba(img)
	imgui.PlotCreateContext()
}

func loop() {
	showWidgetsDemo()
	showPictureLoadingDemo()
	showImPlotDemo()
}

func beforeDestroyContext() {
	imgui.PlotDestroyContext()
}

func init() {
	runtime.LockOSThread()
}

func main() {
	var err error
	img, err = imgui.LoadImage("./test.jpeg")
	if err != nil {
		panic("Failed to load test.jpeg")
	}

	for i := 0; i < 10; i++ {
		barValues = append(barValues, int64(i+1))
	}

	backend, _ = imgui.CreateBackend(imgui.NewGLFWBackend())
	backend.SetAfterCreateContextHook(afterCreateContext)
	backend.SetBeforeDestroyContextHook(beforeDestroyContext)

	backend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	backend.CreateWindow("Hello from cimgui-go", 1200, 900)

	backend.SetDropCallback(func(p []string) {
		fmt.Printf("drop triggered: %v", p)
	})

	backend.SetCloseCallback(func(b imgui.Backend[imgui.GLFWWindowFlags]) {
		fmt.Println("window is closing")
	})

	backend.SetIcons(img)

	backend.Run(loop)
}

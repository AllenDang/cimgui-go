package main

import (
	"fmt"
	"image"
	"runtime"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
	_ "github.com/AllenDang/cimgui-go/immarkdown"
	_ "github.com/AllenDang/cimgui-go/imnodes"
	"github.com/AllenDang/cimgui-go/implot"
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
	currentBackend backend.Backend[glfwbackend.GLFWWindowFlags]
	img            *image.RGBA
	texture        *backend.Texture
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
		w, h := currentBackend.DisplaySize()
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
	if implot.PlotBeginPlotV("Plot", imgui.NewVec2(-1, -1), 0) {
		implot.PlotPlotBarsS64PtrInt("Bar", barValues, int32(len(barValues)))
		implot.PlotPlotLineS64PtrInt("Line", barValues, int32(len(barValues)))
		implot.PlotEndPlot()
	}
	imgui.End()
}

func afterCreateContext() {
	texture = backend.NewTextureFromRgba(img)
	implot.PlotCreateContext()
}

func loop() {
	showWidgetsDemo()
	showPictureLoadingDemo()
	showImPlotDemo()
}

func beforeDestroyContext() {
	implot.PlotDestroyContext()
}

func init() {
	runtime.LockOSThread()
}

func main() {
	var err error
	img, err = backend.LoadImage("../assets/test.jpeg")
	if err != nil {
		panic("Failed to load test.jpeg")
	}

	for i := 0; i < 10; i++ {
		barValues = append(barValues, int64(i+1))
	}

	currentBackend, _ = backend.CreateBackend(glfwbackend.NewGLFWBackend())
	currentBackend.SetAfterCreateContextHook(afterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(beforeDestroyContext)

	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	currentBackend.CreateWindow("Hello from cimgui-go", 1200, 900)

	currentBackend.SetDropCallback(func(p []string) {
		fmt.Printf("drop triggered: %v", p)
	})

	currentBackend.SetCloseCallback(func(b backend.Backend[glfwbackend.GLFWWindowFlags]) {
		fmt.Println("window is closing")
	})

	currentBackend.SetIcons(img)

	currentBackend.Run(loop)
}

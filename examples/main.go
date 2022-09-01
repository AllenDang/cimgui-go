package main

import (
	"fmt"

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
	texture        *cimgui.Texture
	barValues      []int64
)

func callback(data cimgui.ImGuiInputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

func showWidgetsDemo() {
	if showDemoWindow {
		cimgui.ShowDemoWindow(&showDemoWindow)
	}

	cimgui.SetNextWindowSize(cimgui.NewImVec2(300, 300), cimgui.ImGuiCond_Once)
	cimgui.Begin("Window 1", nil, 0)
	if cimgui.Button("Click Me", cimgui.NewImVec2(80, 20)) {
		w, h := window.DisplaySize()
		fmt.Println(w, h)
	}
	cimgui.TextUnformatted("Unformated text")
	cimgui.Checkbox("Show demo window", &showDemoWindow)
	if cimgui.BeginCombo("Combo", "Combo preview", cimgui.ImGuiComboFlags_HeightLarge) {
		cimgui.Selectable_BoolPtr("Item 1", &selected, 0, cimgui.NewImVec2(100, 20))
		cimgui.Selectable_Bool("Item 2", false, 0, cimgui.NewImVec2(100, 20))
		cimgui.Selectable_Bool("Item 3", false, 0, cimgui.NewImVec2(100, 20))
		cimgui.EndCombo()
	}

	if cimgui.RadioButton_Bool("Radio button1", selected) {
		selected = true
	}

	cimgui.SameLine(0, 0)

	if cimgui.RadioButton_Bool("Radio button2", !selected) {
		selected = false
	}

	cimgui.InputTextWithHint("Name", "write your name here", &content, 0, callback)
	cimgui.Text(content)
	cimgui.SliderInt("Slider int", &value3, 0, 100, "%d", 0)
	cimgui.DragInt("Drag int", &value1, 1, 0, 100, "%d", 0)
	cimgui.DragInt2("Drag int2", values, 1, 0, 100, "%d", 0)
	cimgui.ColorEdit4("Color Edit3", color4, 0)
	cimgui.End()
}

func showPictureLoadingDemo() {
	// demo of showing a picture
	basePos := cimgui.GetMainViewport().GetPos()
	cimgui.SetNextWindowPos(cimgui.NewImVec2(basePos.X+60, 600), cimgui.ImGuiCond_Appearing, cimgui.NewImVec2(0, 0))
	cimgui.Begin("Image", nil, 0)
	cimgui.Text(fmt.Sprintf("pointer = %v", texture.ID()))
	cimgui.Image(texture.ID(), cimgui.NewImVec2(float32(texture.Width), float32(texture.Height)), cimgui.NewImVec2(0, 0), cimgui.NewImVec2(1, 1), cimgui.NewImVec4(1, 1, 1, 1), cimgui.NewImVec4(0, 0, 0, 0))
	cimgui.End()
}

func showImPlotDemo() {
	basePos := cimgui.GetMainViewport().GetPos()
	cimgui.SetNextWindowPos(cimgui.NewImVec2(basePos.X+400, basePos.Y+60), cimgui.ImGuiCond_Appearing, cimgui.NewImVec2(0, 0))
	cimgui.SetNextWindowSize(cimgui.NewImVec2(500, 300), cimgui.ImGuiCond_Appearing)
	cimgui.Begin("Plot window", nil, 0)
	if cimgui.Plot_BeginPlot("Plot", cimgui.NewImVec2(-1, -1), 0) {
		cimgui.Plot_PlotBars_S64PtrInt("Bar", barValues, int32(len(barValues)), 0.67, 0, 0, 0, 8)
		cimgui.Plot_PlotLine_S64PtrInt("Line", barValues, int32(len(barValues)), 1, 0, 0, 0, 8)
		cimgui.Plot_EndPlot()
	}
	cimgui.End()
}

func afterCreateContext() {
	cimgui.Plot_CreateContext()
}

func loop() {
	showWidgetsDemo()
	showPictureLoadingDemo()
	showImPlotDemo()
}

func beforeDestoryContext() {
	cimgui.Plot_DestroyContext(0)
}

func main() {
	img, err := cimgui.LoadImage("./test.jpeg")
	if err != nil {
		panic("Failed to load gopher.png")
	}

	for i := 0; i < 10; i++ {
		barValues = append(barValues, int64(i+1))
	}

	cimgui.SetAfterCreateContextHook(afterCreateContext)
	cimgui.SetBeforeDestroyContextHook(beforeDestoryContext)

	window = cimgui.CreateGlfwWindow("Hello from cimgui-go", 1200, 900, 0)

	texture, err = cimgui.NewTextureFromRgba(img)
	if err != nil {
		panic("Failed to create texture")
	}

	window.Run(loop)
}

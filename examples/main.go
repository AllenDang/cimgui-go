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
)

func callback(data cimgui.ImGuiInputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

func loop() {
	if showDemoWindow {
		cimgui.ShowDemoWindow(&showDemoWindow)
	}

	cimgui.SetNextWindowSize(cimgui.NewImVec2(300, 300), cimgui.ImGuiCond_Once)
	cimgui.Begin("Window 1", nil, 0)
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

func main() {
	window := cimgui.CreateGlfwWindow("Hello from cimgui-go", 1200, 900, 0)

	window.Run(loop, nil, nil)
}

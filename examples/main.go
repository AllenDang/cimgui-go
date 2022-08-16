package main

import (
	"github.com/AllenDang/cimgui-go"
)

var (
	showDemoWindow bool = true
)

func loop() {
	cimgui.ShowDemoWindow(&showDemoWindow)
}

func main() {
	window := cimgui.CreateGlfwWindow("Hello from cimgui-go", 1200, 900)
	window.Run(loop)
}

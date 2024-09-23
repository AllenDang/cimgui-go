// This code is a modified version of examples/shapes from https://github.com/ebiten/ebiten
// Ebiten is licensed under the following licens:
//
//
// Copyright 2017 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/AllenDang/cimgui-go/backend"
	ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 640
	screenHeight = 480
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
		fmt.Println("I was clicked!!")
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
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+60, 400), imgui.CondOnce, imgui.NewVec2(0, 0))
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
	imgui.CreateContextV(imgui.NewFontAtlas())
	texture = backend.NewTextureFromRgba(img)
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

type Game struct {
	count  int
	cimgui *ebitenbackend.EbitenBackend
}

func (g *Game) Update() error {
	g.cimgui.BeginFrame()
	loop()
	g.cimgui.EndFrame()
	g.count++
	g.count %= 240
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cf := float32(g.count)
	vector.StrokeLine(screen, 100, 100, 300, 100, 1, color.RGBA{0xff, 0xff, 0xff, 0xff}, true)
	vector.StrokeLine(screen, 50, 150, 50, 350, 1, color.RGBA{0xff, 0xff, 0x00, 0xff}, true)
	vector.StrokeLine(screen, 50, 100+cf, 200+cf, 250, 4, color.RGBA{0x00, 0xff, 0xff, 0xff}, true)

	vector.DrawFilledRect(screen, 50+cf, 50+cf, 100+cf, 100+cf, color.RGBA{0x80, 0x80, 0x80, 0xc0}, true)
	vector.StrokeRect(screen, 300-cf, 50, 120, 120, 10+cf/4, color.RGBA{0x00, 0x80, 0x00, 0xff}, true)

	vector.DrawFilledCircle(screen, 400, 400, 100, color.RGBA{0x80, 0x00, 0x80, 0x80}, true)
	vector.StrokeCircle(screen, 400, 400, 10+cf, 10+cf/2, color.RGBA{0xff, 0x80, 0xff, 0xff}, true)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
	g.cimgui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	g.cimgui.Layout(outsideWidth, outsideHeight)
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Shapes (Ebitengine Demo)")
	ebitenBackend := ebitenbackend.NewEbitenBackend()
	_, err := backend.CreateBackend(ebitenBackend)
	if err != nil && !errors.Is(err, backend.CExposerError) {
		panic(err)
	}

	game := &Game{cimgui: ebitenBackend}

	img, err = backend.LoadImage("../assets/test.jpeg")
	if err != nil {
		panic("Failed to load test.jpeg")
	}

	for i := 0; i < 10; i++ {
		barValues = append(barValues, int64(i+1))
	}

	/*
		game.cimgui.SetAfterCreateContextHook(afterCreateContext)
		game.cimgui.SetBeforeDestroyContextHook(beforeDestroyContext)
	*/

	afterCreateContext()
	defer beforeDestroyContext()

	game.cimgui.CreateWindow("Hello, world!", 800, 600)

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

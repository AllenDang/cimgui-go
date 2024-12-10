package main

import (
	"fmt"
	"image/color"

	"github.com/AllenDang/cimgui-go/backend"
	ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	"github.com/AllenDang/cimgui-go/examples/common"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/implot"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var (
	currentBackend *ebitenbackend.EbitenBackend
	texture        *backend.Texture
)

func showPictureLoadingDemo() {
	// demo of showing a picture
	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+60, 600), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.Begin("Ebiten Active Image")
	imgui.Text(fmt.Sprintf("pointer = %v", texture.ID))
	imgui.ImageV(texture.ID, imgui.NewVec2(float32(texture.Width), float32(texture.Height)), imgui.NewVec2(0, 0), imgui.NewVec2(1, 1), imgui.NewVec4(1, 1, 1, 1), imgui.NewVec4(0, 0, 0, 0))
	imgui.End()
}

func afterCreateContext() {
	texture = &backend.Texture{
		ID:     currentBackend.CreateTextureFromGame(&Game{}, screenWidth, screenHeight),
		Width:  screenWidth,
		Height: screenHeight,
	}

	implot.CreateContext()
}

func loop() {
	imgui.ClearSizeCallbackPool()
	common.ShowWidgetsDemo()
	showPictureLoadingDemo()
	common.ShowImPlotDemo()
}

func beforeDestroyContext() {
	implot.DestroyContext()
}

func main() {
	common.Initialize()

	currentBackend = ebitenbackend.NewEbitenBackend()
	_, _ = backend.CreateBackend(currentBackend)
	currentBackend.SetAfterCreateContextHook(afterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(beforeDestroyContext)

	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	currentBackend.CreateWindow("Hello from cimgui-go", 1200, 900)

	// TODO: not implemented
	/*
		backend.SetDropCallback(func(p []string) {
			fmt.Printf("drop triggered: %v", p)
		})
	*/

	currentBackend.SetCloseCallback(func() {
		fmt.Println("window is closing")
	})

	currentBackend.Run(loop)
}

// code from ebiten/examples/shape (https://github.com/hajimehoshi/ebiten
// for details see LICENSE of ebiten

const (
	screenWidth, screenHeight = 800, 600
)

type Game struct {
	count int
}

func (g *Game) Update() error {
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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

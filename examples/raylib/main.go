package main

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/backend"
	raylibbackend "github.com/AllenDang/cimgui-go/backend/raylib-backend"
	"github.com/AllenDang/cimgui-go/examples/common"
	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var currentBackend *raylibbackend.RaylibBackend

func main() {
	common.Initialize()

	currentBackend = raylibbackend.NewRaylibBackend()
	backend.CreateBackend[raylibbackend.RaylibBackendFlags](currentBackend)
	currentBackend.SetAfterCreateContextHook(common.AfterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(common.BeforeDestroyContext)

	currentBackend.SetConfigFlags(
		raylibbackend.RaylibBackendFlagsVsyncHint,
		raylibbackend.RaylibBackendFlagsResizable,
		raylibbackend.RaylibBackendFlagsMSAA4X,
	)

	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	currentBackend.CreateWindow("Hello from cimgui-go (raylib)", 1200, 900)

	currentBackend.SetCloseCallback(func() {
		fmt.Println("window is closing")
	})

	currentBackend.SetIcons(common.Image())

	showCimgui := true
	count := 0

	currentBackend.SetBeforeImGuiRenderHook(func() {
		count++
		if count >= 240 {
			count = 0
		}
		cf := float32(count)

		rl.MatrixMode(rl.Projection)
		rl.LoadIdentity()
		w, h := float64(rl.GetScreenWidth()), float64(rl.GetScreenHeight())
		rl.Ortho(0, w, h, 0, 0, 1)
		rl.MatrixMode(rl.Modelview)
		rl.LoadIdentity()

		rl.DrawLineEx(rl.Vector2{X: 100, Y: 100}, rl.Vector2{X: 300, Y: 100}, 2, rl.White)
		rl.DrawLineEx(rl.Vector2{X: 50, Y: 150}, rl.Vector2{X: 50, Y: 350}, 2, rl.Yellow)
		rl.DrawLineEx(rl.Vector2{X: 50, Y: 100 + cf}, rl.Vector2{X: 200 + cf, Y: 250}, 5, rl.SkyBlue)

		rl.DrawRectangle(int32(50+cf), int32(50+cf), int32(100+cf), int32(100+cf), rl.NewColor(0x80, 0x80, 0x80, 0xC0))
		rl.DrawRectangleLinesEx(rl.Rectangle{X: 300 - cf, Y: 50, Width: 120, Height: 120}, 10+cf/4, rl.Green)

		rl.DrawCircle(int32(400), int32(400), 100, rl.NewColor(0x80, 0x00, 0x80, 0x80))
		rl.DrawCircleLines(int32(400), int32(400), 10+cf, rl.NewColor(0xFF, 0x80, 0xFF, 0xFF))

		rl.DrawText(
			fmt.Sprintf("FPS: %d  Enable CImGui [q]: %v", rl.GetFPS(), showCimgui),
			int32(10), int32(10), int32(20), rl.RayWhite,
		)
	})

	currentBackend.Run(func() {
		imgui.ClearSizeCallbackPool()

		if rl.IsKeyPressed(rl.KeyQ) {
			showCimgui = !showCimgui
		}

		if showCimgui {
			common.Loop()
		}
	})
}

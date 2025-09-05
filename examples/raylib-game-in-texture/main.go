package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/AllenDang/cimgui-go/backend"
	raylibbackend "github.com/AllenDang/cimgui-go/backend/raylib-backend"
	"github.com/AllenDang/cimgui-go/examples/common"
	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	texWidth  = 800
	texHeight = 600
)

var (
	currentBackend *raylibbackend.RaylibBackend
	rt             rl.RenderTexture2D
	frameCount     int
)

func afterCreateContext() {
	rt = rl.LoadRenderTexture(int32(texWidth), int32(texHeight))
	rl.SetTextureFilter(rt.Texture, rl.FilterPoint)
}

func beforeDestroyContext() {
	rl.UnloadRenderTexture(rt)
}

func loop() {
	imgui.ClearSizeCallbackPool()

	frameCount++
	if frameCount >= 360 {
		frameCount = 0
	}
	t := float32(frameCount)

	var (
		appFocused = rl.IsWindowFocused()
		winHovered bool
	)

	imgui.SetNextWindowPosV(imgui.NewVec2(60, 60), imgui.CondOnce, imgui.NewVec2(0, 0))

	imgui.Begin("Raylib RenderTexture Preview")

	// Track if window is hovered
	winHovered = imgui.IsWindowHovered()

	prep := currentBackend.PrepareImageRenderTexture(rt, &raylibbackend.ImageRTOptions{
		AlignCursor:          true,
		FlipV:                true,
		RequireAppFocus:      true,
		RequireWindowHovered: true,
	})

	rl.BeginTextureMode(rt)
	rl.ClearBackground(rl.NewColor(70, 74, 82, 255))
	rl.DrawLineEx(rl.NewVector2(40, 40), rl.NewVector2(40+t*0.5, 200), 3, rl.SkyBlue)
	rl.DrawRectangle(int32(100+t*0.4), 80, 160, 120, rl.NewColor(0x50, 0x90, 0xB0, 0xE0))
	rl.DrawRectangleLines(int32(90), int32(70), int32(180), int32(140), rl.Yellow)
	rl.DrawCircle(400, 260, 90+20*float32(math.Sin(float64(t*rl.Deg2rad))), rl.NewColor(0x80, 0x20, 0x90, 0xB0))
	rl.DrawCircleLines(400, 260, 30+10*float32(math.Cos(float64(t*rl.Deg2rad*2))), rl.NewColor(0xFF, 0x80, 0xFF, 0xFF))

	// If app focused + this window hovered + mouse over the image, draw a square at mapped texture coords
	if appFocused && winHovered && prep.Hovered {
		sz := int32(14)
		half := sz / 2
		mx := int32(prep.TexX)
		my := int32(prep.TexY)
		rl.DrawRectangle(mx-half, my-half, sz, sz, rl.NewColor(0x20, 0xE0, 0x80, 0xCC))
		rl.DrawRectangleLines(mx-half, my-half, sz, sz, rl.NewColor(0x00, 0xFF, 0xAA, 0xFF))

		// Collect pressed keys
		pressed := make([]string, 0, 32)
		add := func(name string, key int32) {
			if rl.IsKeyDown(key) {
				pressed = append(pressed, name)
			}
		}
		add("W", rl.KeyW)
		add("A", rl.KeyA)
		add("S", rl.KeyS)
		add("D", rl.KeyD)
		add("Up", rl.KeyUp)
		add("Down", rl.KeyDown)
		add("Left", rl.KeyLeft)
		add("Right", rl.KeyRight)
		add("Space", rl.KeySpace)
		add("Enter", rl.KeyEnter)
		add("Esc", rl.KeyEscape)
		add("1", rl.KeyOne)
		add("2", rl.KeyTwo)
		add("3", rl.KeyThree)
		add("4", rl.KeyFour)
		add("5", rl.KeyFive)
		add("6", rl.KeySix)
		add("7", rl.KeySeven)
		add("8", rl.KeyEight)
		add("9", rl.KeyNine)
		add("0", rl.KeyZero)
		add("F1", rl.KeyF1)
		add("F2", rl.KeyF2)
		add("F3", rl.KeyF3)
		add("F4", rl.KeyF4)
		add("F5", rl.KeyF5)
		add("F6", rl.KeyF6)
		add("F7", rl.KeyF7)
		add("F8", rl.KeyF8)
		add("F9", rl.KeyF9)
		add("F10", rl.KeyF10)
		add("F11", rl.KeyF11)
		add("F12", rl.KeyF12)

		mods := make([]string, 0, 4)
		if rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl) {
			mods = append(mods, "Ctrl")
		}
		if rl.IsKeyDown(rl.KeyRightShift) || rl.IsKeyDown(rl.KeyLeftShift) {
			mods = append(mods, "Shift")
		}
		if rl.IsKeyDown(rl.KeyRightAlt) || rl.IsKeyDown(rl.KeyLeftAlt) {
			mods = append(mods, "Alt")
		}
		if rl.IsKeyDown(rl.KeyRightSuper) || rl.IsKeyDown(rl.KeyLeftSuper) {
			mods = append(mods, "Super")
		}

		// Mouse buttons
		mouseBtns := make([]string, 0, 5)
		if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
			mouseBtns = append(mouseBtns, "L")
		}
		if rl.IsMouseButtonDown(rl.MouseButtonRight) {
			mouseBtns = append(mouseBtns, "R")
		}
		if rl.IsMouseButtonDown(rl.MouseButtonMiddle) {
			mouseBtns = append(mouseBtns, "M")
		}
		if rl.IsMouseButtonDown(rl.MouseButtonSide) {
			mouseBtns = append(mouseBtns, "Side")
		}
		if rl.IsMouseButtonDown(rl.MouseButtonExtra) {
			mouseBtns = append(mouseBtns, "Extra")
		}
		if rl.IsMouseButtonDown(rl.MouseButtonForward) {
			mouseBtns = append(mouseBtns, "Fwd")
		}
		if rl.IsMouseButtonDown(rl.MouseButtonBack) {
			mouseBtns = append(mouseBtns, "Back")
		}

		info := fmt.Sprintf(
			"FPS: %d\nMouse: [%d, %d]\nWheel:%f\nMouseBtn: %s\nKeys: %s\nMods:%s",
			rl.GetFPS(),
			int(prep.TexX),
			int(prep.TexY),
			rl.GetMouseWheelMove(),
			strings.Join(mouseBtns, ","),
			strings.Join(pressed, ","),
			strings.Join(mods, "+"),
		)

		rl.DrawText(info, int32(10), int32(10), int32(20), rl.Black)
		rl.DrawText(info, int32(9), int32(9), int32(20), rl.RayWhite)
	} else if !appFocused {
		rl.DrawText("Application not focused", int32(10), int32(10), int32(20), rl.Black)
		rl.DrawText("Application not focused", int32(9), int32(9), int32(20), rl.Gray)
	}
	rl.EndTextureMode()

	currentBackend.DrawImageRenderTexture(rt, prep.ImageSize, true, false)

	imgui.End()

	common.ShowWidgetsDemo()
}

func main() {
	common.Initialize()

	currentBackend = raylibbackend.NewRaylibBackend()
	_, _ = backend.CreateBackend[raylibbackend.RaylibBackendFlags](currentBackend)

	currentBackend.SetAfterCreateContextHook(afterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(beforeDestroyContext)

	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))
	currentBackend.CreateWindow("Hello from cimgui-go (raylib off-screen texture)", 1200, 900)
	currentBackend.SetCloseCallback(func() { fmt.Println("window is closing") })

	currentBackend.Run(loop)
}

package main

import (
	"fmt"
	"runtime"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/sdlbackend"
	"github.com/AllenDang/cimgui-go/examples/common"
	"github.com/AllenDang/cimgui-go/imgui"
)

var currentBackend backend.Backend[sdlbackend.SDLWindowFlags]

func init() {
	runtime.LockOSThread()
}

func main() {
	common.Initialize()

	currentBackend, _ = backend.CreateBackend(sdlbackend.NewSDLBackend())
	currentBackend.SetAfterCreateContextHook(common.AfterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(common.BeforeDestroyContext)

	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	currentBackend.CreateWindow("Hello from cimgui-go", 1200, 900)

	currentBackend.SetDropCallback(func(p []string) {
		fmt.Printf("drop triggered: %v", p)
	})

	currentBackend.SetCloseCallback(func() {
		fmt.Println("window is closing")
	})

	currentBackend.SetIcons(common.Image())

	currentBackend.Run(common.Loop)
}

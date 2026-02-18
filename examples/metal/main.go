//go:build darwin

package main

import (
	"fmt"
	"runtime"

	"github.com/AllenDang/cimgui-go/backend"
	metalbackend "github.com/AllenDang/cimgui-go/backend/metal"
	"github.com/AllenDang/cimgui-go/examples/common"
	"github.com/AllenDang/cimgui-go/imgui"
)

var currentBackend backend.Backend[metalbackend.MetalWindowFlags]

func init() {
	runtime.LockOSThread()
}

func main() {
	common.Initialize()

	currentBackend, _ = backend.CreateBackend(metalbackend.NewMetalBackend())
	currentBackend.SetAfterCreateContextHook(common.AfterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(common.BeforeDestroyContext)
	currentBackend.SetTargetFPS(120) // enable ProMotion

	currentBackend.SetBgColor(imgui.NewVec4(0.45, 0.55, 0.6, 1.0))

	currentBackend.CreateWindow("Hello from cimgui-go", 1200, 900)

	currentBackend.SetDropCallback(func(p []string) {
		fmt.Printf("drop triggered: %v", p)
	})

	currentBackend.SetCloseCallback(func() {
		fmt.Println("window is closing")
	})

	currentBackend.Run(common.Loop)
}

package main

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/backend"
	ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	"github.com/AllenDang/cimgui-go/examples/common"
	"github.com/AllenDang/cimgui-go/imgui"
)

var currentBackend backend.Backend[ebitenbackend.EbitenBackendFlags]

func main() {
	common.Initialize()

	currentBackend, _ = backend.CreateBackend(ebitenbackend.NewEbitenBackend())
	currentBackend.SetAfterCreateContextHook(common.AfterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(common.BeforeDestroyContext)

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

	currentBackend.SetIcons(common.Image())

	currentBackend.Run(common.Loop)
}

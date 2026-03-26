//go:build linux && cgo

package main

import (
	"runtime"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/drmeglbackend"
	"github.com/AllenDang/cimgui-go/examples/common"
	"github.com/AllenDang/cimgui-go/imgui"
)

var currentBackend backend.Backend[drmeglbackend.DRMEGLWindowFlags]

func init() {
	runtime.LockOSThread()
}

func main() {
	common.Initialize()
	currentBackend, _ = backend.CreateBackend(drmeglbackend.NewDRMEGLBackend())
	currentBackend.SetAfterCreateContextHook(common.AfterCreateContext)
	currentBackend.SetBeforeDestroyContextHook(common.BeforeDestroyContext)
	currentBackend.SetBgColor(imgui.NewVec4(0.08, 0.10, 0.12, 1.0))
	currentBackend.CreateWindow("Hello from cimgui-go DRM/EGL", 0, 0)
	currentBackend.Run(common.Loop)
}

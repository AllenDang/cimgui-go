package ebitenbackend

import (
	"image"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
)

type EbitenBackendFlags int

var _ imgui.Backend[EbitenBackendFlags] = &EbitenBackend{}

type EbitenBackend struct{}

func NewEbitenBackend() *EbitenBackend {
	return &EbitenBackend{}
}

func (b *EbitenBackend) SetAfterCreateContextHook(func())   {}
func (b *EbitenBackend) SetBeforeDestroyContextHook(func()) {}
func (b *EbitenBackend) SetBeforeRenderHook(func())         {}
func (b *EbitenBackend) SetAfterRenderHook(func())          {}

func (b *EbitenBackend) SetBgColor(color imgui.Vec4) {}
func (b *EbitenBackend) Run(func())                  {}
func (b *EbitenBackend) Refresh()                    {}

func (b *EbitenBackend) SetWindowPos(x, y int)                                            {}
func (b *EbitenBackend) GetWindowPos() (x, y int32)                                       { return 0, 0 }
func (b *EbitenBackend) SetWindowSize(width, height int)                                  {}
func (b *EbitenBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {}
func (b *EbitenBackend) SetWindowTitle(title string)                                      {}
func (b *EbitenBackend) DisplaySize() (width, height int32)                               { return 0, 0 }
func (b *EbitenBackend) SetShouldClose(bool)                                              {}
func (b *EbitenBackend) ContentScale() (xScale, yScale float32)                           { return 1, 1 }

func (b *EbitenBackend) SetTargetFPS(fps uint) {}

func (b *EbitenBackend) SetDropCallback(imgui.DropCallback)                             {}
func (b *EbitenBackend) SetCloseCallback(imgui.WindowCloseCallback[EbitenBackendFlags]) {}
func (b *EbitenBackend) SetKeyCallback(imgui.KeyCallback)                               {}
func (b *EbitenBackend) SetSizeChangeCallback(imgui.SizeChangeCallback)                 {}
func (b *EbitenBackend) SetWindowFlags(flag EbitenBackendFlags, value int)              {}
func (b *EbitenBackend) SetIcons(icons ...image.Image)                                  {}
func (b *EbitenBackend) SetSwapInterval(interval EbitenBackendFlags) error              { return nil }
func (b *EbitenBackend) SetCursorPos(x, y float64)                                      {}
func (b *EbitenBackend) SetInputMode(mode, value EbitenBackendFlags)                    {}

func (b *EbitenBackend) CreateWindow(title string, width, height int) {}

func (b *EbitenBackend) CreateTexture(pixels unsafe.Pointer, width, Height int) imgui.TextureID {
	return imgui.TextureID{}
}

func (b *EbitenBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureID {
	return imgui.TextureID{}
}

func (b *EbitenBackend) DeleteTexture(texture imgui.TextureID) {}

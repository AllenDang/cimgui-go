package ebitenbackend

import (
	"fmt"
	"image"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type EbitenBackendFlags int

var (
	_ imgui.Backend[EbitenBackendFlags] = &EbitenBackend{}
	_ ebiten.Game                       = &EbitenBackend{}
)

type EbitenBackend struct {
	loop           func()
	showDemoWindow bool
	dscale         float64
	retina         bool
	w, h           int
}

func NewEbitenBackend() *EbitenBackend {
	return &EbitenBackend{}
}

func (b *EbitenBackend) SetAfterCreateContextHook(func())   {}
func (b *EbitenBackend) SetBeforeDestroyContextHook(func()) {}
func (b *EbitenBackend) SetBeforeRenderHook(func())         {}
func (b *EbitenBackend) SetAfterRenderHook(func())          {}

func (b *EbitenBackend) SetBgColor(color imgui.Vec4) {}
func (b *EbitenBackend) Run(loop func()) {
	b.loop = loop
	b.dscale = ebiten.DeviceScaleFactor()
	b.showDemoWindow = true

	ebiten.RunGame(b)
}
func (b *EbitenBackend) Refresh() {}

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

// ebiten
func (g *EbitenBackend) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f\nFPS: %.2f\n[C]lipMask: %t", ebiten.ActualTPS(), ebiten.ActualFPS(), ClipMask()), 10, 2)
	Draw(screen)
}

func (g *EbitenBackend) Update() error {
	Update(1.0 / 60.0)

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		SetClipMask(!ClipMask())
	}

	BeginFrame()
	defer EndFrame()

	g.loop()

	return nil
}

func (g *EbitenBackend) Layout(outsideWidth, outsideHeight int) (int, int) {
	if g.retina {
		m := ebiten.DeviceScaleFactor()
		g.w = int(float64(outsideWidth) * m)
		g.h = int(float64(outsideHeight) * m)
	} else {
		g.w = outsideWidth
		g.h = outsideHeight
	}
	SetDisplaySize(float32(g.w), float32(g.h))
	return g.w, g.h
}

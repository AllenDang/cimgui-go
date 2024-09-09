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
	afterCreateContext,
	beforeRender,
	afterRender,
	beforeDestroy, // TODO This is called nowhere
	loop func()
	closeCb      imgui.WindowCloseCallback[EbitenBackendFlags]
	dscale       float64
	retina       bool
	w, h         int
	textureCache TextureCache
}

func NewEbitenBackend() *EbitenBackend {
	return &EbitenBackend{
		textureCache: NewCache(),
	}
}

func (b *EbitenBackend) SetAfterCreateContextHook(fn func()) {
	b.afterCreateContext = fn
}

func (b *EbitenBackend) SetBeforeDestroyContextHook(fn func()) {
	b.beforeDestroy = fn
}

func (b *EbitenBackend) SetBeforeRenderHook(fn func()) {
	b.beforeRender = fn
}

func (b *EbitenBackend) SetAfterRenderHook(fn func()) {
	b.afterRender = fn
}

func (b *EbitenBackend) SetBgColor(color imgui.Vec4) {}

func (b *EbitenBackend) Run(loop func()) {
	b.loop = loop
	b.dscale = ebiten.DeviceScaleFactor()

	if b.afterCreateContext != nil {
		b.afterCreateContext()
	}

	ebiten.RunGame(b)
}

func (b *EbitenBackend) Refresh() {}

func (b *EbitenBackend) SetWindowPos(x, y int) {
	ebiten.SetWindowPosition(x, y)
}

func (b *EbitenBackend) GetWindowPos() (x, y int32) {
	xInt, yInt := ebiten.WindowPosition()
	return int32(xInt), int32(yInt)
}

func (b *EbitenBackend) SetWindowSize(width, height int) {
	ebiten.SetWindowSize(width, height)
	b.w = width
	b.h = height
}

func (b *EbitenBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {}

func (b *EbitenBackend) SetWindowTitle(title string) {
	ebiten.SetWindowTitle(title)
}

func (b *EbitenBackend) DisplaySize() (width, height int32)     { return 0, 0 }
func (b *EbitenBackend) SetShouldClose(bool)                    {}
func (b *EbitenBackend) ContentScale() (xScale, yScale float32) { return 1, 1 }

func (b *EbitenBackend) SetTargetFPS(fps uint) {}

func (b *EbitenBackend) SetDropCallback(imgui.DropCallback) {}

func (b *EbitenBackend) SetCloseCallback(cb imgui.WindowCloseCallback[EbitenBackendFlags]) {
	b.closeCb = cb
}

func (b *EbitenBackend) SetKeyCallback(imgui.KeyCallback)                  {}
func (b *EbitenBackend) SetSizeChangeCallback(imgui.SizeChangeCallback)    {}
func (b *EbitenBackend) SetWindowFlags(flag EbitenBackendFlags, value int) {}
func (b *EbitenBackend) SetIcons(icons ...image.Image)                     {}
func (b *EbitenBackend) SetSwapInterval(interval EbitenBackendFlags) error { return nil }
func (b *EbitenBackend) SetCursorPos(x, y float64)                         {}
func (b *EbitenBackend) SetInputMode(mode, value EbitenBackendFlags)       {}

func (b *EbitenBackend) CreateWindow(title string, width, height int) {}

func (b *EbitenBackend) CreateTexture(pixels unsafe.Pointer, width, height int) imgui.TextureID {
	eimg := ebiten.NewImage(width, height)
	eimg.WritePixels(premultiplyPixels(pixels, width, height))

	tid := imgui.TextureID{Data: uintptr(b.textureCache.NextId())}
	b.textureCache.SetTexture(tid, eimg)
	return tid
}

func (b *EbitenBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureID {
	pix := img.Pix
	return b.CreateTexture(unsafe.Pointer(&pix), width, height)
}

func (b *EbitenBackend) DeleteTexture(id imgui.TextureID) {
	b.textureCache.RemoveTexture(id)
}

// ebiten
func (g *EbitenBackend) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f\nFPS: %.2f\n[C]lipMask: %t", ebiten.ActualTPS(), ebiten.ActualFPS(), ClipMask()), 10, 2)
	Draw(screen)
}

func (g *EbitenBackend) Update() error {
	if ebiten.IsWindowBeingClosed() {
		g.closeCb(g)
	}

	if g.beforeRender != nil {
		g.beforeRender()
	}

	Update(1.0 / 60.0)

	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		SetClipMask(!ClipMask())
	}

	BeginFrame()
	defer EndFrame()

	g.loop()

	defer func() {
		if g.afterRender != nil {
			g.afterRender()
		}
	}()

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

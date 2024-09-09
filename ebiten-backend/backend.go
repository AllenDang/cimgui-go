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

const (
	// EbitenBackendFlagsCursorMode sets the cursor mode.
	// refer ebiten.CursorModeType
	// CursorModeVisible, CursorModeHidden, CursorModeCaptured
	EbitenBackendFlagsCursorMode EbitenBackendFlags = iota
	// EbitenBackendFlagsCursorShape sets the cursor shape.
	// refer ebiten.CursorShapeType
	// CursorShapeDefault, CursorShapeText, CursorShapeCrosshair, CursorShapePointer,
	// CursorShapeEWResize, CursorShapeNSResize, CursorShapeNESWResize, CursorShapeNWSEResize,
	// CursorShapeMove, CursorShapeNotAllowed CursorShapeType, EbitenBackendFlagsCursorShape
	EbitenBackendFlagsCursorShape
	EbitenBackendFlagsResizingMode
	EbitenBackendFlagsFPSMode
	EbitenBackendFlagsDecorated
	EbitenBackendFlagsFloating
	EbitenBackendFlagsMaximized
	EbitenBackendFlagsMinimized
	EbitenBackendFlagsClosingHandled
	EbitenBackendFlagsMousePassthrough
)

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
	closeCb imgui.WindowCloseCallback[EbitenBackendFlags]
	dscale  float64
	retina  bool
	w, h    int
	manager *Manager
	fps     uint
}

func NewEbitenBackend() *EbitenBackend {
	return &EbitenBackend{
		manager: NewManager(nil),
		fps:     60,
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

func (e *EbitenBackend) SetTargetFPS(fps uint) {
	e.fps = fps
}

func (b *EbitenBackend) SetDropCallback(imgui.DropCallback) {}

func (b *EbitenBackend) SetCloseCallback(cb imgui.WindowCloseCallback[EbitenBackendFlags]) {
	b.closeCb = cb
}

func (b *EbitenBackend) SetKeyCallback(imgui.KeyCallback)               {}
func (b *EbitenBackend) SetSizeChangeCallback(imgui.SizeChangeCallback) {}

func (b *EbitenBackend) SetWindowFlags(flag EbitenBackendFlags, value int) {
	switch flag {
	case EbitenWindowFlagsCursorMode:
		ebiten.SetCursorMode(ebiten.CursorModeType(value))
	case EbitenWindowFlagsCursorShape:
		ebiten.SetCursorShape(ebiten.CursorShapeType(value))
	case EbitenWindowFlagsFPSMode:
		ebiten.SetVsyncEnabled(value == 0)
	case EbitenWindowFlagsResizingMode:
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeType(value))
	case EbitenWindowFlagsDecorated:
		ebiten.SetWindowDecorated(value != 0)
	case EbitenWindowFlagsFloating:
		ebiten.SetWindowFloating(value != 0)
	case EbitenWindowFlagsMaximized:
		if value != 0 {
			ebiten.MaximizeWindow()
		} else {
			ebiten.RestoreWindow()
		}
	case EbitenWindowFlagsMinimized:
		if value != 0 {
			ebiten.MinimizeWindow()
		} else {
			ebiten.RestoreWindow()
		}
	case EbitenWindowFlagsClosingHandled:
		ebiten.SetWindowClosingHandled(value != 0)
	case EbitenWindowFlagsMousePassthrough:
		ebiten.SetWindowMousePassthrough(value != 0)
	default:
		panic("Invalid flag for SetWindowFlags.")
	}
}

func (b *EbitenBackend) SetIcons(icons ...image.Image) {
	ebiten.SetWindowIcon(icons)
}

func (b *EbitenBackend) SetSwapInterval(interval EbitenBackendFlags) error { return nil }
func (b *EbitenBackend) SetCursorPos(x, y float64)                         {}
func (b *EbitenBackend) SetInputMode(mode, value EbitenBackendFlags)       {}

func (e *EbitenBackend) CreateWindow(title string, width, height int) {
	e.SetWindowTitle(title)
	e.SetWindowSize(width, height)
}

func (e *EbitenBackend) CreateTexture(pixels unsafe.Pointer, width, height int) imgui.TextureID {
	eimg := ebiten.NewImage(width, height)
	eimg.WritePixels(premultiplyPixels(pixels, width, height))

	tid := imgui.TextureID{Data: uintptr(e.manager.Cache.NextId())}
	e.manager.Cache.SetTexture(tid, eimg)
	return tid
}

func (b *EbitenBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureID {
	pix := img.Pix
	return b.CreateTexture(unsafe.Pointer(&pix[0]), width, height)
}

func (e *EbitenBackend) DeleteTexture(id imgui.TextureID) {
	e.manager.Cache.RemoveTexture(id)
}

// ebiten

// Draw draws the generated imgui frame to the screen.
// This is usually called inside the game's Draw() function.
func (e *EbitenBackend) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("TPS: %.2f\nFPS: %.2f\n[C]lipMask: %t", ebiten.ActualTPS(), ebiten.ActualFPS(), e.ClipMask()), 10, 2)
	e.manager.Draw(screen)
}

// Update needs to be called on every frame, before cimgui-go calls.
// This is usually called inside the game's Update() function.
// delta is the time in seconds since the last frame.
func (e *EbitenBackend) Update() error {
	if ebiten.IsWindowBeingClosed() {
		e.closeCb(e)
	}

	if e.beforeRender != nil {
		e.beforeRender()
	}

	e.manager.Update(1.0 / float32(e.fps))

	// TODO: what is that?
	if inpututil.IsKeyJustPressed(ebiten.KeyC) {
		e.SetClipMask(!e.ClipMask())
	}

	e.BeginFrame()
	defer e.EndFrame()

	e.loop()

	defer func() {
		if e.afterRender != nil {
			e.afterRender()
		}
	}()

	return nil
}

func (e *EbitenBackend) Layout(outsideWidth, outsideHeight int) (int, int) {
	if e.retina {
		m := ebiten.DeviceScaleFactor()
		e.w = int(float64(outsideWidth) * m)
		e.h = int(float64(outsideHeight) * m)
	} else {
		e.w = outsideWidth
		e.h = outsideHeight
	}

	e.SetDisplaySize(float32(e.w), float32(e.h))

	return e.w, e.h
}

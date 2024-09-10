package ebitenbackend

import (
	"image"
	"runtime"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
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
	// EbitenBackendFlagsDebug is a flag to enable debug mode. It will show FPS, TPS, ClipMask and enable ClipMask shortcut.
	// 0 (default) disabled, 1 (or anything else) enabled
	EbitenBackendFlagsDebug
)

var (
	_ imgui.Backend[EbitenBackendFlags] = &EbitenBackend{}
	_ ebiten.Game                       = &EbitenBackend{}
)

type EbitenBackend struct {
	// cimgui-go backend specific:
	afterCreateContext,
	beforeRender,
	afterRender,
	beforeDestroy,
	loop func()
	closeCb imgui.WindowCloseCallback[EbitenBackendFlags]

	// ebiten stuff
	filter   ebiten.Filter
	dscale   float64
	retina   bool
	w, h     int
	fps      uint // target FPS
	bgColor  imgui.Vec4
	debug    bool // if true render some extra debug info
	clipMask bool // https://github.com/ocornut/imgui/issues/7372
	lmask    *ebiten.Image

	syncInputs         bool
	syncInputsFn       func()
	syncCursor         bool
	controlCursorShape bool

	cache TextureCache
	ctx   *imgui.Context // imgui context

	manager *Manager // TODO: remove this entirely
}

// this is "pointer" to the first texture used for font atlas texture.
// it needs to be a var and cannot be more private because of CGO stuff.
var id1 = 1

// NewEbitenBackend creates a new Ebiten backend.
// it takes font atlas which could be nil
// - TODO: make font atlas a factory method and move some stuff to CreateWindow
func NewEbitenBackend(fontAtlas *imgui.FontAtlas) *EbitenBackend {
	var imctx *imgui.Context

	if fontAtlas != nil {
		imctx = imgui.CreateContextV(fontAtlas)
	} else {
		imctx = imgui.CreateContext()
	}

	result := &EbitenBackend{
		ctx:                imctx,
		cache:              NewCache(),
		manager:            NewManager(nil),
		fps:                60,
		clipMask:           true,
		syncInputs:         true,
		syncCursor:         true,
		controlCursorShape: true,
	}

	// Build texture atlas
	fonts := imgui.CurrentIO().Fonts()
	_, _, _, _ = fonts.GetTextureDataAsRGBA32() // call this to force imgui to build the font atlas cache

	texID := imgui.TextureID{}
	texID.Data = uintptr(unsafe.Pointer(&id1)) // TODO: this shit will cause -race issues
	fonts.SetTexID(texID)

	result.cache.SetFontAtlasTextureID(texID)

	runtime.SetFinalizer(result, (*EbitenBackend).onfinalize)

	result.setKeyMapping()

	return result
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

func (b *EbitenBackend) SetBgColor(color imgui.Vec4) {
	b.bgColor = color
}

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
	case EbitenBackendFlagsCursorMode:
		ebiten.SetCursorMode(ebiten.CursorModeType(value))
	case EbitenBackendFlagsCursorShape:
		ebiten.SetCursorShape(ebiten.CursorShapeType(value))
	case EbitenBackendFlagsFPSMode:
		ebiten.SetVsyncEnabled(value == 0)
	case EbitenBackendFlagsResizingMode:
		ebiten.SetWindowResizingMode(ebiten.WindowResizingModeType(value))
	case EbitenBackendFlagsDecorated:
		ebiten.SetWindowDecorated(value != 0)
	case EbitenBackendFlagsFloating:
		ebiten.SetWindowFloating(value != 0)
	case EbitenBackendFlagsMaximized:
		if value != 0 {
			ebiten.MaximizeWindow()
		} else {
			ebiten.RestoreWindow()
		}
	case EbitenBackendFlagsMinimized:
		if value != 0 {
			ebiten.MinimizeWindow()
		} else {
			ebiten.RestoreWindow()
		}
	case EbitenBackendFlagsClosingHandled:
		ebiten.SetWindowClosingHandled(value != 0)
	case EbitenBackendFlagsMousePassthrough:
		ebiten.SetWindowMousePassthrough(value != 0)
	case EbitenBackendFlagsDebug:
		b.debug = value != 0
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

	tid := imgui.TextureID{Data: uintptr(e.cache.NextId())}
	e.cache.SetTexture(tid, eimg)
	return tid
}

func (b *EbitenBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureID {
	pix := img.Pix
	return b.CreateTexture(unsafe.Pointer(&pix[0]), width, height)
}

func (e *EbitenBackend) DeleteTexture(id imgui.TextureID) {
	e.cache.RemoveTexture(id)
}

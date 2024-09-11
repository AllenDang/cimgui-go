package ebitenbackend

import (
	"runtime"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

type getCursorFn func() (x, y float32)

var (
	_ imgui.Backend[EbitenBackendFlags] = &EbitenBackend{}
	_ ebiten.Game                       = &EbitenBackend{}
)

// EbitenBackend implements imgui.Backend and ebiten.Game.
// It allows to render imgui UI using ebitengine (https://github.com/hajimehoshi/ebiten).
type EbitenBackend struct {
	customFontAtlas *imgui.FontAtlas
	customCtx       *imgui.Context

	// cimgui-go backend specific:
	afterCreateContext,
	beforeRender,
	afterRender,
	beforeDestroy,
	loop func()
	closeCb imgui.WindowCloseCallback[EbitenBackendFlags]

	// ebiten stuff
	filter                      ebiten.Filter
	dscale                      float64
	retina                      bool
	currentWidth, currentHeight int
	fps                         uint // target FPS
	bgColor                     imgui.Vec4
	debug                       bool // if true render some extra debug info
	clipMask                    bool // https://github.com/ocornut/imgui/issues/7372
	lmask                       *ebiten.Image

	syncInputs         bool
	syncInputsFn       func()
	syncCursor         bool
	controlCursorShape bool
	getCursor          getCursorFn
	cliptxt            string
	inputChars         []rune

	cache TextureCache
	ctx   *imgui.Context // imgui context
}

// this is "pointer" to the first texture used for font atlas texture.
// it needs to be a var and cannot be more private because of CGO stuff.
var id1 = 1

// NewEbitenBackend creates a new Ebiten backend.
// it takes font atlas which could be nil
func NewEbitenBackend() *EbitenBackend {
	result := &EbitenBackend{
		cache:              NewCache(),
		fps:                60,
		clipMask:           true,
		syncInputs:         true,
		syncCursor:         true,
		controlCursorShape: true,
		inputChars:         make([]rune, 0, 256),
	}

	runtime.SetFinalizer(result, (*EbitenBackend).onfinalize)

	result.setKeyMapping()

	return result
}

// SetFontAtlas sets custom font atlas
// * do not use SetContext along with this
// * if not called, CreateWindow will create new context with nil font atlas
func (e *EbitenBackend) SetFontAtlas(fa *imgui.FontAtlas) *EbitenBackend {
	e.customFontAtlas = fa
	return e
}

// SetContext sets imgui.Context (if not set, CreateWindow will create one)
func (e *EbitenBackend) SetContext(ctx *imgui.Context) *EbitenBackend {
	e.customCtx = ctx
	return e
}

// BeginFrame needs to be called on every frame, before cimgui-go calls.
// This is usually called inside the game's Update() function.
func (e *EbitenBackend) BeginFrame() {
	imgui.NewFrame()
}

// EndFrame needs to be called on every frame, after cimgui-go calls.
// This is usually called inside the game's Update() function.
func (e *EbitenBackend) EndFrame() {
	imgui.EndFrame()
}

// SetClipMask sets if clipmask is enabled or not. This is usually called for debugging purposes.
func (e *EbitenBackend) SetClipMask(value bool) {
	e.clipMask = value
}

// ClipMask returns if clipmask is enabled or not.
func (e *EbitenBackend) ClipMask() bool {
	return e.clipMask
}

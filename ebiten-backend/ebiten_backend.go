package ebitenbackend

import (
	"runtime"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	_ imgui.Backend[EbitenBackendFlags] = &EbitenBackend{}
	_ ebiten.Game                       = &EbitenBackend{}
)

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
func NewEbitenBackend() *EbitenBackend {
	result := &EbitenBackend{
		cache:              NewCache(),
		manager:            NewManager(nil),
		fps:                60,
		clipMask:           true,
		syncInputs:         true,
		syncCursor:         true,
		controlCursorShape: true,
	}

	runtime.SetFinalizer(result, (*EbitenBackend).onfinalize)

	result.setKeyMapping()

	return result
}

func (e *EbitenBackend) SetFontAtlas(fa *imgui.FontAtlas) *EbitenBackend {
	e.customFontAtlas = fa
	return e
}

func (e *EbitenBackend) SetContext(ctx *imgui.Context) *EbitenBackend {
	e.customCtx = ctx
	return e
}

package ebitenbackend

import (
	"runtime"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
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

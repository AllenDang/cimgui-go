package ebitenbackend

import (
	"image/color"
	"runtime"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type getCursorFn func() (x, y float32)

var (
	_ backend.Backend[EbitenBackendFlags] = &EbitenBackend{}
	_ ebiten.Game                         = &EbitenBackend{}
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
	closeCb     backend.WindowCloseCallback
	shouldClose bool
	resizeCb    backend.SizeChangeCallback

	// ebiten stuff
	filter                      ebiten.Filter
	dscale                      float64
	retina                      bool
	currentWidth, currentHeight int
	fps                         uint // target FPS
	bgColor                     imgui.Vec4
	bgColorMagic                struct { // this is code from https://github.com/tinne26/mipix/blob/main/internal/utils.go#L113, thank you @tinne26
		pkgMask1x1           *ebiten.Image
		pkgFillVertices      []ebiten.Vertex
		pkgFillVertIndices   []uint16
		pkgFillTrianglesOpts ebiten.DrawTrianglesOptions
	}
	debug    bool // if true render some extra debug info
	clipMask bool // https://github.com/ocornut/imgui/issues/7372
	lmask    *ebiten.Image

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
var id1 imgui.TextureID = 1

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

	result.bgColorMagic.pkgMask1x1 = ebiten.NewImage(1, 1)
	result.bgColorMagic.pkgFillVertices = make([]ebiten.Vertex, 4)
	result.bgColorMagic.pkgFillVertIndices = []uint16{0, 1, 3, 3, 1, 2}
	for i := range 4 {
		result.bgColorMagic.pkgFillVertices[i].SrcX = 0.5
		result.bgColorMagic.pkgFillVertices[i].SrcY = 0.5
	}

	result.bgColorMagic.pkgMask1x1.Fill(color.White)

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
	if ebiten.IsWindowBeingClosed() && e.closeCb != nil {
		e.closeCb()
	}

	if e.beforeRender != nil {
		e.beforeRender()
	}

	io := imgui.CurrentIO()
	io.SetDisplaySize(imgui.Vec2{X: float32(e.currentWidth), Y: float32(e.currentHeight)})

	io.SetDeltaTime(1.0 / float32(e.fps))
	if e.syncCursor {
		if e.getCursor != nil {
			x, y := e.getCursor()
			io.SetMousePos(imgui.Vec2{X: x, Y: y})
		} else {
			mx, my := ebiten.CursorPosition()
			io.SetMousePos(imgui.Vec2{X: float32(mx), Y: float32(my)})
		}
		io.SetMouseButtonDown(0, ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft))
		io.SetMouseButtonDown(1, ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight))
		io.SetMouseButtonDown(2, ebiten.IsMouseButtonPressed(ebiten.MouseButtonMiddle))
		xoff, yoff := ebiten.Wheel()
		io.AddMouseWheelDelta(float32(xoff), float32(yoff))
		e.controlCursorShapeFn()
	}

	if e.syncInputs {
		if e.syncInputsFn != nil {
			e.syncInputsFn()
		} else {
			e.inputChars = sendInput(imgui.CurrentIO(), e.inputChars)
		}
	}

	if e.debug { // add ClipMask switch
		if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			e.SetClipMask(!e.ClipMask())
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyI) {
			e.syncInputs = !e.syncInputs
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyU) {
			e.syncCursor = !e.syncCursor
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyS) {
			e.controlCursorShape = !e.controlCursorShape
		}
	}

	imgui.NewFrame()
}

// EndFrame needs to be called on every frame, after cimgui-go calls.
// This is usually called inside the game's Update() function.
func (e *EbitenBackend) EndFrame() {
	imgui.EndFrame()
	if e.afterRender != nil {
		e.afterRender()
	}
}

// SetClipMask sets if clipmask is enabled or not. This is usually called for debugging purposes.
func (e *EbitenBackend) SetClipMask(value bool) {
	e.clipMask = value
}

// ClipMask returns if clipmask is enabled or not.
func (e *EbitenBackend) ClipMask() bool {
	return e.clipMask
}

package ebitenbackend

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

type GetCursorFn func() (x, y float32)

type Manager struct {
	Filter             ebiten.Filter
	cliptxt            string
	GetCursor          GetCursorFn
	SyncInputsFn       func()
	SyncCursor         bool
	SyncInputs         bool
	ControlCursorShape bool
	lmask              *ebiten.Image
	ClipMask           bool

	width        float32
	height       float32
	screenWidth  int
	screenHeight int

	inputChars []rune
}

func NewManager(fontAtlas *imgui.FontAtlas) *Manager {
	m := &Manager{
		SyncCursor:         true,
		SyncInputs:         true,
		ClipMask:           true,
		ControlCursorShape: true,
		inputChars:         make([]rune, 0, 256),
	}

	m.setKeyMapping()

	return m
}

func NewManagerWithContext(ctx *imgui.Context) *Manager {
	m := &Manager{
		SyncCursor:         true,
		SyncInputs:         true,
		ClipMask:           true,
		ControlCursorShape: true,
	}
	m.setKeyMapping()
	return m
}

func (m *Manager) controlCursorShape() {
	if !m.ControlCursorShape {
		return
	}

	switch imgui.CurrentMouseCursor() {
	case imgui.MouseCursorNone:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	case imgui.MouseCursorArrow:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	case imgui.MouseCursorTextInput:
		ebiten.SetCursorShape(ebiten.CursorShapeText)
	case imgui.MouseCursorResizeAll:
		ebiten.SetCursorShape(ebiten.CursorShapeCrosshair)
	case imgui.MouseCursorResizeEW:
		ebiten.SetCursorShape(ebiten.CursorShapeEWResize)
	case imgui.MouseCursorResizeNS:
		ebiten.SetCursorShape(ebiten.CursorShapeNSResize)
	case imgui.MouseCursorHand:
		ebiten.SetCursorShape(ebiten.CursorShapePointer)
	default:
		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	}
}

func (m *Manager) setKeyMapping() {
	// Keyboard mapping. ImGui will use those indices to peek into the io.KeysDown[] array.
	/*
		io := cimgui.GetIO()
		for imguiKey, nativeKey := range keys {
			// io.KeyMap(int(imguiKey), nativeKey)
		}
	*/
}

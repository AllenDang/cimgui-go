package ebitenbackend

import (
	"image"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

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

// TODO: Not implemented
func (b *EbitenBackend) Refresh() {
	panic("Refresh is not implemented for Ebiten backend yet.")
}

func (b *EbitenBackend) SetWindowPos(x, y int) {
	ebiten.SetWindowPosition(x, y)
}

func (b *EbitenBackend) GetWindowPos() (x, y int32) {
	xInt, yInt := ebiten.WindowPosition()
	return int32(xInt), int32(yInt)
}

func (b *EbitenBackend) SetWindowSize(width, height int) {
	ebiten.SetWindowSize(width, height)
	/*
		b.width = float32(width)
		b.height = float32(height)
	*/
}

func (b *EbitenBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {} // TODO

func (b *EbitenBackend) SetWindowTitle(title string) {
	ebiten.SetWindowTitle(title)
}

// TODO: Not implemented
func (b *EbitenBackend) DisplaySize() (width, height int32) {
	panic("SetShouldClose is not implemented for Ebiten backend yet.")
	return 0, 0
}

// TODO: Not implemented
func (b *EbitenBackend) SetShouldClose(bool) {
	panic("SetShouldClose is not implemented for Ebiten backend yet.")
}

func (b *EbitenBackend) ContentScale() (xScale, yScale float32) {
	scale := ebiten.DeviceScaleFactor()
	return float32(scale), float32(scale)
}

func (e *EbitenBackend) SetTargetFPS(fps uint) {
	e.fps = fps
}

// TODO: Not implemented
func (b *EbitenBackend) SetDropCallback(imgui.DropCallback) {
	panic("SetDropCallback is not implemented for Ebiten backend yet.")
}

func (b *EbitenBackend) SetCloseCallback(cb imgui.WindowCloseCallback[EbitenBackendFlags]) {
	b.closeCb = cb
}

func (b *EbitenBackend) SetKeyCallback(imgui.KeyCallback)               {} // TODO
func (b *EbitenBackend) SetSizeChangeCallback(imgui.SizeChangeCallback) {} // TODO

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

// TODO: Not implemented
func (b *EbitenBackend) SetSwapInterval(interval EbitenBackendFlags) error {
	panic("SetSwapInterval is not implemented for Ebiten backend yet.")
	return nil
}

// TODO: Not implemented
func (b *EbitenBackend) SetCursorPos(x, y float64) {
	panic("SetCursorPos is not implemented for Ebiten backend yet.")
}

// TODO: Not implemented
func (b *EbitenBackend) SetInputMode(mode, value EbitenBackendFlags) {
	panic("SetInputMode is not implemented for Ebiten backend yet.")
}

func (e *EbitenBackend) CreateWindow(title string, width, height int) {
	// create context
	var imctx *imgui.Context

	switch {
	case e.customCtx != nil:
		imctx = e.customCtx
	case e.customFontAtlas != nil:
		imctx = imgui.CreateContextV(e.customFontAtlas)
	default:
		imctx = imgui.CreateContext()
	}

	e.ctx = imctx

	// Build texture atlas
	fonts := imgui.CurrentIO().Fonts()
	_, _, _, _ = fonts.GetTextureDataAsRGBA32() // call this to force imgui to build the font atlas cache

	texID := imgui.TextureID{}
	texID.Data = uintptr(unsafe.Pointer(&id1)) // TODO: this shit will cause -race issues
	fonts.SetTexID(texID)

	e.cache.SetFontAtlasTextureID(texID)

	// initialize ebiten stuff
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

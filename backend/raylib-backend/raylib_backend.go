package raylibbackend

import (
	"image"
	"unsafe"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	_ backend.Backend[RaylibBackendFlags] = &RaylibBackend{}
)

type RaylibBackend struct {
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

	cache TextureCache
	ctx   *imgui.Context // imgui context

	window unsafe.Pointer
}

func (r *RaylibBackend) SetAfterCreateContextHook(f func()) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetBeforeDestroyContextHook(f func()) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetBeforeRenderHook(f func()) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetAfterRenderHook(f func()) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetBgColor(color imgui.Vec4) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) Run(f func()) {
	r.loop = f

	if r.afterCreateContext != nil {
		r.afterCreateContext()
	}

}

func (r *RaylibBackend) Refresh() {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetWindowPos(x, y int) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) GetWindowPos() (x, y int32) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetWindowSize(width, height int) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetWindowSizeLimits(minWidth, minHeight, maxWidth, maxHeight int) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetWindowTitle(title string) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) DisplaySize() (width, height int32) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetShouldClose(b bool) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) ContentScale() (xScale, yScale float32) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetTargetFPS(fps uint) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetDropCallback(callback backend.DropCallback) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetCloseCallback(callback backend.WindowCloseCallback) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetKeyCallback(callback backend.KeyCallback) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetSizeChangeCallback(callback backend.SizeChangeCallback) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetWindowFlags(flag RaylibBackendFlags, value int) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetIcons(icons ...image.Image) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetSwapInterval(interval RaylibBackendFlags) error {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetCursorPos(x, y float64) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) SetInputMode(mode, value RaylibBackendFlags) {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) CreateWindow(title string, width, height int) {
	rl.InitWindow(int32(width), int32(height), title)

}

func (r *RaylibBackend) CreateTexture(pixels unsafe.Pointer, width, Height int) imgui.TextureRef {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureRef {
	//TODO implement me
	panic("implement me")
}

func (r *RaylibBackend) DeleteTexture(id imgui.TextureRef) {
	//TODO implement me
	panic("implement me")
}

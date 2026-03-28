//go:build linux && cgo

package drmeglbackend

import (
	"errors"
	"image"
	"time"
	"unsafe"

	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/impl/opengl3"
	drm "github.com/tthhr/go-drm-egl/drm"
)

type voidCallbackFunc func()

type DRMEGLWindowFlags int

const (
	DRMEGLWindowFlagsNone      DRMEGLWindowFlags = 0
	DRMEGLWindowFlagsOffscreen DRMEGLWindowFlags = 1 << iota
)

const (
	DRMEGLSwapIntervalImmediate = DRMEGLWindowFlags(0)
	DRMEGLSwapIntervalVsync     = DRMEGLWindowFlags(1)
)

var (
	ErrContextNotCreated        = errors.New("drmeglbackend: context not created")
	ErrImmediateSwapUnsupported = errors.New("drmeglbackend: immediate swap interval is not supported by go-drm-egl")
)

var _ backend.Backend[DRMEGLWindowFlags] = &DRMEGLBackend{}

type DRMEGLBackend struct {
	afterCreateContext   voidCallbackFunc
	loop                 voidCallbackFunc
	beforeRender         voidCallbackFunc
	afterRender          voidCallbackFunc
	beforeDestroyContext voidCallbackFunc
	dropCB               backend.DropCallback
	closeCB              func(pointer unsafe.Pointer)
	keyCb                backend.KeyCallback
	sizeCb               backend.SizeChangeCallback

	ctx         *drm.Context
	shouldClose bool
	targetFPS   uint
	bgColor     imgui.Vec4
	offscreen   bool
	title       string
	width       int
	height      int
}

func NewDRMEGLBackend() *DRMEGLBackend {
	return &DRMEGLBackend{
		targetFPS: 60,
		bgColor:   imgui.NewVec4(0.45, 0.55, 0.6, 1.0),
		closeCB:   func(unsafe.Pointer) {},
		dropCB:    func([]string) {},
		keyCb:     func(int, int, int, int) {},
		sizeCb:    func(int, int) {},
	}
}

func (b *DRMEGLBackend) SetAfterCreateContextHook(hook func()) {
	b.afterCreateContext = hook
}

func (b *DRMEGLBackend) AfterCreateHook() func() {
	return b.afterCreateContext
}

func (b *DRMEGLBackend) SetBeforeDestroyContextHook(hook func()) {
	b.beforeDestroyContext = hook
}

func (b *DRMEGLBackend) BeforeDestroyHook() func() {
	return b.beforeDestroyContext
}

func (b *DRMEGLBackend) SetBeforeRenderHook(hook func()) {
	b.beforeRender = hook
}

func (b *DRMEGLBackend) BeforeRenderHook() func() {
	return b.beforeRender
}

func (b *DRMEGLBackend) SetAfterRenderHook(hook func()) {
	b.afterRender = hook
}

func (b *DRMEGLBackend) AfterRenderHook() func() {
	return b.afterRender
}

func (b *DRMEGLBackend) SetBgColor(color imgui.Vec4) {
	b.bgColor = color
}

func (b *DRMEGLBackend) Run(loop func()) {
	if b.ctx == nil {
		panic(ErrContextNotCreated)
	}

	b.loop = loop
	lastFrame := time.Now()
	frameDuration := time.Second / time.Duration(maxUint(b.targetFPS, 1))

	for !b.shouldClose {
		now := time.Now()
		delta := now.Sub(lastFrame)
		lastFrame = now

		if hook := b.beforeRender; hook != nil {
			hook()
		}

		io := imgui.CurrentIO()
		io.SetDisplaySize(imgui.NewVec2(float32(b.ctx.Width()), float32(b.ctx.Height())))
		io.SetDeltaTime(float32(delta.Seconds()))

		drm.ViewPort(0, 0, b.ctx.Width(), b.ctx.Height())
		opengl3.NewFrame()
		imgui.NewFrame()

		if b.loop != nil {
			b.loop()
		}

		imgui.Render()
		drm.ClearColor(b.bgColor.X, b.bgColor.Y, b.bgColor.Z, b.bgColor.W)
		drm.Clear(drm.COLOR_BUFFER_BIT)
		opengl3.RenderDrawData(imgui.CurrentDrawData())
		drm.RenderFrame(b.ctx)

		if hook := b.afterRender; hook != nil {
			hook()
		}

		if sleep := frameDuration - time.Since(now); sleep > 0 {
			time.Sleep(sleep)
		}
	}

	b.destroy()
}

func (b *DRMEGLBackend) destroy() {
	opengl3.Shutdown()
	if hook := b.beforeDestroyContext; hook != nil {
		hook()
	}
	imgui.DestroyContext()
	if b.ctx != nil {
		b.ctx.Cleanup()
		b.ctx = nil
	}
	if b.closeCB != nil {
		b.closeCB(nil)
	}
}

func (b *DRMEGLBackend) LoopFunc() func() {
	return b.loop
}

func (b *DRMEGLBackend) DropCallback() backend.DropCallback {
	return b.dropCB
}

func (b *DRMEGLBackend) CloseCallback() func(wnd unsafe.Pointer) {
	return b.closeCB
}

func (b *DRMEGLBackend) SetWindowPos(_, _ int) {}

func (b *DRMEGLBackend) GetWindowPos() (x, y int32) {
	return 0, 0
}

func (b *DRMEGLBackend) SetWindowSize(width, height int) {
	b.width = width
	b.height = height
}

func (b *DRMEGLBackend) SetWindowSizeLimits(_, _, _, _ int) {}

func (b *DRMEGLBackend) SetWindowTitle(title string) {
	b.title = title
}

func (b *DRMEGLBackend) DisplaySize() (width, height int32) {
	if b.ctx == nil {
		return int32(b.width), int32(b.height)
	}
	return int32(b.ctx.Width()), int32(b.ctx.Height())
}

func (b *DRMEGLBackend) ContentScale() (xScale, yScale float32) {
	return 1, 1
}

func (b *DRMEGLBackend) SetShouldClose(value bool) {
	b.shouldClose = value
}

func (b *DRMEGLBackend) SetTargetFPS(fps uint) {
	b.targetFPS = maxUint(fps, 1)
}

func (b *DRMEGLBackend) Refresh() {}

func (b *DRMEGLBackend) CreateTexture(pixels unsafe.Pointer, width, height int) imgui.TextureRef {
	tex := drm.GenTextures(1)
	drm.BindTexture(drm.TEXTURE_2D, tex)
	drm.TexParameteri(drm.TEXTURE_2D, drm.TEXTURE_MIN_FILTER, int32(drm.LINEAR))
	drm.TexParameteri(drm.TEXTURE_2D, drm.TEXTURE_MAG_FILTER, int32(drm.LINEAR))
	drm.TexParameteri(drm.TEXTURE_2D, drm.TEXTURE_WRAP_S, int32(drm.CLAMP_TO_EDGE))
	drm.TexParameteri(drm.TEXTURE_2D, drm.TEXTURE_WRAP_T, int32(drm.CLAMP_TO_EDGE))
	drm.TexImage2D(
		drm.TEXTURE_2D,
		0,
		drm.RGBA,
		int32(width),
		int32(height),
		0,
		drm.RGBA,
		drm.UNSIGNED_BYTE,
		pixels,
	)
	drm.BindTexture(drm.TEXTURE_2D, 0)
	return *imgui.NewTextureRefTextureID(imgui.TextureID(tex))
}

func (b *DRMEGLBackend) CreateTextureRgba(img *image.RGBA, width, height int) imgui.TextureRef {
	return b.CreateTexture(unsafe.Pointer(&img.Pix[0]), width, height)
}

func (b *DRMEGLBackend) DeleteTexture(id imgui.TextureRef) {
	tex := uint32(id.TexID())
	drm.DeleteTextures(tex)
}

func (b *DRMEGLBackend) CreateWindow(title string, width, height int) {
	b.title = title
	b.width = width
	b.height = height

	requestWidth, requestHeight := 0, 0
	if b.offscreen {
		requestWidth = width
		requestHeight = height
	}

	ctx, err := drm.Init(requestWidth, requestHeight)
	if err != nil {
		panic(err)
	}
	if !drm.MakeCurrent(ctx) {
		ctx.Cleanup()
		panic("drmeglbackend: failed to make EGL context current")
	}

	b.ctx = ctx
	imgui.CreateContext()
	imgui.StyleColorsDark()

	if hook := b.afterCreateContext; hook != nil {
		hook()
	}

	if !opengl3.InitV("#version 300 es") {
		b.destroy()
		panic("drmeglbackend: failed to initialize imgui opengl3 renderer")
	}

	if cb := b.sizeCb; cb != nil {
		cb(ctx.Width(), ctx.Height())
	}
}

func (b *DRMEGLBackend) SetDropCallback(cbfun backend.DropCallback) {
	if cbfun == nil {
		b.dropCB = func([]string) {}
		return
	}
	b.dropCB = cbfun
}

func (b *DRMEGLBackend) SetCloseCallback(cbfun backend.WindowCloseCallback) {
	if cbfun == nil {
		b.closeCB = func(unsafe.Pointer) {}
		return
	}
	b.closeCB = func(_ unsafe.Pointer) {
		cbfun()
	}
}

func (b *DRMEGLBackend) SetWindowFlags(flag DRMEGLWindowFlags, value int) {
	if flag == DRMEGLWindowFlagsOffscreen {
		b.offscreen = value != 0
	}
}

func (b *DRMEGLBackend) SetIcons(...image.Image) {}

func (b *DRMEGLBackend) SetKeyCallback(cbfun backend.KeyCallback) {
	if cbfun == nil {
		b.keyCb = func(int, int, int, int) {}
		return
	}
	b.keyCb = cbfun
}

func (b *DRMEGLBackend) KeyCallback() backend.KeyCallback {
	return b.keyCb
}

func (b *DRMEGLBackend) SetSizeChangeCallback(cbfun backend.SizeChangeCallback) {
	if cbfun == nil {
		b.sizeCb = func(int, int) {}
		return
	}
	b.sizeCb = cbfun
}

func (b *DRMEGLBackend) SizeCallback() backend.SizeChangeCallback {
	return b.sizeCb
}

func (b *DRMEGLBackend) SetSwapInterval(interval DRMEGLWindowFlags) error {
	if interval == DRMEGLSwapIntervalVsync {
		return nil
	}
	return ErrImmediateSwapUnsupported
}

func (b *DRMEGLBackend) SetCursorPos(_, _ float64) {}

func (b *DRMEGLBackend) SetInputMode(_, _ DRMEGLWindowFlags) {}

func maxUint(a, floor uint) uint {
	if a < floor {
		return floor
	}
	return a
}

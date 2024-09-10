package ebitenbackend

import imgui "github.com/AllenDang/cimgui-go"

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

// SetDisplaySize sets the display size for imgui.
// This is usually called inside the game's Layout() function.
func (e *EbitenBackend) SetDisplaySize(width, height float32) {
	e.manager.width = width
	e.manager.height = height
}

// SetClipMask sets if clipmask is enabled or not. This is usually called for debugging purposes.
func (e *EbitenBackend) SetClipMask(value bool) {
	e.clipMask = value
}

// ClipMask returns if clipmask is enabled or not.
func (e *EbitenBackend) ClipMask() bool {
	return e.clipMask
}

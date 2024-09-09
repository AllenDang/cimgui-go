package ebitenbackend

// BeginFrame needs to be called on every frame, before cimgui-go calls.
// This is usually called inside the game's Update() function.
func (e *EbitenBackend) BeginFrame() {
	e.manager.BeginFrame()
}

// EndFrame needs to be called on every frame, after cimgui-go calls.
// This is usually called inside the game's Update() function.
func (e *EbitenBackend) EndFrame() {
	e.manager.EndFrame()
}

// SetDisplaySize sets the display size for imgui.
// This is usually called inside the game's Layout() function.
func (e *EbitenBackend) SetDisplaySize(width, height float32) {
	e.manager.SetDisplaySize(width, height)
}

// SetClipMask sets if clipmask is enabled or not. This is usually called for debugging purposes.
func (e *EbitenBackend) SetClipMask(value bool) {
	e.manager.ClipMask = value
}

// ClipMask returns if clipmask is enabled or not.
func (e *EbitenBackend) ClipMask() bool {
	return e.manager.ClipMask
}

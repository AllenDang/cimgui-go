package ebitenbackend

import "github.com/hajimehoshi/ebiten/v2"

var globalManager *Manager = NewManager(nil)

// Draw draws the generated imgui frame to the screen.
// This is usually called inside the game's Draw() function.
func Draw(screen *ebiten.Image) {
	globalManager.Draw(screen)
}

// Update needs to be called on every frame, before cimgui-go calls.
// This is usually called inside the game's Update() function.
// delta is the time in seconds since the last frame.
func Update(delta float32) {
	globalManager.Update(delta)
}

// BeginFrame needs to be called on every frame, before cimgui-go calls.
// This is usually called inside the game's Update() function.
func BeginFrame() {
	globalManager.BeginFrame()
}

// EndFrame needs to be called on every frame, after cimgui-go calls.
// This is usually called inside the game's Update() function.
func EndFrame() {
	globalManager.EndFrame()
}

// SetDisplaySize sets the display size for imgui.
// This is usually called inside the game's Layout() function.
func SetDisplaySize(width, height float32) {
	globalManager.SetDisplaySize(width, height)
}

// SetClipMask sets if clipmask is enabled or not. This is usually called for debugging purposes.
func SetClipMask(value bool) {
	globalManager.ClipMask = value
}

// ClipMask returns if clipmask is enabled or not.
func ClipMask() bool {
	return globalManager.ClipMask
}

// GlobalManager returns the global manager instance.
func GlobalManager() *Manager {
	return globalManager
}

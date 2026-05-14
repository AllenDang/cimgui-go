package raylibbackend

import (
	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Track touch state
var prevTouchDown bool

// GetPrevTouchDown returns whether touch was down in the previous frame
func GetPrevTouchDown() bool {
	return prevTouchDown
}

// sendRaylibTouch forwards primary touch as ImGui touch-screen mouse events
func sendRaylibTouch(io *imgui.IO) {
	// Number of active touchpoints
	count := rl.GetTouchPointCount()
	down := count > 0

	if down {
		// First finger position
		p := rl.GetTouchPosition(0)

		io.AddMouseSourceEvent(imgui.MouseSourceTouchScreen)
		io.AddMousePosEvent(p.X, p.Y)

		if !prevTouchDown {
			// Touch began -> send button down on left
			io.AddMouseButtonEvent(0, true)
		}
	} else {
		if prevTouchDown {
			// Touch ended -> send button up
			io.AddMouseButtonEvent(0, false)

			// Optionally revert source to Mouse
			io.AddMouseSourceEvent(imgui.MouseSourceMouse)
		}
	}

	prevTouchDown = down
}

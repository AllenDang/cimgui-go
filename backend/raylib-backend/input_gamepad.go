package raylibbackend

import (
	"math"

	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	// Previous button state
	prevPadBtn = map[int32]bool{}

	// Map ImGui gamepad keys to raylib buttons
	gamepadButtonMap = map[imgui.Key]int32{
		// Center buttons
		imgui.KeyGamepadStart: rl.GamepadButtonMiddleRight,
		imgui.KeyGamepadBack:  rl.GamepadButtonMiddleLeft,

		// Face buttons (raylib calls them right face)
		// A / Cross
		imgui.KeyGamepadFaceDown: rl.GamepadButtonRightFaceDown,
		// B / Circle
		imgui.KeyGamepadFaceRight: rl.GamepadButtonRightFaceRight,
		// X / Square
		imgui.KeyGamepadFaceLeft: rl.GamepadButtonRightFaceLeft,
		// Y / Triangle
		imgui.KeyGamepadFaceUp: rl.GamepadButtonRightFaceUp,

		// D-Pad (raylib calls them left face)
		imgui.KeyGamepadDpadLeft:  rl.GamepadButtonLeftFaceLeft,
		imgui.KeyGamepadDpadRight: rl.GamepadButtonLeftFaceRight,
		imgui.KeyGamepadDpadUp:    rl.GamepadButtonLeftFaceUp,
		imgui.KeyGamepadDpadDown:  rl.GamepadButtonLeftFaceDown,

		// Shoulders and stick-press
		imgui.KeyGamepadL1: rl.GamepadButtonLeftTrigger1,
		imgui.KeyGamepadL2: rl.GamepadButtonLeftTrigger2,
		imgui.KeyGamepadL3: rl.GamepadButtonLeftThumb,
		imgui.KeyGamepadR1: rl.GamepadButtonRightTrigger1,
		imgui.KeyGamepadR2: rl.GamepadButtonRightTrigger2,
		imgui.KeyGamepadR3: rl.GamepadButtonRightThumb,
	}
)

// sendRaylibGamepad forwards gamepad buttons and analog axes to ImGui
func sendRaylibGamepad(io *imgui.IO) {
	const pad = int32(0)

	has := rl.IsGamepadAvailable(pad)

	// Toggle BackendFlagsHasGamepad
	if has {
		io.SetBackendFlags(io.BackendFlags() | imgui.BackendFlagsHasGamepad)
	} else {
		io.SetBackendFlags(io.BackendFlags() &^ imgui.BackendFlagsHasGamepad)
	}

	// If no pad, also release any previously pressed buttons
	if !has {
		for key := range gamepadButtonMap {
			io.AddKeyEvent(key, false)
		}

		// Zero-out analogs
		zeroAnalog := func(k imgui.Key) { io.AddKeyAnalogEvent(k, false, 0) }
		zeroAnalog(imgui.KeyGamepadLStickLeft)
		zeroAnalog(imgui.KeyGamepadLStickRight)
		zeroAnalog(imgui.KeyGamepadLStickUp)
		zeroAnalog(imgui.KeyGamepadLStickDown)
		zeroAnalog(imgui.KeyGamepadRStickLeft)
		zeroAnalog(imgui.KeyGamepadRStickRight)
		zeroAnalog(imgui.KeyGamepadRStickUp)
		zeroAnalog(imgui.KeyGamepadRStickDown)
		zeroAnalog(imgui.KeyGamepadL2)
		zeroAnalog(imgui.KeyGamepadR2)

		return
	}

	// Buttons -> ImGui keys
	for k, rb := range gamepadButtonMap {
		down := rl.IsGamepadButtonDown(pad, rb)
		if prevPadBtn[rb] != down {
			io.AddKeyEvent(k, down)
			prevPadBtn[rb] = down
		}
	}

	// Axes -> analog keys
	const deadZone = 0.15

	// Left stick
	lx := rl.GetGamepadAxisMovement(pad, rl.GamepadAxisLeftX)
	ly := rl.GetGamepadAxisMovement(pad, rl.GamepadAxisLeftY)
	addAxisAnalogEvents(io, lx, deadZone, imgui.KeyGamepadLStickLeft, imgui.KeyGamepadLStickRight)
	addAxisAnalogEvents(io, ly, deadZone, imgui.KeyGamepadLStickUp, imgui.KeyGamepadLStickDown)

	// Right stick
	rx := rl.GetGamepadAxisMovement(pad, rl.GamepadAxisRightX)
	ry := rl.GetGamepadAxisMovement(pad, rl.GamepadAxisRightY)
	addAxisAnalogEvents(io, rx, deadZone, imgui.KeyGamepadRStickLeft, imgui.KeyGamepadRStickRight)
	addAxisAnalogEvents(io, ry, deadZone, imgui.KeyGamepadRStickUp, imgui.KeyGamepadRStickDown)

	// Triggers
	lt := rl.GetGamepadAxisMovement(pad, rl.GamepadAxisLeftTrigger)
	rt := rl.GetGamepadAxisMovement(pad, rl.GamepadAxisRightTrigger)

	norm := func(v float32) float32 {
		// Try to support both [-1..1] and [1..-1]
		// If v in [-1..1], (v+1)/2; if v in [1..-1], (1-v)/2 gives same mapping.
		// Blend using symmetrical formula: (1 - v) * 0.5 maps 1->0, -1->1.
		return clamp01((1 - v) * 0.5)
	}

	ltv := norm(lt)
	rtv := norm(rt)

	io.AddKeyAnalogEvent(imgui.KeyGamepadL2, ltv > deadZone, ltv)
	io.AddKeyAnalogEvent(imgui.KeyGamepadR2, rtv > deadZone, rtv)
}

// minf float32 min
func minf(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

// maxf float32 max
func maxf(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}

// clamp01 clamps a float32 value to [0, 1]
func clamp01(v float32) float32 {
	if v < 0 {
		return 0
	}
	if v > 1 {
		return 1
	}
	return v
}

// addAxisAnalogEvents processes axis and sends analog events
func addAxisAnalogEvents(io *imgui.IO, axis float32, deadZone float32, negKey, posKey imgui.Key) {
	io.AddKeyAnalogEvent(negKey, axis < -deadZone, clamp01(float32(math.Abs(float64(minf(axis, 0))))))
	io.AddKeyAnalogEvent(posKey, axis > deadZone, clamp01(float32(math.Abs(float64(maxf(axis, 0))))))
}

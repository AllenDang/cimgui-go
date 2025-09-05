package raylibbackend

type RaylibBackendFlags int

const (
	// RaylibBackendFlagsVsyncHint sets VSync hint for the GPU
	RaylibBackendFlagsVsyncHint RaylibBackendFlags = iota

	// RaylibBackendFlagsFullscreen sets fullscreen mode
	RaylibBackendFlagsFullscreen

	// RaylibBackendFlagsResizable sets window resizable
	RaylibBackendFlagsResizable

	// RaylibBackendFlagsUndecorated sets window without decorations
	RaylibBackendFlagsUndecorated

	// RaylibBackendFlagsHidden sets window hidden on initialization
	RaylibBackendFlagsHidden

	// RaylibBackendFlagsMinimized sets window minimized
	RaylibBackendFlagsMinimized

	// RaylibBackendFlagsMaximized sets window maximized
	RaylibBackendFlagsMaximized

	// RaylibBackendFlagsUnfocused sets window without focus on initialization
	RaylibBackendFlagsUnfocused

	// RaylibBackendFlagsTopmost sets window always on top
	RaylibBackendFlagsTopmost

	// RaylibBackendFlagsAlwaysRun sets window to keep running while minimized
	RaylibBackendFlagsAlwaysRun

	// RaylibBackendFlagsTransparent sets transparent framebuffer
	RaylibBackendFlagsTransparent

	// RaylibBackendFlagsHighDPI sets HighDPI support
	RaylibBackendFlagsHighDPI

	// RaylibBackendFlagsMousePassthrough sets mouse passthrough
	RaylibBackendFlagsMousePassthrough

	// RaylibBackendFlagsBorderlessWindowed sets borderless windowed mode
	RaylibBackendFlagsBorderlessWindowed

	// RaylibBackendFlagsMSAA4X sets MSAA 4X hint (antialiasing)
	RaylibBackendFlagsMSAA4X

	// RaylibBackendFlagsInterlaced sets interlaced hint for V3D (Raspberry Pi)
	RaylibBackendFlagsInterlaced

	// Input mode constants for SetInputMode

	// RaylibInputModeCursorVisible is the mode flag for cursor visibility
	RaylibInputModeCursorVisible

	// RaylibInputModeCursorEnabled is the mode flag for cursor enable/disable
	RaylibInputModeCursorEnabled

	// Input mode values for SetInputMode

	// RaylibInputModeDisabled represents a disabled input mode
	RaylibInputModeDisabled

	// RaylibInputModeEnabled represents an enabled input mode
	RaylibInputModeEnabled

	// Swap interval values for SetSwapInterval

	// RaylibSwapIntervalDisabled disables VSync (swap interval = 0)
	RaylibSwapIntervalDisabled

	// RaylibSwapIntervalEnabled enables VSync (swap interval = 1)
	RaylibSwapIntervalEnabled
)

package ebitenbackend

type EbitenBackendFlags int

const (
	// EbitenBackendFlagsCursorMode sets the cursor mode.
	// refer ebiten.CursorModeType
	// CursorModeVisible, CursorModeHidden, CursorModeCaptured
	EbitenBackendFlagsCursorMode EbitenBackendFlags = iota
	// EbitenBackendFlagsCursorShape sets the cursor shape.
	// refer ebiten.CursorShapeType
	// CursorShapeDefault, CursorShapeText, CursorShapeCrosshair, CursorShapePointer,
	// CursorShapeEWResize, CursorShapeNSResize, CursorShapeNESWResize, CursorShapeNWSEResize,
	// CursorShapeMove, CursorShapeNotAllowed CursorShapeType, EbitenBackendFlagsCursorShape
	EbitenBackendFlagsCursorShape
	// EbitenBackendFlagsResizingMode sets the resizing mode.
	// Possible values: WindowResizingModeDisabled, WindowResizingModeOnlyFullscreenEnabled, WindowResizingModeEnabled
	EbitenBackendFlagsResizingMode
	// EbitenBackendFlagsFPSMode if 0 set enables vsync. See also: ebiten.SetVsyncEnabled
	EbitenBackendFlagsFPSMode
	// EbitenBackendFlagsDecorated sets the window decoration if != 0.
	EbitenBackendFlagsDecorated
	EbitenBackendFlagsFloating
	EbitenBackendFlagsMaximized
	EbitenBackendFlagsMinimized
	EbitenBackendFlagsClosingHandled
	EbitenBackendFlagsMousePassthrough
	// EbitenBackendFlagsDebug is a flag to enable debug mode. It will show FPS, TPS, ClipMask and enable ClipMask shortcut.
	// 0 (default) disabled, 1 (or anything else) enabled
	// The following keys are bound to the following actions:
	// - C: Toggle ClipMask
	// - I: Toggle input sync (if disabled no input will be handled)
	// - S: Toggle cursor shape sync
	EbitenBackendFlagsDebug
)

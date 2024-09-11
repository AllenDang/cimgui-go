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
	EbitenBackendFlagsFPSMode
	EbitenBackendFlagsDecorated
	EbitenBackendFlagsFloating
	EbitenBackendFlagsMaximized
	EbitenBackendFlagsMinimized
	EbitenBackendFlagsClosingHandled
	EbitenBackendFlagsMousePassthrough
	// EbitenBackendFlagsDebug is a flag to enable debug mode. It will show FPS, TPS, ClipMask and enable ClipMask shortcut.
	// 0 (default) disabled, 1 (or anything else) enabled
	EbitenBackendFlagsDebug
)

package ebitenbackend

type GetCursorFn func() (x, y float32)

type Manager struct {
	width        float32
	height       float32
	screenWidth  int
	screenHeight int
}

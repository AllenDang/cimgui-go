package ebitenbackend

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func keysMap() map[imgui.Key]int {
	return map[imgui.Key]int{
		imgui.KeyTab:        int(ebiten.KeyTab),
		imgui.KeyLeftArrow:  int(ebiten.KeyLeft),
		imgui.KeyRightArrow: int(ebiten.KeyRight),
		imgui.KeyUpArrow:    int(ebiten.KeyUp),
		imgui.KeyDownArrow:  int(ebiten.KeyDown),
		imgui.KeyPageUp:     int(ebiten.KeyPageUp),
		imgui.KeyPageDown:   int(ebiten.KeyPageDown),
		imgui.KeyHome:       int(ebiten.KeyHome),
		imgui.KeyEnd:        int(ebiten.KeyEnd),
		imgui.KeyInsert:     int(ebiten.KeyInsert),
		imgui.KeyDelete:     int(ebiten.KeyDelete),
		imgui.KeyBackspace:  int(ebiten.KeyBackspace),
		imgui.KeySpace:      int(ebiten.KeySpace),
		imgui.KeyEnter:      int(ebiten.KeyEnter),
		imgui.KeyEscape:     int(ebiten.KeyEscape),
		imgui.KeyA:          int(ebiten.KeyA),
		imgui.KeyC:          int(ebiten.KeyC),
		imgui.KeyV:          int(ebiten.KeyV),
		imgui.KeyX:          int(ebiten.KeyX),
		imgui.KeyY:          int(ebiten.KeyY),
		imgui.KeyZ:          int(ebiten.KeyZ),
	}
}

func sendInput(io *imgui.IO, inputChars []rune) []rune {
	// Ebiten hides the LeftAlt RightAlt implementation (inside the uiDriver()), so
	// here only the left alt is sent
	if ebiten.IsKeyPressed(ebiten.KeyAlt) {
		io.SetKeyAlt(true)
	} else {
		io.SetKeyAlt(false)
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		io.SetKeyShift(true)
	} else {
		io.SetKeyShift(false)
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		io.SetKeyCtrl(true)
	} else {
		io.SetKeyCtrl(false)
	}
	// TODO: get KeySuper somehow (GLFW: KeyLeftSuper    = Key(343), R: 347)
	inputChars = ebiten.AppendInputChars(inputChars)
	if len(inputChars) > 0 {
		io.AddInputCharactersUTF8(string(inputChars))
		inputChars = inputChars[:0]
	}
	for ik, iv := range keysMap() {
		if inpututil.IsKeyJustPressed(ebiten.Key(iv)) {
			io.AddKeyEvent(ik, true)
		}
		if inpututil.IsKeyJustReleased(ebiten.Key(iv)) {
			io.AddKeyEvent(ik, false)
		}
	}
	return inputChars
}

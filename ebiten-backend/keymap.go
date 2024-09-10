package ebitenbackend

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func keyMap() map[imgui.Key]int {
	return map[imgui.Key]int{
		imgui.KeyTab:          int(ebiten.KeyTab),
		imgui.KeyCapsLock:     int(ebiten.KeyCapsLock),
		imgui.KeyComma:        int(ebiten.KeyComma),
		imgui.KeyMenu:         int(ebiten.KeyContextMenu),
		imgui.KeyLeftArrow:    int(ebiten.KeyArrowLeft),
		imgui.KeyRightArrow:   int(ebiten.KeyArrowRight),
		imgui.KeyUpArrow:      int(ebiten.KeyArrowUp),
		imgui.KeyDownArrow:    int(ebiten.KeyArrowDown),
		imgui.KeyBackslash:    int(ebiten.KeyBackslash),
		imgui.KeyBackspace:    int(ebiten.KeyBackspace),
		imgui.KeyLeftCtrl:     int(ebiten.KeyControlLeft),
		imgui.KeyRightCtrl:    int(ebiten.KeyControlRight),
		imgui.KeyLeftAlt:      int(ebiten.KeyAltLeft),
		imgui.KeyRightAlt:     int(ebiten.KeyAltRight),
		imgui.KeyLeftShift:    int(ebiten.KeyShiftLeft),
		imgui.KeyRightShift:   int(ebiten.KeyShiftRight),
		imgui.KeyLeftSuper:    int(ebiten.KeyMetaLeft),
		imgui.KeyRightSuper:   int(ebiten.KeyMetaRight),
		imgui.KeyLeftBracket:  int(ebiten.KeyBracketLeft),
		imgui.KeyRightBracket: int(ebiten.KeyBracketRight),
		imgui.KeyPageUp:       int(ebiten.KeyPageUp),
		imgui.KeyPageDown:     int(ebiten.KeyPageDown),
		imgui.KeyEnd:          int(ebiten.KeyEnd),
		imgui.KeyHome:         int(ebiten.KeyHome),
		imgui.KeyInsert:       int(ebiten.KeyInsert),
		imgui.KeyDelete:       int(ebiten.KeyDelete),
		imgui.KeySpace:        int(ebiten.KeySpace),
		imgui.KeyEnter:        int(ebiten.KeyEnter),
		imgui.KeyEscape:       int(ebiten.KeyEscape),
		imgui.KeyEqual:        int(ebiten.KeyEqual),
		imgui.Key0:            int(ebiten.KeyDigit0),
		imgui.Key1:            int(ebiten.KeyDigit1),
		imgui.Key2:            int(ebiten.KeyDigit2),
		imgui.Key3:            int(ebiten.KeyDigit3),
		imgui.Key4:            int(ebiten.KeyDigit4),
		imgui.Key5:            int(ebiten.KeyDigit5),
		imgui.Key6:            int(ebiten.KeyDigit6),
		imgui.Key7:            int(ebiten.KeyDigit7),
		imgui.Key8:            int(ebiten.KeyDigit8),
		imgui.Key9:            int(ebiten.KeyDigit9),
		imgui.KeyA:            int(ebiten.KeyA),
		imgui.KeyB:            int(ebiten.KeyB),
		imgui.KeyC:            int(ebiten.KeyC),
		imgui.KeyD:            int(ebiten.KeyD),
		imgui.KeyE:            int(ebiten.KeyE),
		imgui.KeyF:            int(ebiten.KeyF),
		imgui.KeyG:            int(ebiten.KeyG),
		imgui.KeyH:            int(ebiten.KeyH),
		imgui.KeyI:            int(ebiten.KeyI),
		imgui.KeyJ:            int(ebiten.KeyJ),
		imgui.KeyK:            int(ebiten.KeyK),
		imgui.KeyL:            int(ebiten.KeyL),
		imgui.KeyM:            int(ebiten.KeyM),
		imgui.KeyN:            int(ebiten.KeyN),
		imgui.KeyO:            int(ebiten.KeyO),
		imgui.KeyP:            int(ebiten.KeyP),
		imgui.KeyQ:            int(ebiten.KeyQ),
		imgui.KeyR:            int(ebiten.KeyR),
		imgui.KeyS:            int(ebiten.KeyS),
		imgui.KeyT:            int(ebiten.KeyT),
		imgui.KeyU:            int(ebiten.KeyU),
		imgui.KeyV:            int(ebiten.KeyV),
		imgui.KeyW:            int(ebiten.KeyW),
		imgui.KeyX:            int(ebiten.KeyX),
		imgui.KeyY:            int(ebiten.KeyY),
		imgui.KeyZ:            int(ebiten.KeyZ),
		imgui.KeyKeypad0:      int(ebiten.KeyNumpad0),
		imgui.KeyKeypad1:      int(ebiten.KeyNumpad1),
		imgui.KeyKeypad2:      int(ebiten.KeyNumpad2),
		imgui.KeyKeypad3:      int(ebiten.KeyNumpad3),
		imgui.KeyKeypad4:      int(ebiten.KeyNumpad4),
		imgui.KeyKeypad5:      int(ebiten.KeyNumpad5),
		imgui.KeyKeypad6:      int(ebiten.KeyNumpad6),
		imgui.KeyKeypad7:      int(ebiten.KeyNumpad7),
		imgui.KeyKeypad8:      int(ebiten.KeyNumpad8),
		imgui.KeyKeypad9:      int(ebiten.KeyNumpad9),
		imgui.KeyF1:           int(ebiten.KeyF1),
		imgui.KeyF2:           int(ebiten.KeyF2),
		imgui.KeyF3:           int(ebiten.KeyF3),
		imgui.KeyF4:           int(ebiten.KeyF4),
		imgui.KeyF5:           int(ebiten.KeyF5),
		imgui.KeyF6:           int(ebiten.KeyF6),
		imgui.KeyF7:           int(ebiten.KeyF7),
		imgui.KeyF8:           int(ebiten.KeyF8),
		imgui.KeyF9:           int(ebiten.KeyF9),
		imgui.KeyF10:          int(ebiten.KeyF10),
		imgui.KeyF11:          int(ebiten.KeyF11),
		imgui.KeyF12:          int(ebiten.KeyF12),
		imgui.KeyF13:          int(ebiten.KeyF13),
		imgui.KeyF14:          int(ebiten.KeyF14),
		imgui.KeyF15:          int(ebiten.KeyF15),
		imgui.KeyF16:          int(ebiten.KeyF16),
		imgui.KeyF17:          int(ebiten.KeyF17),
		imgui.KeyF18:          int(ebiten.KeyF18),
		imgui.KeyF19:          int(ebiten.KeyF19),
		imgui.KeyF20:          int(ebiten.KeyF20),
		imgui.KeyF21:          int(ebiten.KeyF21),
		imgui.KeyF22:          int(ebiten.KeyF22),
		imgui.KeyF23:          int(ebiten.KeyF23),
		imgui.KeyF24:          int(ebiten.KeyF24),
		// TODO: ebiten.KeyApostrophe was deprecated as of v2.1
		imgui.KeyApostrophe:     int(ebiten.KeyApostrophe),
		imgui.KeyMinus:          int(ebiten.KeyMinus),
		imgui.KeyPeriod:         int(ebiten.KeyPeriod),
		imgui.KeySlash:          int(ebiten.KeySlash),
		imgui.KeySemicolon:      int(ebiten.KeySemicolon),
		imgui.KeyGraveAccent:    int(ebiten.KeyBackquote),
		imgui.KeyScrollLock:     int(ebiten.KeyScrollLock),
		imgui.KeyNumLock:        int(ebiten.KeyNumLock),
		imgui.KeyPrintScreen:    int(ebiten.KeyPrintScreen),
		imgui.KeyPause:          int(ebiten.KeyPause),
		imgui.KeyKeypadDecimal:  int(ebiten.KeyKPDecimal),
		imgui.KeyKeypadDivide:   int(ebiten.KeyKPDivide),
		imgui.KeyKeypadMultiply: int(ebiten.KeyKPMultiply),
		imgui.KeyKeypadSubtract: int(ebiten.KeyKPSubtract),
		imgui.KeyKeypadAdd:      int(ebiten.KeyKPAdd),
		imgui.KeyKeypadEnter:    int(ebiten.KeyKPEnter),
		imgui.KeyKeypadEqual:    int(ebiten.KeyKPEqual),
	}
}

func sendInput(io *imgui.IO, inputChars []rune) []rune {
	// TODO: Is this still necessary? Ebiten now has both Left and Right key codes.
	if ebiten.IsKeyPressed(ebiten.KeyAltRight) {
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
	inputChars = ebiten.AppendInputChars(inputChars)
	if len(inputChars) > 0 {
		io.AddInputCharactersUTF8(string(inputChars))
		inputChars = inputChars[:0]
	}
	for ik, iv := range keyMap() {
		if inpututil.IsKeyJustPressed(ebiten.Key(iv)) {
			io.AddKeyEvent(ik, true)
		}
		if inpututil.IsKeyJustReleased(ebiten.Key(iv)) {
			io.AddKeyEvent(ik, false)
		}
	}
	return inputChars
}

func (e *EbitenBackend) setKeyMapping() {
	// Keyboard mapping. ImGui will use those indices to peek into the io.KeysDown[] array.
	/*
		io := cimgui.GetIO()
		for imguiKey, nativeKey := range keys {
			// io.KeyMap(int(imguiKey), nativeKey)
		}
	*/
}

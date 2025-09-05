package raylibbackend

import (
    "github.com/AllenDang/cimgui-go/imgui"
    rl "github.com/gen2brain/raylib-go/raylib"
)

// Key mapping from ImGui virtual keys to raylib key codes.
// Filled once; used each frame to generate edge events.
func rlKeyMap() map[imgui.Key]int32 {
    return map[imgui.Key]int32{
        // Letters
        imgui.KeyA: int32(rl.KeyA), imgui.KeyB: int32(rl.KeyB), imgui.KeyC: int32(rl.KeyC), imgui.KeyD: int32(rl.KeyD),
        imgui.KeyE: int32(rl.KeyE), imgui.KeyF: int32(rl.KeyF), imgui.KeyG: int32(rl.KeyG), imgui.KeyH: int32(rl.KeyH),
        imgui.KeyI: int32(rl.KeyI), imgui.KeyJ: int32(rl.KeyJ), imgui.KeyK: int32(rl.KeyK), imgui.KeyL: int32(rl.KeyL),
        imgui.KeyM: int32(rl.KeyM), imgui.KeyN: int32(rl.KeyN), imgui.KeyO: int32(rl.KeyO), imgui.KeyP: int32(rl.KeyP),
        imgui.KeyQ: int32(rl.KeyQ), imgui.KeyR: int32(rl.KeyR), imgui.KeyS: int32(rl.KeyS), imgui.KeyT: int32(rl.KeyT),
        imgui.KeyU: int32(rl.KeyU), imgui.KeyV: int32(rl.KeyV), imgui.KeyW: int32(rl.KeyW), imgui.KeyX: int32(rl.KeyX),
        imgui.KeyY: int32(rl.KeyY), imgui.KeyZ: int32(rl.KeyZ),

        // Numbers row
        imgui.Key0: int32(rl.KeyZero), imgui.Key1: int32(rl.KeyOne), imgui.Key2: int32(rl.KeyTwo), imgui.Key3: int32(rl.KeyThree),
        imgui.Key4: int32(rl.KeyFour), imgui.Key5: int32(rl.KeyFive), imgui.Key6: int32(rl.KeySix), imgui.Key7: int32(rl.KeySeven),
        imgui.Key8: int32(rl.KeyEight), imgui.Key9: int32(rl.KeyNine),

        // Punctuation
        imgui.KeyMinus: int32(rl.KeyMinus),
        imgui.KeyEqual: int32(rl.KeyEqual),
        imgui.KeyBackslash: int32(rl.KeyBackSlash),
        imgui.KeySlash: int32(rl.KeySlash),
        imgui.KeyPeriod: int32(rl.KeyPeriod),
        imgui.KeyComma: int32(rl.KeyComma),
        imgui.KeySemicolon: int32(rl.KeySemicolon),
        imgui.KeyGraveAccent: int32(rl.KeyGrave),
        imgui.KeyLeftBracket: int32(rl.KeyLeftBracket),
        imgui.KeyRightBracket: int32(rl.KeyRightBracket),
        imgui.KeyApostrophe: int32(rl.KeyApostrophe),

        // Navigation
        imgui.KeyTab: int32(rl.KeyTab),
        imgui.KeyEnter: int32(rl.KeyEnter),
        imgui.KeyEscape: int32(rl.KeyEscape),
        imgui.KeySpace: int32(rl.KeySpace),
        imgui.KeyBackspace: int32(rl.KeyBackspace),
        imgui.KeyLeftArrow: int32(rl.KeyLeft),
        imgui.KeyRightArrow: int32(rl.KeyRight),
        imgui.KeyUpArrow: int32(rl.KeyUp),
        imgui.KeyDownArrow: int32(rl.KeyDown),
        imgui.KeyHome: int32(rl.KeyHome),
        imgui.KeyEnd: int32(rl.KeyEnd),
        imgui.KeyInsert: int32(rl.KeyInsert),
        imgui.KeyDelete: int32(rl.KeyDelete),
        imgui.KeyPageUp: int32(rl.KeyPageUp),
        imgui.KeyPageDown: int32(rl.KeyPageDown),

        // Function keys
        imgui.KeyF1: int32(rl.KeyF1), imgui.KeyF2: int32(rl.KeyF2), imgui.KeyF3: int32(rl.KeyF3), imgui.KeyF4: int32(rl.KeyF4),
        imgui.KeyF5: int32(rl.KeyF5), imgui.KeyF6: int32(rl.KeyF6), imgui.KeyF7: int32(rl.KeyF7), imgui.KeyF8: int32(rl.KeyF8),
        imgui.KeyF9: int32(rl.KeyF9), imgui.KeyF10: int32(rl.KeyF10), imgui.KeyF11: int32(rl.KeyF11), imgui.KeyF12: int32(rl.KeyF12),

        // Explicit left/right modifiers so ImGui can query them as keys
        imgui.KeyLeftCtrl:  int32(rl.KeyLeftControl),
        imgui.KeyRightCtrl: int32(rl.KeyRightControl),
        imgui.KeyLeftShift: int32(rl.KeyLeftShift),
        imgui.KeyRightShift:int32(rl.KeyRightShift),
        imgui.KeyLeftAlt:   int32(rl.KeyLeftAlt),
        imgui.KeyRightAlt:  int32(rl.KeyRightAlt),
        imgui.KeyLeftSuper: int32(rl.KeyLeftSuper),
        imgui.KeyRightSuper:int32(rl.KeyRightSuper),
    }
}

// Track previous key state for edge detection (press/release)
var rlPrevDown = map[int32]bool{}
// Track previous state for ImGui's modifier key-chord virtual keys
var prevModDown = map[imgui.Key]bool{}

// sendRaylibInput feeds keyboard modifiers, key edge events, and UTF-8 text input to ImGui.
func sendRaylibInput(io *imgui.IO) {
    // Modifiers
    ctrlDown := rl.IsKeyDown(rl.KeyLeftControl) || rl.IsKeyDown(rl.KeyRightControl)
    shiftDown := rl.IsKeyDown(rl.KeyLeftShift) || rl.IsKeyDown(rl.KeyRightShift)
    altDown := rl.IsKeyDown(rl.KeyLeftAlt) || rl.IsKeyDown(rl.KeyRightAlt)
    superDown := rl.IsKeyDown(rl.KeyLeftSuper) || rl.IsKeyDown(rl.KeyRightSuper)

    // Keep legacy IO flags in sync (some helper APIs still read these)
    io.SetKeyCtrl(ctrlDown)
    io.SetKeyShift(shiftDown)
    io.SetKeyAlt(altDown)
    io.SetKeySuper(superDown)

    // Also report modifiers as ImGui key events so IsKeyDown() on mod keys works.
    // Note: bindings expose reserved keys for key-chords.
    setMod := func(k imgui.Key, down bool) {
        if prevModDown[k] != down {
            io.AddKeyEvent(k, down)
            prevModDown[k] = down
        }
    }
    setMod(imgui.KeyReservedForModCtrl, ctrlDown)
    setMod(imgui.KeyReservedForModShift, shiftDown)
    setMod(imgui.KeyReservedForModAlt, altDown)
    setMod(imgui.KeyReservedForModSuper, superDown)

    // Keys edge events
    for ik, rk := range rlKeyMap() {
        down := rl.IsKeyDown(rk)
        was := rlPrevDown[rk]
        if down != was {
            io.AddKeyEvent(ik, down)
            rlPrevDown[rk] = down
        }
    }

    // UTF-8 text input (queue of codepoints this frame)
    for {
        cp := rl.GetCharPressed()
        if cp == 0 { // no more chars
            break
        }
        io.AddInputCharactersUTF8(string(rune(cp)))
    }
}

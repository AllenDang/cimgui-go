package ebitenbackend

import (
	imgui "github.com/AllenDang/cimgui-go"
)

type GetCursorFn func() (x, y float32)

type Manager struct {
	cliptxt   string
	GetCursor GetCursorFn

	width        float32
	height       float32
	screenWidth  int
	screenHeight int

	inputChars []rune
}

func NewManager(fontAtlas *imgui.FontAtlas) *Manager {
	m := &Manager{
		inputChars: make([]rune, 0, 256),
	}

	return m
}

func NewManagerWithContext(ctx *imgui.Context) *Manager {
	m := &Manager{}

	return m
}

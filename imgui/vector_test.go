package imgui

import "testing"

// A vector of pointers gives every element through Slice: once the font
// atlas moves to a new texture, the draw data and the platform IO both list
// the old texture and the new one (#492, #554).
func TestVectorOfPointersSlice(t *testing.T) {
	ctx := CreateContext()
	defer DestroyContextV(ctx)
	io := CurrentIO()
	io.SetIniFilename("") // no imgui.ini
	io.SetBackendFlags(io.BackendFlags() | BackendFlagsRendererHasTextures)
	io.SetDisplaySize(Vec2{X: 640, Y: 360})

	// frame draws text at size, when not 0, and completes the textures'
	// requests as a backend would, reading them through Slice.
	frame := func(size float32) {
		NewFrame()
		if size != 0 {
			Begin("text")
			PushFont(CurrentFont(), size)
			Text("The quick brown fox jumps over the lazy dog")
			PopFont()
			End()
		}
		Render()
		for _, tex := range CurrentPlatformIO().Textures().Slice() {
			if tex.Status() == TextureStatusWantCreate {
				tex.SetTexID(TextureID(tex.UniqueID() + 1))
			}
			tex.SetStatus(TextureStatusOK)
		}
	}

	frame(40)
	if got := CurrentPlatformIO().Textures().Slice(); len(got) != 1 || got[0].CData != io.Fonts().TexData().CData {
		t.Fatalf("first frame: %d textures, want the atlas's alone", len(got))
	}
	frame(0)
	frame(0)
	io.Fonts().CompactCache() // drops the unused size, moving the atlas to a new texture
	frame(0)
	for _, list := range []struct {
		name     string
		textures []TextureData
	}{
		{"platform IO", CurrentPlatformIO().Textures().Slice()},
		{"draw data", CurrentDrawData().Textures().Slice()},
	} {
		if len(list.textures) != 2 {
			t.Fatalf("%s: %d textures after the atlas moved, want 2", list.name, len(list.textures))
		}
		for i, tex := range list.textures {
			if tex.CData == nil {
				t.Errorf("%s: texture %d is nil", list.name, i)
			}
		}
		if list.textures[1].CData != io.Fonts().TexData().CData {
			t.Errorf("%s: the second texture is not the atlas's new one", list.name)
		}
	}
}

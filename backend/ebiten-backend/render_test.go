package ebitenbackend

import (
	"testing"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/hajimehoshi/ebiten/v2"
)

// The imgui functions that wrap a C pointer in a Go value are called through
// these function values: with Go 1.25, go test -race fails to link a test
// package outside imgui that inlines them ("inlined function
// internal.ReinterpretCast[...] missing func info").
var (
	createContext     = imgui.CreateContext
	currentIO         = imgui.CurrentIO
	currentFont       = imgui.CurrentFont
	currentDrawData   = imgui.CurrentDrawData
	currentPlatformIO = imgui.CurrentPlatformIO
)

func TestAppendTexturesListsEveryTexture(t *testing.T) {
	ctx := createContext()
	defer imgui.DestroyContextV(ctx)
	io := currentIO()
	io.SetIniFilename("") // no imgui.ini
	io.SetBackendFlags(io.BackendFlags() | imgui.BackendFlagsRendererHasTextures)
	io.SetDisplaySize(imgui.Vec2{X: 640, Y: 360})

	// frame draws text at size, when not 0, and completes the textures'
	// requests as a backend would.
	frame := func(size float32) []imgui.TextureData {
		imgui.NewFrame()
		if size != 0 {
			imgui.Begin("text")
			imgui.PushFont(currentFont(), size)
			imgui.Text("The quick brown fox jumps over the lazy dog")
			imgui.PopFont()
			imgui.End()
		}
		imgui.Render()
		textures := appendTextures(currentDrawData(), nil)
		for _, tex := range textures {
			if tex.Status() == imgui.TextureStatusWantCreate {
				tex.SetTexID(imgui.TextureID(tex.UniqueID() + 1))
			}
			tex.SetStatus(imgui.TextureStatusOK)
		}
		return textures
	}

	first := frame(40)
	if len(first) != 1 || first[0].CData != io.Fonts().TexData().CData {
		t.Fatalf("first frame: got %d textures, want the atlas alone", len(first))
	}
	frame(0)
	frame(0)
	io.Fonts().CompactCache() // drops the unused size, moving the atlas to a new texture
	got := frame(0)
	want := []imgui.TextureData{first[0], *io.Fonts().TexData()}
	if len(got) != len(want) {
		t.Fatalf("after the atlas moved: got %d textures, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i].CData != want[i].CData {
			t.Errorf("texture %d: got #%d, want #%d", i, got[i].UniqueID(), want[i].UniqueID())
		}
	}
	if n := currentPlatformIO().Textures().Size; int(n) != len(got) {
		t.Errorf("got %d textures, ImGui lists %d", len(got), n)
	}
}

// Render does what ImGui asks of every texture it lists, creating, updating
// and destroying them, the second one included once the font atlas moves to
// a new texture: after each frame, a texture ImGui keeps has an image of its
// size in the cache, a destroyed one has none, and the cache holds nothing
// else.
func TestRenderKeepsTheCacheToImGuisTextures(t *testing.T) {
	ctx := createContext()
	defer imgui.DestroyContextV(ctx)
	io := currentIO()
	io.SetIniFilename("") // no imgui.ini
	io.SetBackendFlags(io.BackendFlags() | imgui.BackendFlagsRendererHasTextures)
	io.SetDisplaySize(imgui.Vec2{X: 640, Y: 360})

	cache := NewCache()
	target := ebiten.NewImage(640, 360)
	most := 0                               // the most textures listed at once
	asked := map[imgui.TextureStatus]bool{} // the requests ImGui made
	frame := func(size float32) {
		t.Helper()
		imgui.NewFrame()
		if size != 0 {
			imgui.Begin("text")
			imgui.PushFont(currentFont(), size)
			imgui.Text("The quick brown fox jumps over the lazy dog")
			imgui.PopFont()
			imgui.End()
		}
		imgui.Render()
		for _, tex := range appendTextures(currentDrawData(), nil) {
			asked[tex.Status()] = true
		}
		Render(target, currentDrawData(), cache, ebiten.FilterNearest)

		textures := appendTextures(currentDrawData(), nil)
		most = max(most, len(textures))
		kept := 0
		for _, tex := range textures {
			switch tex.Status() {
			case imgui.TextureStatusOK:
				kept++
				img := cache.GetTexture(tex.TexID())
				if img == nil || img.Bounds().Dx() != int(tex.Width()) || img.Bounds().Dy() != int(tex.Height()) {
					t.Fatalf("texture #%d, kept, has no image of its size in the cache", tex.UniqueID())
				}
			case imgui.TextureStatusDestroyed:
				if tex.TexID() != 0 {
					t.Fatalf("texture #%d, destroyed, keeps ID %d", tex.UniqueID(), tex.TexID())
				}
			default:
				t.Fatalf("texture #%d still asks for %v after rendering", tex.UniqueID(), tex.Status())
			}
		}
		if n := len(cache.(*textureCache).cache); n != kept {
			t.Fatalf("the cache holds %d images, ImGui keeps %d textures", n, kept)
		}
	}

	frame(40)
	frame(30) // the same text at another size: new glyphs written into the atlas
	frame(0)
	io.Fonts().CompactCache() // drops the unused size, moving the atlas to a new texture
	for range 10 {
		frame(0)
	}
	if most < 2 {
		t.Fatalf("the atlas never moved: at most %d textures listed", most)
	}
	for _, s := range []imgui.TextureStatus{imgui.TextureStatusWantCreate, imgui.TextureStatusWantUpdates, imgui.TextureStatusWantDestroy} {
		if !asked[s] {
			t.Errorf("ImGui never asked for %v", s)
		}
	}
}

package ebitenbackend_test

import (
	"testing"

	ebitenbackend "github.com/AllenDang/cimgui-go/backend/ebiten-backend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/hajimehoshi/ebiten/v2"
)

func TestTextureCacheNextIdSkipsRemovedIDs(t *testing.T) {
	cache := ebitenbackend.NewCache()
	a := imgui.TextureID(cache.NextId())
	cache.SetTexture(a, ebiten.NewImage(1, 1))
	b := imgui.TextureID(cache.NextId())
	cache.SetTexture(b, ebiten.NewImage(1, 1))
	cache.RemoveTexture(a)
	if c := imgui.TextureID(cache.NextId()); c == a || c == b {
		t.Errorf("NextId gives %d after %d and %d, with %d removed", c, a, b, a)
	}
}

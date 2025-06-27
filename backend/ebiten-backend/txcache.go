package ebitenbackend

import (
	"log"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/hajimehoshi/ebiten/v2"
)

type TextureCache interface {
	UpdateTexture(tex imgui.TextureData)
	GetTexture(id imgui.TextureID) *ebiten.Image
	GetGameTexture(id imgui.TextureID) (ebiten.Game, bool)
	ForEachGame(f func(id imgui.TextureID, game ebiten.Game, target *ebiten.Image))
	SetTexture(id imgui.TextureID, img *ebiten.Image)
	SetGameTexture(id imgui.TextureID, img ebiten.Game)
	RemoveTexture(id imgui.TextureID)
	NextId() int
}

type textureCache struct {
	startIndex int
	cache      map[imgui.TextureID]*ebiten.Image
	cacheGame  map[imgui.TextureID]ebiten.Game
	dfilter    ebiten.Filter
}

var _ TextureCache = (*textureCache)(nil)

func (c *textureCache) NextId() int {
	return len(c.cache) + c.startIndex
}

func (c *textureCache) UpdateTexture(tex imgui.TextureData) {
	switch tex.Status() {
	case imgui.TextureStatusOK:
		// noop.
	case imgui.TextureStatusWantCreate:
		texImage := getTexture(unsafe.Pointer(tex.Pixels()), tex.Width(), tex.Height())
		newID := imgui.TextureID(c.NextId())
		c.SetTexture(newID, texImage)
		tex.SetTexID(newID)
		tex.SetStatus(imgui.TextureStatusOK)
	case imgui.TextureStatusWantDestroy:
		c.RemoveTexture(tex.TexID())
	case imgui.TextureStatusWantUpdates:
		texImage := getTexture(unsafe.Pointer(tex.Pixels()), tex.Width(), tex.Height())
		c.SetTexture(tex.TexID(), texImage)
	default:
		log.Panicf("Unknown texture status: %v", tex.Status())
	}
}

func (c *textureCache) GetTexture(id imgui.TextureID) *ebiten.Image {
	if im, ok := c.cache[id]; ok {
		return im
	}

	panic("Texture not found in cache. Cannot obtain ebiten Image")
}

func (c *textureCache) GetGameTexture(id imgui.TextureID) (res ebiten.Game, isGame bool) {
	if im, ok := c.cacheGame[id]; ok {
		return im, true
	}

	return nil, false
}

func (c *textureCache) SetTexture(id imgui.TextureID, img *ebiten.Image) {
	c.cache[id] = img
}

func (c *textureCache) SetGameTexture(id imgui.TextureID, img ebiten.Game) {
	c.cacheGame[id] = img
}

func (c *textureCache) ForEachGame(f func(id imgui.TextureID, game ebiten.Game, target *ebiten.Image)) {
	for id, game := range c.cacheGame {
		target := c.GetTexture(id)
		f(id, game, target)
	}
}

func (c *textureCache) RemoveTexture(id imgui.TextureID) {
	delete(c.cache, id)
}

func NewCache() TextureCache {
	return &textureCache{
		startIndex: 2,
		cache:      make(map[imgui.TextureID]*ebiten.Image),
		cacheGame:  make(map[imgui.TextureID]ebiten.Game),
	}
}

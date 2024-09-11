package ebitenbackend

import (
	imgui "github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
)

type TextureCache interface {
	FontAtlasTextureID() imgui.TextureID
	SetFontAtlasTextureID(id imgui.TextureID)
	GetTexture(id imgui.TextureID) *ebiten.Image
	SetTexture(id imgui.TextureID, img *ebiten.Image)
	RemoveTexture(id imgui.TextureID)
	ResetFontAtlasCache(filter ebiten.Filter)
	NextId() int
}

type textureCache struct {
	startIndex     int
	fontAtlasID    imgui.TextureID
	fontAtlasImage *ebiten.Image
	cache          map[imgui.TextureID]*ebiten.Image
	dfilter        ebiten.Filter
}

var _ TextureCache = (*textureCache)(nil)

func (c *textureCache) NextId() int {
	return len(c.cache) + c.startIndex
}

func (c *textureCache) getFontAtlas() *ebiten.Image {
	if c.fontAtlasImage == nil {
		pixels, width, height, _ := imgui.CurrentIO().Fonts().GetTextureDataAsRGBA32()
		c.fontAtlasImage = getTexture(pixels, width, height)
		c.SetTexture(c.fontAtlasID, c.fontAtlasImage)
		imgui.CurrentIO().Fonts().SetTexID(c.fontAtlasID)
	}

	return c.fontAtlasImage
}

func (c *textureCache) FontAtlasTextureID() imgui.TextureID {
	return c.fontAtlasID
}

func (c *textureCache) SetFontAtlasTextureID(id imgui.TextureID) {
	c.fontAtlasID = id
}

func (c *textureCache) GetTexture(id imgui.TextureID) *ebiten.Image {
	if id == c.fontAtlasID {
		return c.getFontAtlas()
	}

	if im, ok := c.cache[id]; ok {
		return im
	}

	panic("Texture not found in cache. Cannot obtain ebiten Image")
}

func (c *textureCache) SetTexture(id imgui.TextureID, img *ebiten.Image) {
	c.cache[id] = img
}

func (c *textureCache) RemoveTexture(id imgui.TextureID) {
	delete(c.cache, id)
}

func (c *textureCache) ResetFontAtlasCache(filter ebiten.Filter) {
	c.fontAtlasImage = nil
	c.dfilter = filter
}

func NewCache() TextureCache {
	return &textureCache{
		startIndex:     2,
		fontAtlasID:    imgui.TextureID{Data: uintptr(1)},
		cache:          make(map[imgui.TextureID]*ebiten.Image),
		fontAtlasImage: nil,
	}
}

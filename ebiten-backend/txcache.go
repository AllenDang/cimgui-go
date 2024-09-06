package ebitenbackend

import (
	"unsafe"

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
}

type textureCache struct {
	fontAtlasID    imgui.TextureID
	fontAtlasImage *ebiten.Image
	cache          map[imgui.TextureID]*ebiten.Image
	dfilter        ebiten.Filter
}

var _ TextureCache = (*textureCache)(nil)

func (c *textureCache) getFontAtlas() *ebiten.Image {
	if c.fontAtlasImage == nil {
		pixels, width, height, outBytesPerPixel := imgui.CurrentIO().Fonts().GetTextureDataAsRGBA32() // call this to force imgui to build the font atlas cache
		c.fontAtlasImage = getTexture(pixels, width, height, outBytesPerPixel)
	}
	return c.fontAtlasImage
}

func (c *textureCache) FontAtlasTextureID() imgui.TextureID {
	return c.fontAtlasID
}

func (c *textureCache) SetFontAtlasTextureID(id imgui.TextureID) {
	c.fontAtlasID = id
	// c.fontAtlasImage = nil
}

func (c *textureCache) GetTexture(id imgui.TextureID) *ebiten.Image {
	if id != c.fontAtlasID {
		if im, ok := c.cache[id]; ok {
			return im
		}
	}
	return c.getFontAtlas()
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
	texID := imgui.TextureID{
		Data: uintptr(unsafe.Pointer(&id1)),
	}

	return &textureCache{
		fontAtlasID:    texID,
		cache:          make(map[imgui.TextureID]*ebiten.Image),
		fontAtlasImage: nil,
	}
}

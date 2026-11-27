package ebitenbackend

import (
	"image"
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
	next      int // above every ID set, so a removed texture's ID isn't reused
	cache     map[imgui.TextureID]*ebiten.Image
	cacheGame map[imgui.TextureID]ebiten.Game
	dfilter   ebiten.Filter
	pix       []byte // premultiplied pixels on their way to a texture
}

var _ TextureCache = (*textureCache)(nil)

func (c *textureCache) NextId() int {
	return c.next
}

// UpdateTexture does what ImGui asks of a texture and tells it when done, as
// Dear ImGui 1.92 wants: it makes the texture, writes the pixels ImGui changed
// into it, or frees it.
func (c *textureCache) UpdateTexture(tex imgui.TextureData) {
	switch tex.Status() {
	case imgui.TextureStatusOK, imgui.TextureStatusDestroyed:
		// noop.
	case imgui.TextureStatusWantCreate:
		w, h := int(tex.Width()), int(tex.Height())
		texImage := ebiten.NewImage(w, h)
		c.writePixels(texImage, tex, image.Rect(0, 0, w, h))
		newID := imgui.TextureID(c.NextId())
		c.SetTexture(newID, texImage)
		tex.SetTexID(newID)
		tex.SetStatus(imgui.TextureStatusOK)
	case imgui.TextureStatusWantUpdates:
		r := tex.UpdateRect()
		x, y := int(r.X()), int(r.Y())
		c.writePixels(c.GetTexture(tex.TexID()), tex, image.Rect(x, y, x+int(r.W()), y+int(r.H())))
		tex.SetStatus(imgui.TextureStatusOK)
	case imgui.TextureStatusWantDestroy:
		if img, ok := c.cache[tex.TexID()]; ok {
			img.Deallocate()
		}
		c.RemoveTexture(tex.TexID())
		tex.SetTexID(0) // ImTextureID_Invalid
		tex.SetStatus(imgui.TextureStatusDestroyed)
	default:
		log.Panicf("Unknown texture status: %v", tex.Status())
	}
}

// writePixels writes rect of tex's pixels into img: ImGui keeps them in
// straight alpha, and Ebitengine wants them premultiplied.
func (c *textureCache) writePixels(img *ebiten.Image, tex imgui.TextureData, rect image.Rectangle) {
	if rect.Empty() {
		return
	}
	stride := int(tex.Width()) * 4
	src := unsafe.Slice((*byte)(unsafe.Pointer(tex.Pixels())), stride*int(tex.Height()))
	n := rect.Dx() * rect.Dy() * 4
	if cap(c.pix) < n {
		c.pix = make([]byte, n)
	}
	dst := c.pix[:n]
	row := rect.Dx() * 4
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		at := y*stride + rect.Min.X*4
		premultiply(dst[(y-rect.Min.Y)*row:], src[at:at+row])
	}
	img.SubImage(rect).(*ebiten.Image).WritePixels(dst)
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
	c.next = max(c.next, int(id)+1)
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
		next:      2,
		cache:     make(map[imgui.TextureID]*ebiten.Image),
		cacheGame: make(map[imgui.TextureID]ebiten.Game),
	}
}

package raylibbackend

import (
	"image"
	"image/color"
	"log"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TextureCache interface defines methods for managing textures in Raylib
type TextureCache interface {
	UpdateTexture(tex imgui.TextureData)
	GetTexture(id imgui.TextureID) rl.Texture2D
	HasTexture(id imgui.TextureID) bool
	SetTexture(id imgui.TextureID, img rl.Texture2D)
	RemoveTexture(id imgui.TextureID)
	NextId() int
}

// textureCache implements TextureCache interface for managing textures in Raylib
type textureCache struct {
	startIndex int
	cache      map[imgui.TextureID]rl.Texture2D
	nextID     int
}

var _ TextureCache = (*textureCache)(nil)

// NewTextureCache creates a new texture cache instance
func NewTextureCache() TextureCache {
	return &textureCache{
		startIndex: 2,
		cache:      make(map[imgui.TextureID]rl.Texture2D),
		nextID:     0,
	}
}

// NextId returns the next available texture ID for the cache
func (c *textureCache) NextId() int {
	id := c.nextID + c.startIndex
	c.nextID++

	return id
}

// UpdateTexture updates the texture cache based on the provided ImGui TextureData
func (c *textureCache) UpdateTexture(tex imgui.TextureData) {
	switch tex.Status() {
	case imgui.TextureStatusOK:
		// noop.

	case imgui.TextureStatusWantCreate:
		texImage := getTextureFromTextureData(tex)
		newID := imgui.TextureID(c.NextId())

		// If the ID already exists in the cache, unload it first
		if old, exists := c.cache[newID]; exists {
			rl.UnloadTexture(old)
		}

		c.SetTexture(newID, texImage)
		tex.SetTexID(newID)
		tex.SetStatus(imgui.TextureStatusOK)

	case imgui.TextureStatusWantDestroy:
		id := tex.TexID()
		c.RemoveTexture(id)

	case imgui.TextureStatusWantUpdates:
		id := tex.TexID()

		if existing, ok := c.cache[id]; ok {
			rgba := rgbaColorsFromTextureData(tex)
			rl.UpdateTexture(existing, rgba)

			// Re-ensure parameters after update
			rl.SetTextureFilter(existing, rl.FilterBilinear)
			rl.SetTextureWrap(existing, rl.WrapClamp)
		} else {
			texImage := getTextureFromTextureData(tex)
			c.SetTexture(id, texImage)
		}

		tex.SetStatus(imgui.TextureStatusOK)

	default:
		log.Panicf("Unknown texture status: %v", tex.Status())
	}
}

// GetTexture retrieves a texture from the cache by ID
func (c *textureCache) GetTexture(id imgui.TextureID) rl.Texture2D {
	if im, ok := c.cache[id]; ok {
		return im
	}

	panic("Texture not found in cache.")
}

// HasTexture returns bool if a texture exists in the texture cache or not
func (c *textureCache) HasTexture(id imgui.TextureID) bool {
	_, ok := c.cache[id]
	return ok
}

// SetTexture sets a texture in the cache, unloading any previous texture if present
func (c *textureCache) SetTexture(id imgui.TextureID, img rl.Texture2D) {
	if old, ok := c.cache[id]; ok {
		rl.UnloadTexture(old)
	}

	c.cache[id] = img
}

// RemoveTexture removes a texture from the cache and unloads it
func (c *textureCache) RemoveTexture(id imgui.TextureID) {
	if old, ok := c.cache[id]; ok {
		rl.UnloadTexture(old)
	}

	delete(c.cache, id)
}

// getTextureFromTextureData converts ImGui TextureData into RGBA pixels if needed and uploads it
func getTextureFromTextureData(td imgui.TextureData) rl.Texture2D {
	w, h := int(td.Width()), int(td.Height())
	rgba := rgbaBytesFromTextureData(td)
	img := imageRGBA(w, h, rgba)

	rlImg := rl.NewImageFromImage(img)
	tex := rl.LoadTextureFromImage(rlImg)
	rl.UnloadImage(rlImg)

	// Improve sampling quality for fonts and UI
	rl.SetTextureFilter(tex, rl.FilterBilinear)
	rl.SetTextureWrap(tex, rl.WrapClamp)

	return tex
}

// imageRGBA creates an image.RGBA wrapper around provided pixel data
func imageRGBA(w, h int, pix []byte) *image.RGBA {
	return &image.RGBA{
		Pix:    pix,
		Stride: 4 * w,
		Rect:   image.Rect(0, 0, w, h),
	}
}

// rgbaColorsFromTextureData converts TextureData to []color.RGBA for rl.UpdateTexture
func rgbaColorsFromTextureData(td imgui.TextureData) []color.RGBA {
	bytes := rgbaBytesFromTextureData(td)
	n := len(bytes) / 4
	colors := make([]color.RGBA, n)

	for i := 0; i < n; i++ {
		colors[i] = color.RGBA{
			R: bytes[i*4+0],
			G: bytes[i*4+1],
			B: bytes[i*4+2],
			A: bytes[i*4+3],
		}
	}

	return colors
}

// rgbaBytesFromTextureData converts TextureData (Alpha8 or RGBA32) to a packed []byte RGBA slice
func rgbaBytesFromTextureData(td imgui.TextureData) []uint8 {
	w := int(td.Width())
	h := int(td.Height())
	pitch := int(td.Pitch())
	format := td.Format()
	srcPtr := unsafe.Pointer(td.Pixels())

	if format == imgui.TextureFormatRGBA32 && pitch == w*4 {
		return (*[1 << 28]uint8)(srcPtr)[: w*h*4 : w*h*4]
	}

	dst := make([]uint8, w*h*4)

	switch format {
	case imgui.TextureFormatRGBA32:
		srcBpp := 4
		for y := 0; y < h; y++ {
			srcRow := (*[1 << 28]uint8)(unsafe.Pointer(uintptr(srcPtr) + uintptr(y*pitch)))[:pitch:pitch]
			copy(dst[y*w*4:(y+1)*w*4], srcRow[:w*srcBpp])
		}

	case imgui.TextureFormatAlpha8:
		for y := 0; y < h; y++ {
			srcRow := (*[1 << 28]uint8)(unsafe.Pointer(uintptr(srcPtr) + uintptr(y*pitch)))[:pitch:pitch]
			for x := 0; x < w; x++ {
				a := srcRow[x]
				o := (y*w + x) * 4
				dst[o+0] = 0xFF
				dst[o+1] = 0xFF
				dst[o+2] = 0xFF
				dst[o+3] = a
			}
		}

	default:
		// Assume RGBA32 layout
		for y := 0; y < h; y++ {
			srcRow := (*[1 << 28]uint8)(unsafe.Pointer(uintptr(srcPtr) + uintptr(y*pitch)))[:pitch:pitch]
			copy(dst[y*w*4:(y+1)*w*4], srcRow[:w*4])
		}
	}

	return dst
}

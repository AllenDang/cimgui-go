package backend

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/AllenDang/cimgui-go/imgui"
)

// Texture implements a simple texture loader. It wraps backend's methods to allow creating textures easily.
// IMPORTANT: as the texture is mainly handled by C OpenGL, it is not covered by Garbae Collector (GC).
//
//	Remember to call (*Texture).Release when you no longer need it.
type Texture struct {
	ID     imgui.TextureID
	Width  int
	Height int
}

func NewTextureFromRgba(rgba *image.RGBA) *Texture {
	texID := textureManager.CreateTextureRgba(rgba, rgba.Bounds().Dx(), rgba.Bounds().Dy())

	texture := Texture{
		ID:     texID,
		Width:  rgba.Bounds().Dx(),
		Height: rgba.Bounds().Dy(),
	}

	// I leav it for documentation here:
	// GC runs in a separated thread so this may not work correctly (will crash opengl)
	// runtime.SetFinalizer(&texture, (*Texture).release)

	return &texture
}

// Release tells OpenGL that this texture is no longer needed.
// ATTENTION: This will not be automatically handled by GC so remember to do it manually if you have many textures!
func (t *Texture) Release() {
	textureManager.DeleteTexture(t.ID)
}

// ImageToRgba converts image.Image to *image.RGBA.
func ImageToRgba(img image.Image) *image.RGBA {
	switch trueImg := img.(type) {
	case *image.RGBA:
		return trueImg
	default:
		rgba := image.NewRGBA(trueImg.Bounds())
		draw.Draw(rgba, trueImg.Bounds(), trueImg, image.Pt(0, 0), draw.Src)
		return rgba
	}
}

// LoadImage loads image from file and returns *image.RGBA.
func LoadImage(imgPath string) (*image.RGBA, error) {
	imgFile, err := os.Open(filepath.Clean(imgPath))
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error opening image file %s: %w", imgPath, err)
	}

	defer func() {
		// nolint:govet // we want to reuse this err variable here
		if err := imgFile.Close(); err != nil {
			panic(fmt.Sprintf("error closing image file: %s", imgPath))
		}
	}()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, fmt.Errorf("LoadImage: error decoding png image: %w", err)
	}

	return ImageToRgba(img), nil
}

package raylibbackend

import (
	"math"

	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ImageRTOptions configures how a RenderTexture2D is rendered and interacted with
type ImageRTOptions struct {
	// Align the ImGui cursor to the framebuffer pixel grid before drawing
	AlignCursor bool
	// Flip the image horizontally when drawing
	FlipH bool
	// Flip the image vertically when drawing
	FlipV bool
	// Require app focus for mapping
	RequireAppFocus bool
	// Require the current ImGui window hovered for mapping
	RequireWindowHovered bool
}

// ImageRTResult reports the layout and mouse interaction on the texture
type ImageRTResult struct {
	ImagePos  imgui.Vec2
	ImageSize imgui.Vec2
	Hovered   bool
	TexX      float32
	TexY      float32
}

// defaultOpts configures default options
func defaultOpts(opts *ImageRTOptions) ImageRTOptions {
	if opts == nil {
		return ImageRTOptions{
			AlignCursor:          true,
			FlipV:                true,
			RequireAppFocus:      true,
			RequireWindowHovered: true,
		}
	}

	out := *opts

	if !opts.AlignCursor {
		out.AlignCursor = true
	}

	if !opts.FlipV {
		out.FlipV = true
	}

	if !opts.RequireAppFocus {
		out.RequireAppFocus = true
	}

	if !opts.RequireWindowHovered {
		out.RequireWindowHovered = true
	}

	return out
}

// PixelAlignedImageSize computes an ImGui size with DisplayFramebufferScale applied in pixels
func PixelAlignedImageSize(texW, texH int) imgui.Vec2 {
	io := imgui.CurrentIO()
	ds := io.DisplayFramebufferScale()

	if ds.X == 0 {
		ds.X = 1
	}

	if ds.Y == 0 {
		ds.Y = 1
	}

	sxIm := float32(math.Round(float64(float32(texW) / ds.X)))
	syIm := float32(math.Round(float64(float32(texH) / ds.Y)))

	return imgui.NewVec2(sxIm, syIm)
}

// AlignCursorToFramebufferPixel snaps the current ImGui cursor so items align to framebuffer pixels
func AlignCursorToFramebufferPixel() imgui.Vec2 {
	io := imgui.CurrentIO()
	ds := io.DisplayFramebufferScale()

	if ds.X == 0 {
		ds.X = 1
	}

	if ds.Y == 0 {
		ds.Y = 1
	}

	cur := imgui.CursorScreenPos()
	cur.X = float32(math.Round(float64(cur.X*ds.X)) / float64(ds.X))
	cur.Y = float32(math.Round(float64(cur.Y*ds.Y)) / float64(ds.Y))

	imgui.SetCursorScreenPos(cur)

	return cur
}

// MapMouseToTexture converts current mouse position to texture pixel coords for the given image rect
func MapMouseToTexture(imgPos, imgSize imgui.Vec2, texW, texH int, flipV bool) (tx, ty float32, hovered bool) {
	io := imgui.CurrentIO()
	m := io.MousePos()

	if m.X < imgPos.X ||
		m.X >= imgPos.X+imgSize.X ||
		m.Y < imgPos.Y ||
		m.Y >= imgPos.Y+imgSize.Y {
		return 0, 0, false
	}

	hovered = true
	lx := m.X - imgPos.X
	ly := m.Y - imgPos.Y
	sx := float32(texW) / imgSize.X
	sy := float32(texH) / imgSize.Y
	tx = lx * sx

	if flipV {
		ty = ly * sy
	} else {
		ty = float32(texH-1) - ly*sy
	}

	// Clamp
	if tx < 0 {
		tx = 0
	}

	if ty < 0 {
		ty = 0
	}

	if tx > float32(texW-1) {
		tx = float32(texW - 1)
	}

	if ty > float32(texH-1) {
		ty = float32(texH - 1)
	}

	return
}

// textureRefForTexture returns an ImGui texture reference for a raylib Texture2D
func (r *RaylibBackend) textureRefForTexture(tex rl.Texture2D) imgui.TextureRef {
	return r.CreateTextureFromTexture2D(tex)
}

// PrepareImageRenderTexture computes size/pos and optional mouse mapping for a RenderTexture2D
func (r *RaylibBackend) PrepareImageRenderTexture(rt rl.RenderTexture2D, opts *ImageRTOptions) ImageRTResult {
	o := defaultOpts(opts)

	size := PixelAlignedImageSize(int(rt.Texture.Width), int(rt.Texture.Height))

	var pos imgui.Vec2
	if o.AlignCursor {
		pos = AlignCursorToFramebufferPixel()
	} else {
		pos = imgui.CursorScreenPos()
	}

	canMap := true
	if o.RequireAppFocus {
		canMap = rl.IsWindowFocused()
	}

	if canMap && o.RequireWindowHovered {
		canMap = imgui.IsWindowHovered()
	}

	res := ImageRTResult{ImagePos: pos, ImageSize: size}
	if canMap {
		tx, ty, hov := MapMouseToTexture(pos, size, int(rt.Texture.Width), int(rt.Texture.Height), o.FlipV)
		if hov {
			res.Hovered = true
			res.TexX = tx
			res.TexY = ty
		}
	}

	return res
}

// DrawImageRenderTexture draws the RenderTexture2D as an ImGui image
func (r *RaylibBackend) DrawImageRenderTexture(rt rl.RenderTexture2D, size imgui.Vec2, flipV bool, flipH bool) {
	ref := r.textureRefForTexture(rt.Texture)
	uv0 := imgui.NewVec2(0, 0)
	uv1 := imgui.NewVec2(1, 1)

	if flipV {
		uv0.Y = 1
		uv1.Y = 0
	}

	if flipH {
		uv0.X = 1
		uv1.X = 0
	}

	imgui.ImageV(ref, size, uv0, uv1)
}

// ImageRenderTexture prepares then draws the image, and returns the result
func (r *RaylibBackend) ImageRenderTexture(rt rl.RenderTexture2D, opts *ImageRTOptions) ImageRTResult {
	o := defaultOpts(opts)
	res := r.PrepareImageRenderTexture(rt, &o)
	r.DrawImageRenderTexture(rt, res.ImageSize, o.FlipV, o.FlipH)

	return res
}

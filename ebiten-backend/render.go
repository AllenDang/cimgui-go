package ebitenbackend

import (
	"fmt"
	"image"
	"unsafe"

	imgui "github.com/AllenDang/cimgui-go"
	"github.com/AllenDang/cimgui-go/ebiten-backend/internal/native"
	"github.com/hajimehoshi/ebiten/v2"
)

type cImDrawVertx32 struct {
	Pos struct{ X, Y float32 }
	UV  struct{ X, Y float32 }
	Col uint32
}

type cVec2x64 struct {
	X float64
	Y float64
}

type cImDrawVertx64 struct {
	Pos cVec2x64
	UV  cVec2x64
	Col uint32
}

func premultiplyPixels(pixels unsafe.Pointer, width, height int) []uint8 {
	n := width * height
	srcPix := (*[1 << 28]uint8)(pixels)[: n*4 : n*4]
	pix := make([]uint8, n*4)
	for i := 0; i < n; i++ {
		alpha := uint16(srcPix[4*i+3])
		pix[4*i] = uint8((uint16(srcPix[4*i])*alpha + 127) / 255)
		pix[4*i+1] = uint8((uint16(srcPix[4*i+1])*alpha + 127) / 255)
		pix[4*i+2] = uint8((uint16(srcPix[4*i+2])*alpha + 127) / 255)
		pix[4*i+3] = uint8(alpha)
	}

	return pix
}

func getVertices(vbuf unsafe.Pointer, vblen, vsize, offpos, offuv,
	offcol int,
) []ebiten.Vertex {
	if native.SzFloat() == 4 {
		return getVerticesx32(vbuf, vblen, vsize, offpos, offuv, offcol)
	}
	if native.SzFloat() == 8 {
		return getVerticesx64(vbuf, vblen, vsize, offpos, offuv, offcol)
	}
	panic("invalid char size")
}

func getVerticesx32(vbuf unsafe.Pointer, vblen, vsize, offpos, offuv,
	offcol int,
) []ebiten.Vertex {
	n := vblen / vsize
	vertices := make([]ebiten.Vertex, 0, vblen/vsize)
	if offpos != 0 || offuv != 8 || offcol != 16 {
		panic("TODO: invalid vertex layout")
	}

	rawverts := (*[1 << 28]cImDrawVertx32)(vbuf)[:n:n]
	for i := 0; i < n; i++ {
		vertices = append(vertices, ebiten.Vertex{
			SrcX:   rawverts[i].UV.X,
			SrcY:   rawverts[i].UV.Y,
			DstX:   rawverts[i].Pos.X,
			DstY:   rawverts[i].Pos.Y,
			ColorR: float32(rawverts[i].Col&0xFF) / 255,
			ColorG: float32(rawverts[i].Col>>8&0xFF) / 255,
			ColorB: float32(rawverts[i].Col>>16&0xFF) / 255,
			ColorA: float32(rawverts[i].Col>>24&0xFF) / 255,
		})
	}
	return vertices
}

func getVerticesx64(vbuf unsafe.Pointer, vblen, vsize, offpos, offuv,
	offcol int,
) []ebiten.Vertex {
	n := vblen / vsize
	vertices := make([]ebiten.Vertex, 0, vblen/vsize)
	if offpos != 0 || offuv != 8 || offcol != 16 {
		panic("TODO: invalid vertex layout (64)")
	}
	rawverts := (*[1 << 28]cImDrawVertx64)(vbuf)[:n:n]
	for i := 0; i < n; i++ {
		vertices = append(vertices, ebiten.Vertex{
			SrcX:   float32(rawverts[i].UV.X),
			SrcY:   float32(rawverts[i].UV.Y),
			DstX:   float32(rawverts[i].Pos.X),
			DstY:   float32(rawverts[i].Pos.Y),
			ColorR: float32(rawverts[i].Col&0xFF) / 255,
			ColorG: float32(rawverts[i].Col>>8&0xFF) / 255,
			ColorB: float32(rawverts[i].Col>>16&0xFF) / 255,
			ColorA: float32(rawverts[i].Col>>24&0xFF) / 255,
		})
	}
	return vertices
}

func lerp(a, b int, t float32) float32 {
	return float32(a)*(1-t) + float32(b)*t
}

func vcopy(v []ebiten.Vertex) []ebiten.Vertex {
	cl := make([]ebiten.Vertex, len(v))
	copy(cl, v)
	return cl
}

func vmultiply(v, vbuf []ebiten.Vertex, bmin, bmax image.Point) {
	for i := range vbuf {
		vbuf[i].SrcX = lerp(bmin.X, bmax.X, v[i].SrcX)
		vbuf[i].SrcY = lerp(bmin.Y, bmax.Y, v[i].SrcY)
	}
}

func getTexture(pixels unsafe.Pointer, width int32, height int32) *ebiten.Image {
	n := int(width * height)
	srcPix := (*[1 << 28]uint8)(pixels)[: n*4 : n*4]
	pix := make([]uint8, n*4)
	// Note: Ebiten expects colors in premultiplied-alpha form.
	// However, the imgui library exports pixmaps in straight-alpha form.
	// Also, not doing this modification in-place,
	// as srcPix points right into an imgui-owned data structure.
	for i := 0; i < n; i++ {
		alpha := uint16(srcPix[4*i+3])
		pix[4*i] = uint8((uint16(srcPix[4*i])*alpha + 127) / 255)
		pix[4*i+1] = uint8((uint16(srcPix[4*i+1])*alpha + 127) / 255)
		pix[4*i+2] = uint8((uint16(srcPix[4*i+2])*alpha + 127) / 255)
		pix[4*i+3] = uint8(alpha)
	}
	img := ebiten.NewImage(int(width), int(height))
	img.WritePixels(pix)
	return img
}

func getIndices(ibuf unsafe.Pointer, iblen, isize int) []uint16 {
	n := iblen / isize
	switch isize {
	case 2:
		// direct conversion (without a data copy)
		// TODO: document the size limit (?) this fits 268435456 bytes
		// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
		return (*[1 << 28]uint16)(ibuf)[:n:n]
	case 4:
		slc := make([]uint16, n)
		for i := 0; i < n; i++ {
			slc[i] = uint16(*(*uint32)(unsafe.Pointer(uintptr(ibuf) + uintptr(i*isize))))
		}
		return slc
	case 8:
		slc := make([]uint16, n)
		for i := 0; i < n; i++ {
			slc[i] = uint16(*(*uint64)(unsafe.Pointer(uintptr(ibuf) + uintptr(i*isize))))
		}
		return slc
	default:
		panic(fmt.Sprint("byte size", isize, "not supported"))
	}
}

// Render the ImGui drawData into the target *ebiten.Image
func Render(target *ebiten.Image, drawData *imgui.DrawData, txcache TextureCache,
	dfilter ebiten.Filter,
) {
	render(target, nil, drawData, txcache, dfilter)
}

// RenderMasked renders the ImGui drawData into the target *ebiten.Image
func RenderMasked(target *ebiten.Image, mask *ebiten.Image, drawData *imgui.DrawData,
	txcache TextureCache, dfilter ebiten.Filter,
) {
	render(target, mask, drawData, txcache, dfilter)
}

func render(target *ebiten.Image, mask *ebiten.Image, drawData *imgui.DrawData,
	txcache TextureCache, dfilter ebiten.Filter,
) {
	targetSize := target.Bounds().Size()
	if !drawData.Valid() {
		return
	}

	vertexSize,
		vertexOffsetPos,
		vertexOffsetUv,
		vertexOffsetCol := imgui.VertexBufferLayout()

	indexSize := imgui.IndexBufferLayout()

	opt := &ebiten.DrawTrianglesOptions{
		Filter: dfilter,
	}
	var opt2 *ebiten.DrawImageOptions
	if mask != nil {
		opt2 = &ebiten.DrawImageOptions{
			Blend: ebiten.BlendSourceOver,
		}
	}

	for _, clist := range drawData.CommandLists() {
		vertexBuffer, vertexLen := clist.GetVertexBuffer()
		indexBuffer, indexLen := clist.GetIndexBuffer()
		vertices := getVertices(vertexBuffer, vertexLen, vertexSize, vertexOffsetPos,
			vertexOffsetUv, vertexOffsetCol)
		vbuf := vcopy(vertices)
		indices := getIndices(indexBuffer, indexLen, indexSize)
		for _, cmd := range clist.Commands() {
			idxOffset := int(cmd.IdxOffset())

			ecount := int(cmd.ElemCount())
			if cmd.HasUserCallback() {
				cmd.CallUserCallback(clist)
			} else {
				clipRect := cmd.ClipRect()
				texid := cmd.TextureId()
				tx := txcache.GetTexture(texid)
				vmultiply(vertices, vbuf, tx.Bounds().Min, tx.Bounds().Max)
				if mask == nil ||
					(clipRect.X == 0 &&
						clipRect.Y == 0 &&
						clipRect.Z == float32(targetSize.X) &&
						clipRect.W == float32(targetSize.Y)) {
					target.DrawTriangles(vbuf, indices[idxOffset:idxOffset+ecount], tx,
						opt)
				} else {
					mask.Clear()
					if opt2 != nil {
						opt2.GeoM.Reset()
						opt2.GeoM.Translate(float64(clipRect.X), float64(clipRect.Y))
						mask.DrawTriangles(vbuf, indices[idxOffset:idxOffset+ecount], tx, opt)
						target.DrawImage(mask.SubImage(image.Rectangle{
							Min: image.Pt(int(clipRect.X), int(clipRect.Y)),
							Max: image.Pt(int(clipRect.Z), int(clipRect.W)),
						}).(*ebiten.Image), opt2)
					}
				}
			}
		}
	}
}

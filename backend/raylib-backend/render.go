package raylibbackend

import (
	"math"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// VertexLayout defines the ImGui vertex buffer
type VertexLayout struct {
	Stride int
	PosOff int
	UVOff  int
	ColOff int
}

// RenderTransform holds the current coordinate and scaling info
type RenderTransform struct {
	DisplayPos imgui.Vec2
	Scale      imgui.Vec2
	UVScale    imgui.Vec2
	FlipV      bool
}

// getMaterialFor retrieves or creates a raylib material for a given texture
func (r *RaylibBackend) getMaterialFor(tex rl.Texture2D) rl.Material {
	if m, ok := r.matCache[tex.ID]; ok {
		return m
	}

	m := rl.LoadMaterialDefault()
	rl.SetMaterialTexture(&m, rl.MapDiffuse, tex)
	r.matCache[tex.ID] = m

	return m
}

// drawMeshTextured creates a temporary Raylib mesh for textured drawing
func (r *RaylibBackend) drawMeshTextured(tex rl.Texture2D, vbuf unsafe.Pointer, indexList []uint16,
	baseIndexOffset, vbase, count int, layout VertexLayout, xfrm RenderTransform) {

	if count <= 0 || tex.Width <= 0 || tex.Height <= 0 {
		return
	}

	if cap(r.meshVerts) < count*3 {
		r.meshVerts = make([]float32, count*3)
	} else {
		r.meshVerts = r.meshVerts[:count*3]
	}
	if cap(r.meshUVs) < count*2 {
		r.meshUVs = make([]float32, count*2)
	} else {
		r.meshUVs = r.meshUVs[:count*2]
	}
	if cap(r.meshCols) < count*4 {
		r.meshCols = make([]uint8, count*4)
	} else {
		r.meshCols = r.meshCols[:count*4]
	}
	if cap(r.meshIdxs) < count {
		r.meshIdxs = make([]uint16, count)
	} else {
		r.meshIdxs = r.meshIdxs[:count]
	}

	verts := r.meshVerts
	uvs := r.meshUVs
	cols := r.meshCols
	idxs := r.meshIdxs

	vi := 0
	ui := 0
	ci := 0

	for i := 0; i < count; i++ {
		vidx := vbase + int(indexList[baseIndexOffset+i])
		base := uintptr(vbuf) + uintptr(vidx*layout.Stride)

		// Transform Position
		vx, vy := transformPosition(base, layout, xfrm)
		verts[vi] = vx
		verts[vi+1] = vy
		verts[vi+2] = 0
		vi += 3

		// Transform UVs
		u := readF32(base+uintptr(layout.UVOff+0)) * xfrm.UVScale.X
		v := readF32(base+uintptr(layout.UVOff+4)) * xfrm.UVScale.Y
		if xfrm.FlipV {
			v = 1 - v
		}
		uvs[ui] = u
		uvs[ui+1] = v
		ui += 2

		// Extract Color
		col := readU32(base + uintptr(layout.ColOff))
		cols[ci] = uint8(col & 0xFF)
		cols[ci+1] = uint8((col >> 8) & 0xFF)
		cols[ci+2] = uint8((col >> 16) & 0xFF)
		cols[ci+3] = uint8((col >> 24) & 0xFF)
		ci += 4

		idxs[i] = uint16(i)
	}

	var mesh rl.Mesh
	mesh.VertexCount = int32(count)
	mesh.TriangleCount = int32(count / 3)
	mesh.Vertices = &verts[0]
	mesh.Texcoords = &uvs[0]
	mesh.Colors = &cols[0]
	mesh.Indices = &idxs[0]

	rl.UploadMesh(&mesh, true)
	rl.DrawMesh(mesh, r.getMaterialFor(tex), rl.MatrixIdentity())
	rl.UnloadMesh(&mesh)
}

// drawRLGL uses immediate mode for untextured colored primitives.
func (r *RaylibBackend) drawRLGL(vbuf unsafe.Pointer, indexList []uint16,
	baseIndexOffset, vbase, count int, layout VertexLayout, xfrm RenderTransform) {

	// Default white texture
	rl.SetTexture(0)

	rl.Begin(rl.Triangles)

	for i := 0; i < count; i++ {
		vidx := vbase + int(indexList[baseIndexOffset+i])
		r.emitVertex(vbuf, vidx, layout, xfrm)
	}

	rl.End()
}

func (r *RaylibBackend) emitVertex(vbuf unsafe.Pointer, index int, layout VertexLayout, xfrm RenderTransform) {
	base := uintptr(vbuf) + uintptr(index*layout.Stride)

	x, y := transformPosition(base, layout, xfrm)

	u := readF32(base+uintptr(layout.UVOff+0)) * xfrm.UVScale.X
	v := readF32(base+uintptr(layout.UVOff+4)) * xfrm.UVScale.Y

	if xfrm.FlipV {
		v = 1 - v
	}

	hexCol := readU32(base + uintptr(layout.ColOff))
	c := hexToColor(hexCol)

	rl.Color4ub(c.R, c.G, c.B, c.A)
	rl.TexCoord2f(u, v)
	rl.Vertex2f(x, y)
}

// renderRL renders the ImGui draw data using Raylib
func (r *RaylibBackend) renderRL(drawData *imgui.DrawData) {
	if !drawData.Valid() {
		return
	}

	// Update Textures
	for _, tex := range drawData.Textures().Slice() {
		r.cache.UpdateTexture(tex)
	}

	// Setup layout and transform
	var layout VertexLayout
	if r.cachedLayoutValid {
		layout = *r.cachedLayout
	} else {
		vSize, vPos, vUV, vCol := imgui.VertexBufferLayout()
		layout = VertexLayout{vSize, vPos, vUV, vCol}
		r.cachedLayout = &layout
		r.cachedLayoutValid = true
	}

	dp := drawData.DisplayPos()
	ds := drawData.FramebufferScale()

	if ds.X == 0 {
		ds.X = 1
	}

	if ds.Y == 0 {
		ds.Y = 1
	}

	xfrm := RenderTransform{
		DisplayPos: dp,
		Scale:      ds,
		UVScale:    imgui.Vec2{X: 1, Y: 1},
		FlipV:      false,
	}

	// Render State Setup
	rl.DisableBackfaceCulling()
	rl.DisableDepthTest()
	rl.BeginBlendMode(rl.BlendAlpha)

	// Draw Command Processing
	for _, clist := range drawData.CommandLists() {
		vbuf, _ := clist.GetVertexBuffer()
		ibuf, iblen := clist.GetIndexBuffer()
		indices := getIndices(ibuf, iblen, imgui.IndexBufferLayout())

		for _, cmd := range clist.Commands() {
			if cmd.HasUserCallback() {
				cmd.CallUserCallback(clist)
				continue
			}

			// Handle Clipping
			sx, sy, sw, sh := calculateScissor(cmd.ClipRect(), dp, ds)
			if sw <= 0 || sh <= 0 {
				continue
			}

			if r.disabledScissorFrames == 0 {
				rl.BeginScissorMode(sx, sy, sw, sh)
			}

			// Choose Render Path
			texID := cmd.TexID()
			idxOffset := int(cmd.IdxOffset())
			vtxOffset := int(cmd.VtxOffset())
			elemCount := int(cmd.ElemCount())

			if texID == 0 {
				r.drawRLGL(vbuf, indices, idxOffset, vtxOffset, elemCount, layout, xfrm)
			} else if r.cache.HasTexture(texID) {
				tex := r.cache.GetTexture(texID)
				r.drawMeshTextured(tex, vbuf, indices, idxOffset, vtxOffset, elemCount, layout, xfrm)
			}

			if r.disabledScissorFrames == 0 {
				rl.EndScissorMode()
			}
		}
	}

	// Cleanup
	rl.EndBlendMode()
	rl.EnableDepthTest()
	rl.EnableBackfaceCulling()

	if r.disabledScissorFrames > 0 {
		r.disabledScissorFrames--
	}
}

// readF32 float32 pointer reader
func readF32(ptr uintptr) float32 { return *(*float32)(unsafe.Pointer(ptr)) }

// readU32 uint32 pointer reader
func readU32(ptr uintptr) uint32 { return *(*uint32)(unsafe.Pointer(ptr)) }

// getIndices converts index buffer to uint16 slice
func getIndices(ibuf unsafe.Pointer, iblen, isize int) []uint16 {
	n := iblen / isize
	switch isize {
	case 2:
		return (*[1 << 28]uint16)(ibuf)[:n:n]
	case 4:
		src := (*[1 << 27]uint32)(ibuf)[:n:n]
		slc := make([]uint16, n)
		for i := 0; i < n; i++ {
			slc[i] = uint16(src[i])
		}
		return slc
	case 8:
		src := (*[1 << 26]uint64)(ibuf)[:n:n]
		slc := make([]uint16, n)
		for i := 0; i < n; i++ {
			slc[i] = uint16(src[i])
		}
		return slc
	default:
		panic("unsupported index element size")
	}
}

// transformPosition transforms vertex position coordinates from ImGui space to render space
func transformPosition(base uintptr, layout VertexLayout, xfrm RenderTransform) (float32, float32) {
	vx := (readF32(base+uintptr(layout.PosOff+0)) - xfrm.DisplayPos.X) * xfrm.Scale.X
	vy := (readF32(base+uintptr(layout.PosOff+4)) - xfrm.DisplayPos.Y) * xfrm.Scale.Y
	return vx, vy
}

// hexToColor converts ImGui's packed ABGR/RGBA uint32 to a raylib.Color
func hexToColor(col uint32) rl.Color {
	return rl.NewColor(
		uint8(col&0xFF),
		uint8((col>>8)&0xFF),
		uint8((col>>16)&0xFF),
		uint8((col>>24)&0xFF),
	)
}

// calculateScissor transforms an ImGui clip rect into Raylib screen coordinates
func calculateScissor(cr imgui.Vec4, dp, ds imgui.Vec2) (int32, int32, int32, int32) {
	clipX := (cr.X - dp.X) * ds.X
	clipY := (cr.Y - dp.Y) * ds.Y
	clipR := (cr.Z - dp.X) * ds.X
	clipB := (cr.W - dp.Y) * ds.Y

	sx := int32(math.Floor(float64(clipX)))
	sy := int32(math.Floor(float64(clipY)))
	sw := int32(math.Max(0, float64(clipR-clipX)))
	sh := int32(math.Max(0, float64(clipB-clipY)))

	return sx, sy, sw, sh
}

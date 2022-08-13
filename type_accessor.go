package cimgui

func (v ImVec2) V() (x float32, y float32) {
	return float32(v.x), float32(v.y)
}

func (v ImVec4) V() (x float32, y float32, z float32, w float32) {
	return float32(v.x), float32(v.y), float32(v.z), float32(v.w)
}

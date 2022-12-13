package cimgui

// implement vector operations

// Vec2:

// Add returns a sum of v and another.
func (v ImVec2) Add(another ImVec2) ImVec2 {
	return ImVec2{
		X: v.X + another.X,
		Y: v.Y + another.Y,
	}
}

// Sub returns the vector v - another.
func (v ImVec2) Sub(another ImVec2) ImVec2 {
	return ImVec2{
		X: v.X - another.X,
		Y: v.Y - another.Y,
	}
}

// Mul returns the vector v*k.
func (v ImVec2) Mul(k float32) ImVec2 {
	return ImVec2{
		X: v.X * k,
		Y: v.Y * k,
	}
}

// Div returns the vector v/k.
func (v ImVec2) Div(k float32) ImVec2 {
	return ImVec2{
		X: v.X / k,
		Y: v.Y / k,
	}
}

// Vec4

// Add returns the rectangle r translated by p.
func (v ImVec4) Add(p ImVec2) ImVec4 {
	return ImVec4{
		X: v.X + p.X,
		Y: v.Y + p.Y,
		Z: v.Z + p.X,
		W: v.W + p.Y,
	}
}

// Sub returns the vec4 v translated by -p.
func (v ImVec4) Sub(p ImVec2) ImVec4 {
	return v.Add(p.Mul(-1))
}
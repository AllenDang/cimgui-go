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

// Eq reports whether p and q are equal.
func (v ImVec2) Eq(q ImVec2) bool {
	return v.X == q.X && v.Y == q.Y
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

// Eq reports whether v and r are equal.
func (v ImVec4) Eq(r ImVec4) bool {
	return v.X == r.X && v.Y == r.Y && v.Z == r.Z && v.W == r.W
}

func (v ImVec4) Canon() ImVec4 {
	if v.X > v.Z {
		v.X, v.Z = v.Z, v.X
	}

	if v.Y > v.W {
		v.Y, v.W = v.W, v.Y
	}

	return v
}

func (v ImVec4) In(s ImVec4) bool {
	v = v.Canon()
	s = s.Canon()
	return v.X >= s.X && v.Y >= s.Y && v.Z <= s.Z && v.W <= s.W
}

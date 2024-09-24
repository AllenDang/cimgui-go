package imgui

// implement vector operations

// Vec2:

// Add returns a sum of v and another.
func (v Vec2) Add(another Vec2) Vec2 {
	return Vec2{
		X: v.X + another.X,
		Y: v.Y + another.Y,
	}
}

// Sub returns the vector v - another.
func (v Vec2) Sub(another Vec2) Vec2 {
	return Vec2{
		X: v.X - another.X,
		Y: v.Y - another.Y,
	}
}

// Mul returns the vector v*k.
func (v Vec2) Mul(k float32) Vec2 {
	return Vec2{
		X: v.X * k,
		Y: v.Y * k,
	}
}

// Div returns the vector v/k.
func (v Vec2) Div(k float32) Vec2 {
	return Vec2{
		X: v.X / k,
		Y: v.Y / k,
	}
}

// Vec4

// Add returns the rectangle r translated by p.
func (v Vec4) Add(p Vec2) Vec4 {
	return Vec4{
		X: v.X + p.X,
		Y: v.Y + p.Y,
		Z: v.Z + p.X,
		W: v.W + p.Y,
	}
}

// Sub returns the vec4 v translated by -p.
func (v Vec4) Sub(p Vec2) Vec4 {
	return v.Add(p.Mul(-1))
}

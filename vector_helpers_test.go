package cimgui

import (
	"reflect"
	"testing"
)

func TestVec2_Add(t *testing.T) {
	tests := []struct {
		name string
		p    Vec2
		q    Vec2
		want Vec2
	}{
		{"Add nothing", Vec2{0, 0}, Vec2{0, 0}, Vec2{0, 0}},
		{"Add 1, 5", Vec2{0, 0}, Vec2{1, 5}, Vec2{1, 5}},
		{"(20, 4) + (8, 8)", Vec2{20, 4}, Vec2{8, 8}, Vec2{28, 12}},
		{"positive + negative", Vec2{20, 4}, Vec2{-8, -8}, Vec2{12, -4}},
		{"negative + negative", Vec2{-20, -4}, Vec2{-8, -8}, Vec2{-28, -12}},
		{"float + float", Vec2{20.5, 4.5}, Vec2{8.5, 8.5}, Vec2{29, 13}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Add(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Add(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func TestVec2_Div(t *testing.T) {
	tests := []struct {
		name string
		p    Vec2
		k    float32
		want Vec2
	}{
		{"Simplest", Vec2{0, 0}, 1, Vec2{0, 0}},
		{"Divide by 2", Vec2{2, 2}, 2, Vec2{1, 1}},
		{"divide by -2", Vec2{2, 2}, -2, Vec2{-1, -1}},
		{"divide by float", Vec2{2.5, 2.5}, 2, Vec2{1.25, 1.25}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Div(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Div(%#v) = %v, want %v", tt.p, tt.k, got, tt.want)
			}
		})
	}
}

func TestVec2_Mul(t *testing.T) {
	tests := []struct {
		name string
		p    Vec2
		k    float32
		want Vec2
	}{
		{"multiply by 1", Vec2{1, 1}, 1, Vec2{1, 1}},
		{"multiply by -1", Vec2{1, 1}, -1, Vec2{-1, -1}},
		{"multiply by 2", Vec2{1, 1}, 2, Vec2{2, 2}},
		{"multiply by float", Vec2{1.5, 1.5}, 2, Vec2{3, 3}},
		{"multiply by 0", Vec2{1, 1}, 0, Vec2{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Mul(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Mul(%#v) = %v, want %v", tt.p, tt.k, got, tt.want)
			}
		})
	}
}

func TestVec2_Sub(t *testing.T) {
	tests := []struct {
		name string
		p, q Vec2
		want Vec2
	}{
		{"Simplest", Vec2{0, 0}, Vec2{0, 0}, Vec2{0, 0}},
		{"positive - positive", Vec2{20, 4}, Vec2{8, 8}, Vec2{12, -4}},
		{"positive - negative", Vec2{20, 4}, Vec2{-8, -8}, Vec2{28, 12}},
		{"negative - positive", Vec2{-20, -4}, Vec2{8, 8}, Vec2{-28, -12}},
		{"negative - negative", Vec2{-20, -4}, Vec2{-8, -8}, Vec2{-12, 4}},
		{"float - float", Vec2{20.5, 4.5}, Vec2{8.5, 8.5}, Vec2{12, -4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Sub(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Sub(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func TestVec4_Add(t *testing.T) {
	tests := []struct {
		name string
		p    Vec4
		q    Vec2
		want Vec4
	}{
		{"Simplest", Vec4{0, 0, 0, 0}, Vec2{0, 0}, Vec4{0, 0, 0, 0}},
		{"positive + positive", Vec4{20, 4, 0, 0}, Vec2{8, 8}, Vec4{28, 12, 8, 8}},
		{"positive + negative", Vec4{20, 4, 0, 0}, Vec2{-8, -8}, Vec4{12, -4, -8, -8}},
		{"negative + positive", Vec4{-20, -4, 0, 0}, Vec2{8, 8}, Vec4{-12, 4, 8, 8}},
		{"negative + negative", Vec4{-20, -4, 0, 0}, Vec2{-8, -8}, Vec4{-28, -12, -8, -8}},
		{"float + float", Vec4{20.5, 4.5, 0, 0}, Vec2{8.5, 8.5}, Vec4{29, 13, 8.5, 8.5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Add(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Add(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func TestVec4_Sub(t *testing.T) {
	tests := []struct {
		name string
		p    Vec4
		q    Vec2
		want Vec4
	}{
		{"Simplest", Vec4{0, 0, 0, 0}, Vec2{0, 0}, Vec4{0, 0, 0, 0}},
		{"numbers; no change", Vec4{20, 4, 50, 6}, Vec2{0, 0}, Vec4{20, 4, 50, 6}},
		{"numbers; change", Vec4{20, 4, 50, 6}, Vec2{10, 2}, Vec4{10, 2, 40, 4}},
		{"floats; no change", Vec4{20.5, 4.5, 50.5, 6.5}, Vec2{0, 0}, Vec4{20.5, 4.5, 50.5, 6.5}},
		{"floats; change", Vec4{20.5, 4.5, 50.5, 6.5}, Vec2{10.5, 2.5}, Vec4{10, 2, 40, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Sub(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Sub(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

package cimgui

import (
	"reflect"
	"testing"
)

func TestImVec2_Add(t *testing.T) {
	tests := []struct {
		name string
		p    ImVec2
		q    ImVec2
		want ImVec2
	}{
		{"Add nothing", ImVec2{0, 0}, ImVec2{0, 0}, ImVec2{0, 0}},
		{"Add 1, 5", ImVec2{0, 0}, ImVec2{1, 5}, ImVec2{1, 5}},
		{"(20, 4) + (8, 8)", ImVec2{20, 4}, ImVec2{8, 8}, ImVec2{28, 12}},
		{"positive + negative", ImVec2{20, 4}, ImVec2{-8, -8}, ImVec2{12, -4}},
		{"negative + negative", ImVec2{-20, -4}, ImVec2{-8, -8}, ImVec2{-28, -12}},
		{"float + float", ImVec2{20.5, 4.5}, ImVec2{8.5, 8.5}, ImVec2{29, 13}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Add(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Add(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func TestImVec2_Div(t *testing.T) {
	tests := []struct {
		name string
		p    ImVec2
		k    float32
		want ImVec2
	}{
		{"Simplest", ImVec2{0, 0}, 1, ImVec2{0, 0}},
		{"Divide by 2", ImVec2{2, 2}, 2, ImVec2{1, 1}},
		{"divide by -2", ImVec2{2, 2}, -2, ImVec2{-1, -1}},
		{"divide by float", ImVec2{2.5, 2.5}, 2, ImVec2{1.25, 1.25}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Div(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Div(%#v) = %v, want %v", tt.p, tt.k, got, tt.want)
			}
		})
	}
}

func TestImVec2_Mul(t *testing.T) {
	tests := []struct {
		name string
		p    ImVec2
		k    float32
		want ImVec2
	}{
		{"multiply by 1", ImVec2{1, 1}, 1, ImVec2{1, 1}},
		{"multiply by -1", ImVec2{1, 1}, -1, ImVec2{-1, -1}},
		{"multiply by 2", ImVec2{1, 1}, 2, ImVec2{2, 2}},
		{"multiply by float", ImVec2{1.5, 1.5}, 2, ImVec2{3, 3}},
		{"multiply by 0", ImVec2{1, 1}, 0, ImVec2{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Mul(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Mul(%#v) = %v, want %v", tt.p, tt.k, got, tt.want)
			}
		})
	}
}

func TestImVec2_Sub(t *testing.T) {
	tests := []struct {
		name string
		p, q ImVec2
		want ImVec2
	}{
		{"Simplest", ImVec2{0, 0}, ImVec2{0, 0}, ImVec2{0, 0}},
		{"positive - positive", ImVec2{20, 4}, ImVec2{8, 8}, ImVec2{12, -4}},
		{"positive - negative", ImVec2{20, 4}, ImVec2{-8, -8}, ImVec2{28, 12}},
		{"negative - positive", ImVec2{-20, -4}, ImVec2{8, 8}, ImVec2{-28, -12}},
		{"negative - negative", ImVec2{-20, -4}, ImVec2{-8, -8}, ImVec2{-12, 4}},
		{"float - float", ImVec2{20.5, 4.5}, ImVec2{8.5, 8.5}, ImVec2{12, -4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Sub(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Sub(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func TestImVec4_Add(t *testing.T) {
	tests := []struct {
		name string
		p    ImVec4
		q    ImVec2
		want ImVec4
	}{
		{"Simplest", ImVec4{0, 0, 0, 0}, ImVec2{0, 0}, ImVec4{0, 0, 0, 0}},
		{"positive + positive", ImVec4{20, 4, 0, 0}, ImVec2{8, 8}, ImVec4{28, 12, 8, 8}},
		{"positive + negative", ImVec4{20, 4, 0, 0}, ImVec2{-8, -8}, ImVec4{12, -4, -8, -8}},
		{"negative + positive", ImVec4{-20, -4, 0, 0}, ImVec2{8, 8}, ImVec4{-12, 4, 8, 8}},
		{"negative + negative", ImVec4{-20, -4, 0, 0}, ImVec2{-8, -8}, ImVec4{-28, -12, -8, -8}},
		{"float + float", ImVec4{20.5, 4.5, 0, 0}, ImVec2{8.5, 8.5}, ImVec4{29, 13, 8.5, 8.5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Add(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Add(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func TestImVec4_Sub(t *testing.T) {
	tests := []struct {
		name string
		p    ImVec4
		q    ImVec2
		want ImVec4
	}{
		{"Simplest", ImVec4{0, 0, 0, 0}, ImVec2{0, 0}, ImVec4{0, 0, 0, 0}},
		{"numbers; no change", ImVec4{20, 4, 50, 6}, ImVec2{0, 0}, ImVec4{20, 4, 50, 6}},
		{"numbers; change", ImVec4{20, 4, 50, 6}, ImVec2{10, 2}, ImVec4{10, 2, 40, 4}},
		{"floats; no change", ImVec4{20.5, 4.5, 50.5, 6.5}, ImVec2{0, 0}, ImVec4{20.5, 4.5, 50.5, 6.5}},
		{"floats; change", ImVec4{20.5, 4.5, 50.5, 6.5}, ImVec2{10.5, 2.5}, ImVec4{10, 2, 40, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.Sub(tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%#v.Sub(%#v) = %v, want %v", tt.p, tt.q, got, tt.want)
			}
		})
	}
}

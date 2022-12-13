package cimgui

import (
	"reflect"
	"testing"
)

func TestImVec2_Add(t *testing.T) {
	type fields struct {
		X float32
		Y float32
	}

	type args struct {
		another ImVec2
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ImVec2
	}{
		{"Add nothing", fields{0, 0}, args{ImVec2{0, 0}}, ImVec2{0, 0}},
		{"Add 1, 5", fields{0, 0}, args{ImVec2{1, 5}}, ImVec2{1, 5}},
		{"(20, 4) + (8, 8)", fields{20, 4}, args{ImVec2{8, 8}}, ImVec2{28, 12}},
		{"positive + negative", fields{20, 4}, args{ImVec2{-8, -8}}, ImVec2{12, -4}},
		{"negative + negative", fields{-20, -4}, args{ImVec2{-8, -8}}, ImVec2{-28, -12}},
		{"float + float", fields{20.5, 4.5}, args{ImVec2{8.5, 8.5}}, ImVec2{29, 13}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Add(tt.args.another); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec2_Div(t *testing.T) {
	type fields struct {
		X float32
		Y float32
	}

	type args struct {
		k float32
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ImVec2
	}{
		{"Simplest", fields{0, 0}, args{1}, ImVec2{0, 0}},
		{"Divide by 2", fields{2, 2}, args{2}, ImVec2{1, 1}},
		{"divide by -2", fields{2, 2}, args{-2}, ImVec2{-1, -1}},
		{"divide by float", fields{2.5, 2.5}, args{2}, ImVec2{1.25, 1.25}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Div(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec2_Mul(t *testing.T) {
	type fields struct {
		X float32
		Y float32
	}

	type args struct {
		k float32
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ImVec2
	}{
		{"multiply by 1", fields{1, 1}, args{1}, ImVec2{1, 1}},
		{"multiply by -1", fields{1, 1}, args{-1}, ImVec2{-1, -1}},
		{"multiply by 2", fields{1, 1}, args{2}, ImVec2{2, 2}},
		{"multiply by float", fields{1.5, 1.5}, args{2}, ImVec2{3, 3}},
		{"multiply by 0", fields{1, 1}, args{0}, ImVec2{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Mul(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec2_Sub(t *testing.T) {
	type fields struct {
		X float32
		Y float32
	}

	type args struct {
		another ImVec2
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ImVec2
	}{
		{"Simplest", fields{0, 0}, args{ImVec2{0, 0}}, ImVec2{0, 0}},
		{"positive - positive", fields{20, 4}, args{ImVec2{8, 8}}, ImVec2{12, -4}},
		{"positive - negative", fields{20, 4}, args{ImVec2{-8, -8}}, ImVec2{28, 12}},
		{"negative - positive", fields{-20, -4}, args{ImVec2{8, 8}}, ImVec2{-28, -12}},
		{"negative - negative", fields{-20, -4}, args{ImVec2{-8, -8}}, ImVec2{-12, 4}},
		{"float - float", fields{20.5, 4.5}, args{ImVec2{8.5, 8.5}}, ImVec2{12, -4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec2{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := v.Sub(tt.args.another); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec4_Add(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}

	type args struct {
		p ImVec2
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ImVec4
	}{
		{"Simplest", fields{0, 0, 0, 0}, args{ImVec2{0, 0}}, ImVec4{0, 0, 0, 0}},
		{"positive + positive", fields{20, 4, 0, 0}, args{ImVec2{8, 8}}, ImVec4{28, 12, 8, 8}},
		{"positive + negative", fields{20, 4, 0, 0}, args{ImVec2{-8, -8}}, ImVec4{12, -4, -8, -8}},
		{"negative + positive", fields{-20, -4, 0, 0}, args{ImVec2{8, 8}}, ImVec4{-12, 4, 8, 8}},
		{"negative + negative", fields{-20, -4, 0, 0}, args{ImVec2{-8, -8}}, ImVec4{-28, -12, -8, -8}},
		{"float + float", fields{20.5, 4.5, 0, 0}, args{ImVec2{8.5, 8.5}}, ImVec4{29, 13, 8.5, 8.5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := v.Add(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec4_Canon(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}

	tests := []struct {
		name   string
		fields fields
		want   ImVec4
	}{
		{"Simplest", fields{0, 0, 0, 0}, ImVec4{0, 0, 0, 0}},
		{"no need to reverse anything", fields{0, 0, 20, 4}, ImVec4{0, 0, 20, 4}},
		{"need to reverse vertically", fields{0, 4, 20, 0}, ImVec4{0, 0, 20, 4}},
		{"need to reverse horizontally", fields{20, 0, 0, 4}, ImVec4{0, 0, 20, 4}},
		{"need to reverse both", fields{20, 4, 0, 0}, ImVec4{0, 0, 20, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := v.Canon(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Canon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec4_Eq(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}

	type args struct {
		r ImVec4
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Simplest", fields{0, 0, 0, 0}, args{ImVec4{0, 0, 0, 0}}, true},
		{"numbers; equal", fields{20, 4, 50, 6}, args{ImVec4{20, 4, 50, 6}}, true},
		{"numbers; not equal", fields{20, 4, 50, 6}, args{ImVec4{20, 4, 50, 7}}, false},
		{"floats; equal", fields{20.5, 4.5, 50.5, 6.5}, args{ImVec4{20.5, 4.5, 50.5, 6.5}}, true},
		{"floats; not equal", fields{20.5, 4.5, 50.5, 6.5}, args{ImVec4{20.5, 4.5, 50.5, 7.5}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := v.Eq(tt.args.r); got != tt.want {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec4_In(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}
	type args struct {
		s ImVec4
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Simplest", fields{0, 0, 0, 0}, args{ImVec4{0, 0, 0, 0}}, true},
		{"numbers; already cannon; in", fields{-10, -20, 20, 10}, args{ImVec4{-4, -10, 8, 5}}, true},
		{"numbers; already cannon; out", fields{-10, -20, 20, 10}, args{ImVec4{-4, -10, 8, 15}}, false},
		{"numbers; not cannon; in", fields{10, 20, -20, -10}, args{ImVec4{-4, -10, 8, 5}}, true},
		{"numbers; not cannon; in (case 2)", fields{10, 20, -20, -10}, args{ImVec4{-4, -10, 8, 15}}, true},
		{"numbers; not cannon; out", fields{10, 20, -20, -10}, args{ImVec4{-4, -10, 8, 25}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := v.In(tt.args.s); got != tt.want {
				t.Errorf("In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImVec4_Sub(t *testing.T) {
	type fields struct {
		X float32
		Y float32
		Z float32
		W float32
	}

	type args struct {
		p ImVec2
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   ImVec4
	}{
		{"Simplest", fields{0, 0, 0, 0}, args{ImVec2{0, 0}}, ImVec4{0, 0, 0, 0}},
		{"numbers; no change", fields{20, 4, 50, 6}, args{ImVec2{0, 0}}, ImVec4{20, 4, 50, 6}},
		{"numbers; change", fields{20, 4, 50, 6}, args{ImVec2{10, 2}}, ImVec4{10, 2, 40, 4}},
		{"floats; no change", fields{20.5, 4.5, 50.5, 6.5}, args{ImVec2{0, 0}}, ImVec4{20.5, 4.5, 50.5, 6.5}},
		{"floats; change", fields{20.5, 4.5, 50.5, 6.5}, args{ImVec2{10.5, 2.5}}, ImVec4{10, 2, 40, 4}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ImVec4{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
				W: tt.fields.W,
			}
			if got := v.Sub(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

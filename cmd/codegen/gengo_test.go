package main

import "testing"

func TestIsCallbackTypedef(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"variant with no extra * and ()", "void MarkdownFormalCallback(const MarkdownFormatInfo& markdownFormatInfo_,bool start_);", true},
		{"variant from cimgui", "int(*)(ImGuiInputTextCallbackData* data);", true},
		{"some random structdef that shouldn't be true", "struct ImGuiPlatformImeData", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCallbackTypedef(tt.s); got != tt.want {
				t.Errorf("IsCallbackTypedef() = %v, want %v", got, tt.want)
			}
		})
	}
}

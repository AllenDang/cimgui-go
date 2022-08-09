package cimgui

import "testing"

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v != "1.89 WIP" {
		t.Error("imgui version should be 1.88")
	}
}

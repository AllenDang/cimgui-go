package cimgui

import (
	"testing"
)

func TestGetVersion(t *testing.T) {
	v := GetVersion()
	if v != "1.89 WIP" {
		t.Error("imgui version should be 1.88")
	}
}

func TestSetIOCofigFlags(t *testing.T) {
	CreateContext(0)
	defer DestroyContext(0)

	io := GetIO()
	if io == 0 {
		t.Error("get io failed")
	}

	io.SetBackendFlags(ImGuiConfigFlags_DockingEnable)

	flags := io.GetBackendFlags()
	if flags != ImGuiConfigFlags_DockingEnable {
		t.Error("set io backend flags failed")
		t.Errorf("expect: %v got %v", ImGuiConfigFlags_DockingEnable, flags)
	}
}

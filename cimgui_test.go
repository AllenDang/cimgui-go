package cimgui

import (
	"testing"
)

func TestSetIOCofigFlags(t *testing.T) {
	CreateContext(0)
	defer DestroyContext(0)

	io := GetIO()
	if io == 0 {
		t.Error("get io failed")
	}

	io.SetBackendFlags(ImGuiBackendFlags_RendererHasVtxOffset)

	flags := io.GetBackendFlags()
	if flags != ImGuiBackendFlags_RendererHasVtxOffset {
		t.Error("set io backend flags failed")
		t.Errorf("expect: %v got %v", ImGuiBackendFlags_RendererHasVtxOffset, flags)
	}
}

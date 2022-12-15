package cimgui

import (
	"testing"
)

func TestSetIOConfigFlags(t *testing.T) {
	CreateContext()
	defer DestroyContext()

	io := GetIO()
	if io == 0 {
		t.Error("get io failed")
	}

	io.SetBackendFlags(BackendFlags_RendererHasVtxOffset)

	flags := io.GetBackendFlags()
	if flags != BackendFlags_RendererHasVtxOffset {
		t.Error("set io backend flags failed")
		t.Errorf("expect: %v got %v", BackendFlags_RendererHasVtxOffset, flags)
	}
}

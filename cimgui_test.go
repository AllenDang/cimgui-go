package imgui

import (
	"testing"
)

func TestSetIOConfigFlags(t *testing.T) {
	CreateContext()
	defer DestroyContext()

	io := CurrentIO()
	if io.data == 0 {
		t.Error("get io failed")
	}

	io.SetBackendFlags(BackendFlagsRendererHasVtxOffset)

	flags := io.BackendFlags()
	if flags != BackendFlagsRendererHasVtxOffset {
		t.Error("set io backend flags failed")
		t.Errorf("expect: %v got %v", BackendFlagsRendererHasVtxOffset, flags)
	}
}

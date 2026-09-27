package glfwbackend

/*
#include "glfw_backend.h"
#include <stdlib.h>
*/
import "C"
import (
	"sync"
	"unsafe"

	"github.com/AllenDang/cimgui-go/internal"
)

// GLFW native input action constants.
const (
	GLFWRelease = 0
	GLFWPress   = 1
	GLFWRepeat  = 2
)

// GLFWMonitor is a thin wrapper around a *GLFWmonitor.
type GLFWMonitor struct {
	ptr uintptr
}

func (m *GLFWMonitor) handle() *C.GLFWmonitor {
	return (*C.GLFWmonitor)(unsafe.Pointer(m.ptr))
}

// GLFWWindow is a manually created GLFW window that does NOT initialize ImGui.
// It is intended for secondary outputs (e.g. viewports) while sharing the
// OpenGL context of the main application backend.
type GLFWWindow struct {
	window uintptr
	cb     *glfwWindowCallbacks
}

type glfwWindowCallbacks struct {
	key   KeyCallback
	close CloseCallback
	size  SizeCallback
}

// KeyCallback is invoked on GLFW key events for a GLFWWindow.
type KeyCallback func(key, scancode, action, mods int)

// CloseCallback is invoked when a GLFWWindow is requested to close.
type CloseCallback func()

// SizeCallback is invoked when a GLFWWindow is resized.
type SizeCallback func(width, height int)

var (
	viewportCBMu sync.Mutex
	viewportCBs  = map[uintptr]*glfwWindowCallbacks{}
)

//export goViewportKeyCallback
func goViewportKeyCallback(w unsafe.Pointer, k, s, a, m C.int) {
	viewportCBMu.Lock()
	cb, ok := viewportCBs[uintptr(w)]
	viewportCBMu.Unlock()
	if ok && cb.key != nil {
		cb.key(int(k), int(s), int(a), int(m))
	}
}

//export goViewportCloseCallback
func goViewportCloseCallback(w unsafe.Pointer) {
	viewportCBMu.Lock()
	cb, ok := viewportCBs[uintptr(w)]
	viewportCBMu.Unlock()
	if ok && cb.close != nil {
		cb.close()
	}
}

//export goViewportSizeCallback
func goViewportSizeCallback(w unsafe.Pointer, width, height C.int) {
	viewportCBMu.Lock()
	cb, ok := viewportCBs[uintptr(w)]
	viewportCBMu.Unlock()
	if ok && cb.size != nil {
		cb.size(int(width), int(height))
	}
}

func (w *GLFWWindow) handle() *C.GLFWwindow {
	return (*C.GLFWwindow)(unsafe.Pointer(w.window))
}

// NewGLFWWindow creates a new window sharing the OpenGL context of `share`
// (which may be nil). The native pointer of the main application backend can be
// obtained through GLFWBackend.WindowPointer().
func NewGLFWWindow(title string, width, height int, share unsafe.Pointer) *GLFWWindow {
	var sharePtr *C.GLFWwindow
	if share != nil {
		sharePtr = (*C.GLFWwindow)(share)
	}
	titleArg, titleFin := internal.WrapString[C.char](title)
	defer titleFin()

	win := C.igCreateSharedGLFWWindow((*C.char)(titleArg), C.int(width), C.int(height), sharePtr)
	if win == nil {
		return nil
	}
	w := &GLFWWindow{window: uintptr(unsafe.Pointer(win)), cb: &glfwWindowCallbacks{}}
	viewportCBMu.Lock()
	viewportCBs[w.window] = w.cb
	viewportCBMu.Unlock()
	return w
}

// WindowPointer exposes the underlying native window pointer of a GLFWBackend so
// it can be used as a shared context when creating a GLFWWindow.
func (b *GLFWBackend) WindowPointer() unsafe.Pointer {
	return unsafe.Pointer(b.window)
}

// MakeContextCurrent binds this window's OpenGL context to the calling thread.
func (w *GLFWWindow) MakeContextCurrent() { C.igGLFWWindow_MakeContextCurrent(w.handle()) }

// DetachCurrentContext unbinds any OpenGL context from the calling thread.
func (w *GLFWWindow) DetachCurrentContext() { C.igGLFWWindow_DetachCurrentContext() }

// SwapBuffers swaps the front and back buffers of this window.
func (w *GLFWWindow) SwapBuffers() { C.igGLFWWindow_SwapBuffers(w.handle()) }

// Destroy closes the window. It must not be called on the main backend window.
func (w *GLFWWindow) Destroy() {
	C.igGLFWWindow_DestroyWindow(w.handle())
	viewportCBMu.Lock()
	delete(viewportCBs, w.window)
	viewportCBMu.Unlock()
}

// GetFramebufferSize returns the window's framebuffer size in pixels.
func (w *GLFWWindow) GetFramebufferSize() (int32, int32) {
	var width, height C.int
	C.igGLFWWindow_GetFramebufferSize(w.handle(), &width, &height)
	return int32(width), int32(height)
}

// GetDisplaySize returns the window's size in screen coordinates.
func (w *GLFWWindow) GetDisplaySize() (int32, int32) {
	var width, height C.int
	C.igGLFWWindow_GetDisplaySize(w.handle(), &width, &height)
	return int32(width), int32(height)
}

// Show shows or hides the window.
func (w *GLFWWindow) Show(v bool) {
	C.igGLFWWindow_Show(w.handle(), C.int(boolToInt(v)))
}

// SetAttrib sets a window attribute (e.g. GLFWWindowAutoIconify).
func (w *GLFWWindow) SetAttrib(attrib, value int) {
	C.igGLFWWindow_SetAttrib(w.handle(), C.int(attrib), C.int(value))
}

// ShouldClose reports whether the window close flag is set.
func (w *GLFWWindow) ShouldClose() bool {
	return C.igGLFWWindow_GetShouldClose(w.handle()) != 0
}

// SetShouldClose sets the window close flag.
func (w *GLFWWindow) SetShouldClose(v bool) {
	C.igGLFWWindow_SetShouldClose(w.handle(), C.bool(v))
}

// GetMonitor returns the monitor the window is fullscreen on, or nil if windowed.
func (w *GLFWWindow) GetMonitor() *GLFWMonitor {
	m := C.igGLFWWindow_GetMonitor(w.handle())
	if m == nil {
		return nil
	}
	return &GLFWMonitor{ptr: uintptr(unsafe.Pointer(m))}
}

// SetMonitor switches the window to fullscreen on `m` (or windowed when m is nil).
func (w *GLFWWindow) SetMonitor(m *GLFWMonitor, x, y, width, height, refresh int) {
	var mptr *C.GLFWmonitor
	if m != nil {
		mptr = m.handle()
	}
	C.igGLFWWindow_SetMonitor(w.handle(), mptr, C.int(x), C.int(y), C.int(width), C.int(height), C.int(refresh))
}

// SetKeyCallback installs a per-window key callback.
func (w *GLFWWindow) SetKeyCallback(cb KeyCallback) {
	viewportCBMu.Lock()
	if w.cb == nil {
		w.cb = &glfwWindowCallbacks{}
		viewportCBs[w.window] = w.cb
	}
	w.cb.key = cb
	viewportCBMu.Unlock()
	C.igGLFWWindow_SetKeyCallbackEx(w.handle())
}

// SetCloseCallback installs a per-window close callback.
func (w *GLFWWindow) SetCloseCallback(cb CloseCallback) {
	viewportCBMu.Lock()
	if w.cb == nil {
		w.cb = &glfwWindowCallbacks{}
		viewportCBs[w.window] = w.cb
	}
	w.cb.close = cb
	viewportCBMu.Unlock()
	C.igGLFWWindow_SetCloseCallbackEx(w.handle())
}

// SetSizeCallback installs a per-window resize callback.
func (w *GLFWWindow) SetSizeCallback(cb SizeCallback) {
	viewportCBMu.Lock()
	if w.cb == nil {
		w.cb = &glfwWindowCallbacks{}
		viewportCBs[w.window] = w.cb
	}
	w.cb.size = cb
	viewportCBMu.Unlock()
	C.igGLFWWindow_SetSizeCallbackEx(w.handle())
}

// GetCurrentContext returns the GLFWWindow whose context is current on the
// calling thread, or nil.
func GetCurrentContext() *GLFWWindow {
	w := C.igGLFWWindow_GetCurrentContext()
	if w == nil {
		return nil
	}
	return &GLFWWindow{window: uintptr(unsafe.Pointer(w))}
}

// GetPrimaryMonitor returns the primary monitor.
func GetPrimaryMonitor() *GLFWMonitor {
	m := C.igGLFW_GetPrimaryMonitor()
	if m == nil {
		return nil
	}
	return &GLFWMonitor{ptr: uintptr(unsafe.Pointer(m))}
}

// GetMonitors returns all currently connected monitors.
func GetMonitors() []*GLFWMonitor {
	var count C.int
	monitors := C.igGLFW_GetMonitors(&count)
	if monitors == nil || count == 0 {
		return nil
	}
	monPtrs := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(monitors)), int(count))
	result := make([]*GLFWMonitor, int(count))
	for i, p := range monPtrs {
		result[i] = &GLFWMonitor{ptr: uintptr(p)}
	}
	return result
}

// Name returns the human-readable name of the monitor.
func (m *GLFWMonitor) Name() string {
	name := C.igGLFW_GetMonitorName(m.handle())
	if name == nil {
		return ""
	}
	return C.GoString(name)
}

// GetVideoMode returns the current video mode of a monitor.
func GetVideoMode(m *GLFWMonitor) (int, int, int) {
	var w, h, r C.int
	C.igGLFW_GetVideoMode(m.handle(), &w, &h, &r)
	return int(w), int(h), int(r)
}

func boolToInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

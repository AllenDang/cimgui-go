//go:build linux && cgo

// Package drmeglbackend provides a Dear ImGui backend for Linux DRM/KMS + EGL.
//
// It is intended for embedded or kiosk-like environments where rendering happens
// directly on a TTY without X11 or Wayland. Unlike the GLFW/SDL backends, this
// backend does not expose a real desktop window, so several window-management
// methods are implemented as compatibility no-ops.
package drmeglbackend

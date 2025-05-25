package glfw

// #cgo amd64,linux LDFLAGS: ${SRCDIR}/../../lib/linux/x64/libglfw3.a -ldl -lGL -lX11
// #cgo amd64,windows LDFLAGS: -L${SRCDIR}/../../lib/windows/x64 -l:libglfw3.a -lgdi32 -lopengl32 -limm32
// #cgo darwin LDFLAGS: -framework Cocoa -framework IOKit -framework CoreVideo
// #cgo amd64,darwin LDFLAGS: ${SRCDIR}/../../lib/macos/x64/libglfw3.a
// #cgo arm64,darwin LDFLAGS: ${SRCDIR}/../../lib/macos/arm64/libglfw3.a
// #cgo !gles2,darwin LDFLAGS: -framework OpenGL
// #cgo gles2,darwin LDFLAGS: -lGLESv2
// #cgo CPPFLAGS: -DCIMGUI_USE_GLFW
// #include "../../thirdparty/glfw/include/GLFW/glfw3.h" // Will drag system OpenGL headers
import "C"

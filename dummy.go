//go:build required
// +build required

package imgui

import (

	_ "github.com/AllenDang/cimgui-go/cwrappers"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/backends"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/build"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/example/grep"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/example/snippets"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/example/timer"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/include/boost"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/performance"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/performance/config"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/src"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/captures"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/c_compiler_checks"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/collate_info"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/concepts"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/config_info"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/de_fuzz"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/named_subexpressions"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/object_cache"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/pathology"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/regress"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/static_mutex"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/test/unicode"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuiColorTextEdit/vendor/regex/tools/generate"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_allegro5"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_android_opengl3"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_glfw_opengl2"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_glfw_opengl3"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_glfw_vulkan"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_glfw_wgpu"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_glut_opengl2"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_null"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl2_directx11"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl2_opengl2"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl2_opengl3"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl2_sdlrenderer2"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl2_vulkan"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl3_opengl3"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl3_sdlrenderer3"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_sdl3_vulkan"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_win32_directx10"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_win32_directx11"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_win32_directx12"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_win32_directx9"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/example_win32_opengl3"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/libs/emscripten"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/libs/glfw/include/GLFW"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/examples/libs/usynergy"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui_markdown"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/misc/cpp"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/misc/fonts"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/misc/freetype"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imgui/misc/single_file"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuizmo"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuizmo/example"
	_ "github.com/AllenDang/cimgui-go/cwrappers/ImGuizmo/vcpkg-example"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/example"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/b64"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/cgns"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/chartdir"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/cityhash"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/clapack"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/freeimage"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/gettimeofday"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/graphicsmagick"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/gts"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/igraph"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/libaiff"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/libmspack"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/libu2f-server"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/libuuid"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/modp-base64"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/openblas"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/ports/ragel"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/scripts/test_ports/vcpkg-ci-ankurvdev-embedresource/project"
	_ "github.com/AllenDang/cimgui-go/cwrappers/imnodes/vcpkg/scripts/test_ports/vcpkg-ci-soci/project"
	_ "github.com/AllenDang/cimgui-go/cwrappers/implot"
)

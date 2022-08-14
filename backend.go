package cimgui

// #include "backend_wrapper.h"
import "C"
import "github.com/go-gl/glfw/v3.3/glfw"

func ImGuiImplGlfwInitForOpenGL(window *glfw.Window, installCallbacks bool) bool {
	return C.ImGui_ImplGlfw_InitForOpenGL((*C.GLFWwindow)(window.Handle()), C.bool(installCallbacks)) == C.bool(true)
}

func ImGuiImplGlfwNewFrame() {
	C.ImGui_ImplGlfw_NewFrame()
}

func ImGuiImplGlfwShutdown() {
	C.ImGui_ImplGlfw_Shutdown()
}

func ImGuiImplOpenGL3Init(glslVersion string) bool {
	glslVersionArg, glslVersionFin := wrapString(glslVersion)
	defer glslVersionFin()

	return C.ImGui_ImplOpenGL3_Init(glslVersionArg) == C.bool(true)
}

func ImGuiImplOpenGL3Shutdown() {
	C.ImGui_ImplOpenGL3_Shutdown()
}

func ImGuiImplOpenGL3NewFrame() {
	C.ImGui_ImplOpenGL3_NewFrame()
}

func ImGui_ImplOpenGL3RenderDrawData(drawData *ImDrawData) {
	C.ImGui_ImplOpenGL3_RenderDrawData((*C.ImDrawData)(drawData))
}

func ImGui_ImplOpenGL3CreateFontsTexture() {
	C.ImGui_ImplOpenGL3_CreateFontsTexture()
}

func ImGui_ImplOpenGL3DestroyFontsTexture() {
	C.ImGui_ImplOpenGL3_DestroyFontsTexture()
}

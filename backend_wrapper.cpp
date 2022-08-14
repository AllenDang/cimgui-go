#include "backend_wrapper.h"

bool ImGui_ImplGlfw_InitForOpenGL(GLFWwindow *window, bool install_callbacks) { return ImGui_ImplGlfw_InitForOpenGL(window, install_callbacks); }

void ImGui_ImplGlfw_NewFrame(void) { ImGui_ImplGlfw_NewFrame(); }

void ImGui_ImplGlfw_Shutdown(void) { ImGui_ImplGlfw_Shutdown(); }

bool ImGui_ImplOpenGL3_Init(const char *glsl_version) { return ImGui_ImplOpenGL3_Init(glsl_version); }
void ImGui_ImplOpenGL3_Shutdown(void) { ImGui_ImplOpenGL3_Shutdown(); }
void ImGui_ImplOpenGL3_NewFrame(void) { ImGui_ImplOpenGL3_NewFrame(); }
void ImGui_ImplOpenGL3_RenderDrawData(ImDrawData *draw_data) { ImGui_ImplOpenGL3_RenderDrawData(draw_data); }
bool ImGui_ImplOpenGL3_CreateFontsTexture(void) { return ImGui_ImplOpenGL3_CreateFontsTexture(); }
void ImGui_ImplOpenGL3_DestroyFontsTexture(void) { ImGui_ImplOpenGL3_DestroyFontsTexture(); }

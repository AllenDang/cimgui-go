#pragma once

#define CIMGUI_USE_GFLW
#define CIMGUI_USE_OPENGL3
#include "cimgui/cimgui.h"
#include "cimgui/cimgui_impl.h"

#ifdef __cplusplus
extern "C" {
#endif

typedef struct GLFWwindow GLFWwindow;

extern bool ImGui_ImplGlfw_InitForOpenGL(GLFWwindow *window, bool install_callbacks);
extern void ImGui_ImplGlfw_NewFrame(void);
extern void ImGui_ImplGlfw_Shutdown(void);

extern bool ImGui_ImplOpenGL3_Init(const char *glsl_version);
extern void ImGui_ImplOpenGL3_Shutdown(void);
extern void ImGui_ImplOpenGL3_NewFrame(void);
extern void ImGui_ImplOpenGL3_RenderDrawData(ImDrawData *draw_data);
extern bool ImGui_ImplOpenGL3_CreateFontsTexture(void);
extern void ImGui_ImplOpenGL3_DestroyFontsTexture(void);

#ifdef __cplusplus
}
#endif

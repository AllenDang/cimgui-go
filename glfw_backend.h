#pragma once

#include "cimgui_wrapper.h"
#include "thirdparty/glfw/include/GLFW/glfw3.h" // Will drag system OpenGL headers

#ifdef __cplusplus
extern "C" {
#endif

typedef int GLFWWindowFlags;
enum GLFWWindowFlags_ {
  GLFWWindowNone = 0,
  GLFWWindowResizable = GLFW_RESIZABLE,
  GLFWWindowMaximized = GLFW_MAXIMIZED,
  GLFWWindowTransparentFramebuffer = GLFW_TRANSPARENT_FRAMEBUFFER,
  GLFWWindowDecorated= GLFW_DECORATED,
  GLFWWindowVisible = GLFW_VISIBLE,
};

typedef struct CImage {
int width;
int height;
unsigned char* pixels;
}CImage ;

typedef struct GLFWwindow GLFWwindow;
typedef struct GLFWmonitor GLFWmonitor;
typedef struct GLFWimage GLFWimage;
struct GLFWwindow;
struct GLFWmonitor;
struct GLFWimage;

typedef void (*VoidCallback)();

extern void igSetBgColor(ImVec4 color);
extern void igSetTargetFPS(unsigned int fps);
extern int igInitGLFW();
extern GLFWwindow *igCreateGLFWWindow(const char *title, int width, int height,
                                      VoidCallback afterCreateContext);
extern void igRunLoop(GLFWwindow *window, VoidCallback loop, VoidCallback beforeRender, VoidCallback afterRender,
                      VoidCallback beforeDestroyContext);
extern void igGLFWWindow_GetDisplaySize(GLFWwindow *window, int *width, int *height);
extern void igGLFWWindow_SetWindowPos(GLFWwindow *window, int x, int y);
extern void igGLFWWindow_GetWindowPos(GLFWwindow *window, int *x, int *y);
extern void igGLFWWindow_SetShouldClose(GLFWwindow *window, int value);
extern void igGLFWWindow_SetDropCallbackCB(GLFWwindow *window);
extern void igGLFWWindow_SetSize(GLFWwindow *window, int width, int height);
extern void igGLFWWindow_SetTitle(GLFWwindow *window, const char *title);
extern void igGLFWWindow_SetSizeLimits(GLFWwindow *window, int minWidth, int minHeight, int maxWidth, int maxHeight);
extern void igGLFWWindow_SetCloseCallback(GLFWwindow *window);
extern void igGLFWWindow_SetKeyCallback(GLFWwindow *window);
extern void igGLFWWindow_SetSizeCallback(GLFWwindow *window);
extern void igGLFWWindow_SetIcon(GLFWwindow *window, int count, CImage* images);
extern void igRefresh();
extern ImTextureID igCreateTexture(unsigned char *pixels, int width, int height);
extern void igDeleteTexture(ImTextureID id);
extern void igWindowHint(GLFWWindowFlags hint, int value);

extern void dropCallback(int, char **);
extern void closeCallback(GLFWwindow *window);
extern void keyCallback(int key, int scancode, int action, int mods);
extern void sizeCallback(int width, int height);

#ifdef __cplusplus
}
#endif

#pragma once

#ifdef __cplusplus
extern "C" {
#endif

typedef struct GLFWwindow GLFWwindow;
typedef struct GLFWmonitor GLFWmonitor;
struct GLFWwindow;
struct GLFWmonitor;

typedef void (*VoidCallback)();

extern GLFWwindow *igCreateGlfwWindow(const char *title, int width, int height);
extern void igRunLoop(GLFWwindow *window, VoidCallback loop);

#ifdef __cplusplus
}
#endif

#if defined(CIMGUI_GO_USE_GLFW)
#define GL_SILENCE_DEPRECATION
#define CIMGUI_USE_GLFW
#define CIMGUI_USE_OPENGL3

#include "glfw_backend.h"
#include "cimgui/cimgui.h"
#include "cimgui/cimgui_impl.h"
#include "thirdparty/glfw/include/GLFW/glfw3.h" // Will drag system OpenGL headers
#include <cstdlib>


// [Win32] Our example includes a copy of glfw3.lib pre-compiled with VS2010 to
// maximize ease of testing and compatibility with old VS compilers. To link
// with VS2010-era libraries, VS2015+ requires linking with
// legacy_stdio_definitions.lib, which we do using this pragma. Your own project
// should not be affected, as you are likely to link with a newer binary of GLFW
// that is adequate for your version of Visual Studio.
#if defined(_MSC_VER) && (_MSC_VER >= 1900) && !defined(IMGUI_DISABLE_WIN32_FUNCTIONS)
#pragma comment(lib, "legacy_stdio_definitions")
#endif

#define MAX_EXTRA_FRAME_COUNT 15
unsigned int glfw_target_fps = 30;
int extra_frame_count = MAX_EXTRA_FRAME_COUNT;

ImVec4 clear_color = *ImVec4_ImVec4_Float(0.45, 0.55, 0.6, 1.0);

void glfw_render(GLFWwindow *window, VoidCallback renderLoop);

void igSetBgColor(ImVec4 color) { clear_color = color; }

void igSetTargetFPS(unsigned int fps) { glfw_target_fps = fps; }

static void glfw_error_callback(int error, const char *description) {
  fprintf(stderr, "Glfw Error %d: %s\n", error, description);
}

void glfw_window_refresh_callback(GLFWwindow *window) {
  VoidCallback loopFunc = (VoidCallback)(glfwGetWindowUserPointer(window));
  glfw_render(window, loopFunc);
}

int igInitGLFW() {
    glfwSetErrorCallback(glfw_error_callback);
    return glfwInit();
}

GLFWwindow *igCreateGLFWWindow(const char *title, int width, int height,
                               VoidCallback afterCreateContext) {
    // Decide GL+GLSL versions
#if defined(IMGUI_IMPL_OPENGL_ES2)
  // GL ES 2.0 + GLSL 100
  const char *glsl_version = "#version 100";
  glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 2);
  glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 0);
  glfwWindowHint(GLFW_CLIENT_API, GLFW_OPENGL_ES_API);
#elif defined(__APPLE__)
  // GL 3.2 + GLSL 150
  const char *glsl_version = "#version 150";
  glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3);
  glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 2);
  glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE); // 3.2+ only
  glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GL_TRUE);           // Required on Mac
#else
  // GL 3.0 + GLSL 130
  const char *glsl_version = "#version 130";
  glfwWindowHint(GLFW_CONTEXT_VERSION_MAJOR, 3);
  glfwWindowHint(GLFW_CONTEXT_VERSION_MINOR, 0);
  // glfwWindowHint(GLFW_OPENGL_PROFILE, GLFW_OPENGL_CORE_PROFILE);  // 3.2+
  // only glfwWindowHint(GLFW_OPENGL_FORWARD_COMPAT, GL_TRUE); // 3.0+ only
#endif

  // Create window with graphics context
  GLFWwindow *window = glfwCreateWindow(width, height, title, NULL, NULL);
  if (window == NULL)
    return NULL;
  glfwMakeContextCurrent(window);
  glfwSwapInterval(1); // Enable vsync

  // Setup Dear ImGui context
  igCreateContext(0);

  if (afterCreateContext != NULL) {
    afterCreateContext();
  }

  ImGuiIO *io = igGetIO();
  io->ConfigFlags |= ImGuiConfigFlags_NavEnableKeyboard; // Enable Keyboard Controls
  // io.ConfigFlags |= ImGuiConfigFlags_NavEnableGamepad;      // Enable Gamepad
  // Controls
  io->ConfigFlags |= ImGuiConfigFlags_DockingEnable;   // Enable Docking
  io->ConfigFlags |= ImGuiConfigFlags_ViewportsEnable; // Enable Multi-Viewport
                                                       // / Platform Windows
  // io.ConfigViewportsNoAutoMerge = true;
  // io.ConfigViewportsNoTaskBarIcon = true;

  io->IniFilename = "";

  // Setup Dear ImGui style
  igStyleColorsDark(0);
  // igStyleColorsLight();

  // When viewports are enabled we tweak WindowRounding/WindowBg so platform
  // windows can look identical to regular ones.
  ImGuiStyle *style = igGetStyle();
  if (io->ConfigFlags & ImGuiConfigFlags_ViewportsEnable) {
    style->WindowRounding = 0.0f;
    style->Colors[ImGuiCol_WindowBg].w = 1.0f;
  }

  // Setup Platform/Renderer backends
  ImGui_ImplGlfw_InitForOpenGL(window, true);
  ImGui_ImplOpenGL3_Init(glsl_version);

  // Install extra callback
  glfwSetWindowRefreshCallback(window, glfw_window_refresh_callback);
  glfwMakeContextCurrent(NULL);
  return window;
}

void igRefresh() { glfwPostEmptyEvent(); }

void glfw_render(GLFWwindow *window, VoidCallback renderLoop) {
  // Start the Dear ImGui frame
  ImGui_ImplOpenGL3_NewFrame();
  ImGui_ImplGlfw_NewFrame();
  igNewFrame();

  glfwSetWindowUserPointer(window, (void *)renderLoop);

  // Do ui stuff here
  if (renderLoop != NULL) {
    renderLoop();
  }

  // Rendering
  igRender();
  int display_w, display_h;
  glfwGetFramebufferSize(window, &display_w, &display_h);
  glViewport(0, 0, display_w, display_h);
  glClearColor(clear_color.x * clear_color.w, clear_color.y * clear_color.w, clear_color.z * clear_color.w,
               clear_color.w);
  glClear(GL_COLOR_BUFFER_BIT);
  ImGui_ImplOpenGL3_RenderDrawData(igGetDrawData());

  ImGuiIO *io = igGetIO();

  // Update and Render additional Platform Windows
  // (Platform functions may change the current OpenGL context, so we
  // save/restore it to make it easier to paste this code elsewhere.
  //  For this specific demo app we could also call
  //  glfwMakeContextCurrent(window) directly)
  if (io->ConfigFlags & ImGuiConfigFlags_ViewportsEnable) {
    GLFWwindow *backup_current_context = glfwGetCurrentContext();
    igUpdatePlatformWindows();
    igRenderPlatformWindowsDefault(NULL, NULL);
    glfwMakeContextCurrent(backup_current_context);
  }

  glfwSwapBuffers(window);
}

void igGLFWRunLoop(GLFWwindow *window, VoidCallback loop, VoidCallback beforeRender, VoidCallback afterRender,
               VoidCallback beforeDestroyContext) {
  glfwMakeContextCurrent(window);
  ImGuiIO *io = igGetIO();

  // Load Fonts
  // - If no fonts are loaded, dear imgui will use the default font. You can
  // also load multiple fonts and use igPushFont()/PopFont() to select them.
  // - AddFontFromFileTTF() will return the ImFont* so you can store it if you
  // need to select the font among multiple.
  // - If the file cannot be loaded, the function will return NULL. Please
  // handle those errors in your application (e.g. use an assertion, or display
  // an error and quit).
  // - The fonts will be rasterized at a given size (w/ oversampling) and stored
  // into a texture when calling ImFontAtlas::Build()/GetTexDataAsXXXX(), which
  // ImGui_ImplXXXX_NewFrame below will call.
  // - Read 'docs/FONTS.md' for more instructions and details.
  // - Remember that in C/C++ if you want to include a backslash \ in a string
  // literal you need to write a double backslash \\ !
  // io.Fonts->AddFontDefault();
  // io.Fonts->AddFontFromFileTTF("../../misc/fonts/Roboto-Medium.ttf", 16.0f);
  // io.Fonts->AddFontFromFileTTF("../../misc/fonts/Cousine-Regular.ttf", 15.0f);
  // io.Fonts->AddFontFromFileTTF("../../misc/fonts/DroidSans.ttf", 16.0f);
  // io.Fonts->AddFontFromFileTTF("../../misc/fonts/ProggyTiny.ttf", 10.0f);
  // ImFont* font =
  // io.Fonts->AddFontFromFileTTF("c:\\Windows\\Fonts\\ArialUni.ttf", 18.0f,
  // NULL, io.Fonts->GetGlyphRangesJapanese()); IM_ASSERT(font != NULL);

  // Main loop
  double lasttime = glfwGetTime();
  while (!glfwWindowShouldClose(window)) {
    if (beforeRender != NULL) {
      beforeRender();
    }

    glfw_render(window, loop);

    while (glfwGetTime() < lasttime + 1.0 / glfw_target_fps) {
      // do nothing here
    }
    lasttime += 1.0 / glfw_target_fps;

    if (extra_frame_count > 0) {
      extra_frame_count--;
    } else {
      glfwWaitEvents();
      extra_frame_count = MAX_EXTRA_FRAME_COUNT;
    }

    glfwPollEvents();

    if (afterRender != NULL) {
      afterRender();
    }
  }

  // Cleanup
  ImGui_ImplOpenGL3_Shutdown();
  ImGui_ImplGlfw_Shutdown();

  if (beforeDestroyContext != NULL) {
    beforeDestroyContext();
  }

  igDestroyContext(0);

  glfwDestroyWindow(window);
  glfwTerminate();

  return;
}

ImTextureID igCreateTexture(unsigned char *pixels, int width, int height) {
  GLint last_texture;
  GLuint texId;

  glGetIntegerv(GL_TEXTURE_BINDING_2D, &last_texture);
  glGenTextures(1, &texId);
  glBindTexture(GL_TEXTURE_2D, texId);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MIN_FILTER, GL_LINEAR);
  glTexParameteri(GL_TEXTURE_2D, GL_TEXTURE_MAG_FILTER, GL_LINEAR);
  glTexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, width, height, 0, GL_RGBA, GL_UNSIGNED_BYTE, pixels);

  // Restore state
  glBindTexture(GL_TEXTURE_2D, last_texture);

  return ImTextureID((intptr_t(texId)));
}

void igDeleteTexture(ImTextureID id) {
  glBindTexture(GL_TEXTURE_2D, 0);
  glDeleteTextures(1, (GLuint *)(&id));
}

void igGLFWWindow_GetDisplaySize(GLFWwindow *window, int *width, int *height) {
  glfwGetWindowSize(window, width, height);
}

void igGLFWWindow_GetContentScale(GLFWwindow *window, float *width, float *height) {
  glfwGetWindowContentScale(window, width, height);
}

void igGLFWWindow_SetWindowPos(GLFWwindow *window, int x, int y) { glfwSetWindowPos(window, x, y); }

void igGLFWWindow_GetWindowPos(GLFWwindow *window, int *x, int *y) { glfwGetWindowPos(window, x, y); }

void igGLFWWindow_SetShouldClose(GLFWwindow *window, int value) { glfwSetWindowShouldClose(window, value); }

void igGLFWWindow_SetDropCallbackCB(GLFWwindow *wnd) { glfwSetDropCallback(wnd, (GLFWdropfun)dropCallback); }
void igGLFWWindow_SetKeyCallback(GLFWwindow *wnd) { glfwSetKeyCallback(wnd, (GLFWkeyfun)keyCallback); }
void igGLFWWindow_SetSizeCallback(GLFWwindow *wnd) { glfwSetWindowSizeCallback(wnd, (GLFWwindowsizefun)sizeCallback); }

void igGLFWWindow_SetCloseCallback(GLFWwindow *window) {
  glfwSetWindowCloseCallback(window, (GLFWwindowclosefun)closeCallback);
}

void igGLFWWindow_SetSize(GLFWwindow *window, int width, int height) { glfwSetWindowSize(window, width, height); }

void igGLFWWindow_SetTitle(GLFWwindow *window, const char *title) { glfwSetWindowTitle(window, title); }

void igGLFWWindow_SetSizeLimits(GLFWwindow *window, int minWidth, int minHeight, int maxWidth, int maxHeight) {
  glfwSetWindowSizeLimits(window, minWidth, minHeight, maxWidth, maxHeight);
}

void igWindowHint(int hint, int value) { glfwWindowHint(hint, value); }

void igGLFWWindow_SetIcon(GLFWwindow *window, int count, CImage *images) {
    GLFWimage* glfwImages = (GLFWimage*)malloc(count * sizeof(GLFWimage));
    for (int i = 0; i < count; i++) {
        glfwImages[i].width = images[i].width;
        glfwImages[i].height = images[i].height;
        glfwImages[i].pixels = images[i].pixels;
    }

    glfwSetWindowIcon(window, count, glfwImages);
}

void iggImplGlfw_KeyCallback(GLFWwindow* w, int k,int s,int a,int m) { ImGui_ImplGlfw_KeyCallback(w,k,s,a,m); }

#endif

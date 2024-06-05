#pragma once

#include "cimgui_wrapper.h"
#include "thirdparty/SDL/include/SDL.h" // Will drag system OpenGL headers
#include "extra_types.h"

#ifdef __cplusplus
extern "C" {
#endif

/*
typedef int GLFWKey;
enum GLFWKey_ {
   GLFWKeySpace = GLFW_KEY_SPACE,
   GLFWKeyApostrophe = GLFW_KEY_APOSTROPHE,
   GLFWKeyComma = GLFW_KEY_COMMA,
   GLFWKeyMinus = GLFW_KEY_MINUS,
   GLFWKeyPeriod = GLFW_KEY_PERIOD,
   GLFWKeySlash = GLFW_KEY_SLASH,
   GLFWKey0 = GLFW_KEY_0,
   GLFWKey1 = GLFW_KEY_1,
   GLFWKey2 = GLFW_KEY_2,
   GLFWKey3 = GLFW_KEY_3,
   GLFWKey4 = GLFW_KEY_4,
   GLFWKey5 = GLFW_KEY_5,
   GLFWKey6 = GLFW_KEY_6,
   GLFWKey7 = GLFW_KEY_7,
   GLFWKey8 = GLFW_KEY_8,
   GLFWKey9 = GLFW_KEY_9,
   GLFWKeySemicolon = GLFW_KEY_SEMICOLON,
   GLFWKeyEqual = GLFW_KEY_EQUAL,
   GLFWKeyA = GLFW_KEY_A,
   GLFWKeyB = GLFW_KEY_B,
   GLFWKeyC = GLFW_KEY_C,
   GLFWKeyD = GLFW_KEY_D,
   GLFWKeyE = GLFW_KEY_E,
   GLFWKeyF = GLFW_KEY_F,
   GLFWKeyG = GLFW_KEY_G,
   GLFWKeyH = GLFW_KEY_H,
   GLFWKeyI = GLFW_KEY_I,
   GLFWKeyJ = GLFW_KEY_J,
   GLFWKeyK = GLFW_KEY_K,
   GLFWKeyL = GLFW_KEY_L,
   GLFWKeyM = GLFW_KEY_M,
   GLFWKeyN = GLFW_KEY_N,
   GLFWKeyO = GLFW_KEY_O,
   GLFWKeyP = GLFW_KEY_P,
   GLFWKeyQ = GLFW_KEY_Q,
   GLFWKeyR = GLFW_KEY_R,
   GLFWKeyS = GLFW_KEY_S,
   GLFWKeyT = GLFW_KEY_T,
   GLFWKeyU = GLFW_KEY_U,
   GLFWKeyV = GLFW_KEY_V,
   GLFWKeyW = GLFW_KEY_W,
   GLFWKeyX = GLFW_KEY_X,
   GLFWKeyY = GLFW_KEY_Y,
   GLFWKeyZ = GLFW_KEY_Z,
   GLFWKeyLeftBracket = GLFW_KEY_LEFT_BRACKET,
   GLFWKeyBackslash = GLFW_KEY_BACKSLASH,
   GLFWKeyRightBracket = GLFW_KEY_RIGHT_BRACKET,
   GLFWKeyGraveAccent = GLFW_KEY_GRAVE_ACCENT,
   GLFWKeyWorld1 = GLFW_KEY_WORLD_1,
   GLFWKeyWorld2 = GLFW_KEY_WORLD_2,

   // Function keys
   GLFWKeyEscape = GLFW_KEY_ESCAPE,
   GLFWKeyEnter = GLFW_KEY_ENTER,
   GLFWKeyTab = GLFW_KEY_TAB,
   GLFWKeyBackspace = GLFW_KEY_BACKSPACE,
   GLFWKeyInsert = GLFW_KEY_INSERT,
   GLFWKeyDelete = GLFW_KEY_DELETE,
   GLFWKeyRight = GLFW_KEY_RIGHT,
   GLFWKeyLeft = GLFW_KEY_LEFT,
   GLFWKeyDown = GLFW_KEY_DOWN,
   GLFWKeyUp = GLFW_KEY_UP,
   GLFWKeyPageUp = GLFW_KEY_PAGE_UP,
   GLFWKeyPageDown = GLFW_KEY_PAGE_DOWN,
   GLFWKeyHome = GLFW_KEY_HOME,
   GLFWKeyEnd = GLFW_KEY_END,
   GLFWKeyCapsLock = GLFW_KEY_CAPS_LOCK,
   GLFWKeyScrollLock = GLFW_KEY_SCROLL_LOCK,
   GLFWKeyNumLock = GLFW_KEY_NUM_LOCK,
   GLFWKeyPrintScreen = GLFW_KEY_PRINT_SCREEN,
   GLFWKeyPause = GLFW_KEY_PAUSE,
   GLFWKeyF1 = GLFW_KEY_F1,
   GLFWKeyF2 = GLFW_KEY_F2,
   GLFWKeyF3 = GLFW_KEY_F3,
   GLFWKeyF4 = GLFW_KEY_F4,
   GLFWKeyF5 = GLFW_KEY_F5,
   GLFWKeyF6 = GLFW_KEY_F6,
   GLFWKeyF7 = GLFW_KEY_F7,
   GLFWKeyF8 = GLFW_KEY_F8,
   GLFWKeyF9 = GLFW_KEY_F9,
   GLFWKeyF10 = GLFW_KEY_F10,
   GLFWKeyF11 = GLFW_KEY_F11,
   GLFWKeyF12 = GLFW_KEY_F12,
   GLFWKeyF13 = GLFW_KEY_F13,
   GLFWKeyF14 = GLFW_KEY_F14,
   GLFWKeyF15 = GLFW_KEY_F15,
   GLFWKeyF16 = GLFW_KEY_F16,
   GLFWKeyF17 = GLFW_KEY_F17,
   GLFWKeyF18 = GLFW_KEY_F18,
   GLFWKeyF19 = GLFW_KEY_F19,
   GLFWKeyF20 = GLFW_KEY_F20,
   GLFWKeyF21 = GLFW_KEY_F21,
   GLFWKeyF22 = GLFW_KEY_F22,
   GLFWKeyF23 = GLFW_KEY_F23,
   GLFWKeyF24 = GLFW_KEY_F24,
   GLFWKeyF25 = GLFW_KEY_F25,
   GLFWKeyKp0 = GLFW_KEY_KP_0,
   GLFWKeyKp1 = GLFW_KEY_KP_1,
   GLFWKeyKp2 = GLFW_KEY_KP_2,
   GLFWKeyKp3 = GLFW_KEY_KP_3,
   GLFWKeyKp4 = GLFW_KEY_KP_4,
   GLFWKeyKp5 = GLFW_KEY_KP_5,
   GLFWKeyKp6 = GLFW_KEY_KP_6,
   GLFWKeyKp7 = GLFW_KEY_KP_7,
   GLFWKeyKp8 = GLFW_KEY_KP_8,
   GLFWKeyKp9 = GLFW_KEY_KP_9,
   GLFWKeyKpDecimal = GLFW_KEY_KP_DECIMAL,
   GLFWKeyKpDivide = GLFW_KEY_KP_DIVIDE,
   GLFWKeyKpMultiply = GLFW_KEY_KP_MULTIPLY,
   GLFWKeyKpSubtract = GLFW_KEY_KP_SUBTRACT,
   GLFWKeyKpAdd = GLFW_KEY_KP_ADD,
   GLFWKeyKpEnter = GLFW_KEY_KP_ENTER,
   GLFWKeyKpEqual = GLFW_KEY_KP_EQUAL,
   GLFWKeyLeftShift = GLFW_KEY_LEFT_SHIFT,
   GLFWKeyLeftControl = GLFW_KEY_LEFT_CONTROL,
   GLFWKeyLeftAlt = GLFW_KEY_LEFT_ALT,
   GLFWKeyLeftSuper = GLFW_KEY_LEFT_SUPER,
   GLFWKeyRightShift = GLFW_KEY_RIGHT_SHIFT,
   GLFWKeyRightControl = GLFW_KEY_RIGHT_CONTROL,
   GLFWKeyRightAlt = GLFW_KEY_RIGHT_ALT,
   GLFWKeyRightSuper = GLFW_KEY_RIGHT_SUPER,
   GLFWKeyMenu = GLFW_KEY_MENU,
};

typedef int GLFWModifierKey;
enum GLFWModifierKey_ {
   GLFWModShift = GLFW_MOD_SHIFT,
   GLFWModControl = GLFW_MOD_CONTROL,
   GLFWModAlt = GLFW_MOD_ALT,
   GLFWModSuper = GLFW_MOD_SUPER,
   GLFWModCapsLock = GLFW_MOD_CAPS_LOCK,
   GLFWModNumLock = GLFW_MOD_NUM_LOCK,
};
*/

/*
enum GLFWWindowFlags_ {
  GLFWWindowNone = 0,
  GLFWWindowResizable = GLFW_RESIZABLE,
  GLFWWindowMaximized = GLFW_MAXIMIZED,
  GLFWWindowTransparentFramebuffer = GLFW_TRANSPARENT_FRAMEBUFFER,
  GLFWWindowDecorated= GLFW_DECORATED,
  GLFWWindowVisible = GLFW_VISIBLE,
  GLFWWindowFloating = GLFW_FLOATING,
  GLFWWindowFocused = GLFW_FOCUSED,
  GLFWWindowIconified = GLFW_ICONIFIED,
  GLFWWindowAutoIconify = GLFW_AUTO_ICONIFY,
};
*/

typedef struct CImage {
int width;
int height;
unsigned char* pixels;
}CImage ;

typedef struct SDL_Window SDL_Window;
//typedef struct GLFWmonitor GLFWmonitor;
//typedef struct GLFWimage GLFWimage;
struct SDL_Window;
//struct GLFWmonitor;
//struct GLFWimage;

extern void igSetBgColor(ImVec4 color);
extern void igSetTargetFPS(unsigned int fps);
extern int igInitSDL();
extern SDL_Window *igCreateSDLWindow(const char *title, int width, int height,
                                      VoidCallback afterCreateContext);
extern void igSDLRunLoop(SDL_Window *window, VoidCallback loop, VoidCallback beforeRender, VoidCallback afterRender,
                      VoidCallback beforeDestroyContext);
extern void igSDLWindow_GetDisplaySize(SDL_Window *window, int *width, int *height);
extern void igSDLWindow_GetContentScale(SDL_Window *window, float *width, float *height);
extern void igSDLWindow_SetWindowPos(SDL_Window *window, int x, int y);
extern void igSDLWindow_GetWindowPos(SDL_Window *window, int *x, int *y);
//extern void igSDLWindow_SetShouldClose(SDL_Window *window, int value);
//extern void igSDLWindow_SetDropCallbackCB(SDL_Window *window);
extern void igSDLWindow_SetSize(SDL_Window *window, int width, int height);
extern void igSDLWindow_SetTitle(SDL_Window *window, const char *title);
extern void igSDLWindow_SetSizeLimits(SDL_Window *window, int minWidth, int minHeight, int maxWidth, int maxHeight);
//extern void igSDLWindow_SetCloseCallback(SDL_Window *window);
//extern void igSDLWindow_SetKeyCallback(SDL_Window *window);
//extern void igSDLWindow_SetSizeCallback(SDL_Window *window);
//extern void igSDLWindow_SetIcon(SDL_Window *window, int count, CImage* images);
extern void igRefresh();
extern ImTextureID igCreateTexture(unsigned char *pixels, int width, int height);
extern void igDeleteTexture(ImTextureID id);
extern void igSDLWindowHint(SDL_WindowFlags hint, int value);

//extern void dropCallback(int, char **);
//extern void closeCallback(SDL_Window *window);
//extern void keyCallback(int key, int scancode, int action, int mods);
//extern void sizeCallback(int width, int height);

#ifdef __cplusplus
}
#endif

#pragma once

#include "../../imgui/extra_types.h"
#include <stdbool.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#if !defined(IMGUI_VERSION)
typedef struct ImVec4 {
    float x;
    float y;
    float z;
    float w;
} ImVec4;

typedef uint64_t ImTextureID;
#endif

typedef struct MetalBackendContext MetalBackendContext;

typedef int MetalWindowFlag;
enum MetalWindowFlag_ {
    MetalWindowFlagResizable = 0,
    MetalWindowFlagTitled = 1,
    MetalWindowFlagClosable = 2,
    MetalWindowFlagMiniaturizable = 3,
    MetalWindowFlagFullSizeContentView = 4,
    MetalWindowFlagTransparent = 5,
};

int igInitMetal(void);

MetalBackendContext* igCreateMetalWindow(const char* title, int width, int height);

void igMetalRunLoop(MetalBackendContext* ctx, VoidCallback loop, VoidCallback beforeRender, VoidCallback afterRender,
                    VoidCallback beforeDestroyContext);

void igMetalSetBgColor(MetalBackendContext* ctx, ImVec4 color);
void igMetalSetTargetFPS(MetalBackendContext* ctx, unsigned int fps);
void igMetalRefresh(MetalBackendContext* ctx);

void igMetalWindow_GetDisplaySize(MetalBackendContext* ctx, int* width, int* height);
void igMetalWindow_GetContentScale(MetalBackendContext* ctx, float* width, float* height);
void igMetalWindow_SetWindowPos(MetalBackendContext* ctx, int x, int y);
void igMetalWindow_GetWindowPos(MetalBackendContext* ctx, int* x, int* y);
void igMetalWindow_SetSize(MetalBackendContext* ctx, int width, int height);
void igMetalWindow_SetSizeLimits(MetalBackendContext* ctx, int minWidth, int minHeight, int maxWidth, int maxHeight);
void igMetalWindow_SetTitle(MetalBackendContext* ctx, const char* title);
void igMetalWindow_SetShouldClose(MetalBackendContext* ctx, bool value);

void igMetalWindow_SetCloseCallback(MetalBackendContext* ctx);
void igMetalWindow_SetDropCallback(MetalBackendContext* ctx);
void igMetalWindow_SetSizeCallback(MetalBackendContext* ctx);
void igMetalWindow_SetFlag(MetalBackendContext* ctx, MetalWindowFlag flag, int value);

ImTextureID igMetalCreateTexture(MetalBackendContext* ctx, unsigned char* pixels, int width, int height);
void igMetalDeleteTexture(ImTextureID texture_id);

#ifdef __cplusplus
}
#endif

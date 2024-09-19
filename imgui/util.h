#pragma once

#include "cimgui_wrapper.h"
#include <stddef.h>

#ifdef __cplusplus
extern "C" {
#endif

extern void wrap_GetIndexBufferLayout(size_t *entrySize);

extern void wrap_GetVertexBufferLayout(size_t *entrySize, size_t *posOffset, size_t *uvOffset, size_t *colOffset);

extern ImDrawList *wrap_DrawData_GetDrawListAt(ImDrawData *self, int idx);

extern ImDrawCmd *wrap_DrawList_GetDrawCmdAt(ImDrawList *self, int idx);

extern void wrap_DrawCmd_CallUserCallback(ImDrawList *list, ImDrawCmd *cmd);

extern void wrap_ImGuiIO_SetMouseButtonDown(ImGuiIO *self, int btnIndex, bool isDown);

extern ImVector_ImWchar *wrap_NewGlyphRange();
extern void wrap_DestroyGlyphRange(ImVector_ImWchar *range);
extern ImWchar *wrap_GlyphRange_GetData(ImVector_ImWchar *range);

extern int wrap_ImFontAtlas_GetFontCount(ImFontAtlas *self);

#ifdef __cplusplus
}
#endif

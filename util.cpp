#include "util.h"
#include "cimgui/cimgui.h"
#include "cimgui_wrapper.h"

#define IM_OFFSETOF(_TYPE, _MEMBER) offsetof(_TYPE, _MEMBER) // Offset of _MEMBER within _TYPE. Standardized as offsetof() in C++11

void GetIndexBufferLayout(size_t *entrySize) { *entrySize = sizeof(ImDrawIdx); }

void GetVertexBufferLayout(size_t *entrySize, size_t *posOffset, size_t *uvOffset, size_t *colOffset) {
  *entrySize = sizeof(ImDrawVert);
  *posOffset = IM_OFFSETOF(ImDrawVert, pos);
  *uvOffset = IM_OFFSETOF(ImDrawVert, uv);
  *colOffset = IM_OFFSETOF(ImDrawVert, col);
}

ImDrawList *DrawData_GetDrawListAt(ImDrawData *self, int idx) {
  if (idx >= 0 && idx < self->CmdListsCount)
    return self->CmdLists[idx];

  return NULL;
}

ImDrawCmd *DrawList_GetDrawCmdAt(ImDrawList *self, int idx) {
  if (idx >= 0 && idx < self->CmdBuffer.Size)
    return &(self->CmdBuffer.Data[idx]);

  return NULL;
}

void DrawCmd_CallUserCallback(ImDrawList *list, ImDrawCmd *cmd) { cmd->UserCallback(list, cmd); }

void ImGuiIO_SetMouseButtonDown(ImGuiIO *self, int btnIndex, bool isDown) { self->MouseDown[btnIndex] = isDown; }

ImVector_ImWchar *NewGlyphRange() { return ImVector_ImWchar_create(); }

void DestroyGlyphRange(ImVector_ImWchar *range) { ImVector_ImWchar_destroy(range); }

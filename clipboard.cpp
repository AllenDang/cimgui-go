#include "clipboard.h"
#include "cimgui_wrapper.h"

extern "C" char const *ImGuiIO_GetClipboardText(ImGuiIO *io);
extern "C" void ImGuiIO_SetClipboardText(ImGuiIO *io, const char *text);

static void ImGuiIO_SetClipboardTextWrapper(void *io, char const *text) { ImGuiIO_SetClipboardText((ImGuiIO *)io, text); }

static char const *ImGuiIO_GetClipboardTextWrapper(void *io) { return ImGuiIO_GetClipboardText((ImGuiIO *)io); }

void ImGuiIO_RegisterClipboardFunctions(ImGuiIO *io) {
  io->ClipboardUserData = io;
  io->GetClipboardTextFn = ImGuiIO_GetClipboardTextWrapper;
  io->SetClipboardTextFn = ImGuiIO_SetClipboardTextWrapper;
}

void ImGuiIO_ClearClipboardFunctions(ImGuiIO *io) {
  io->GetClipboardTextFn = NULL;
  io->SetClipboardTextFn = NULL;
  io->ClipboardUserData = NULL;
}

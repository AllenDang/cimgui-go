#pragma once

#include "cimgui_wrapper.h"

#ifdef __cplusplus
extern "C" {
#endif

extern void ImGuiIO_RegisterClipboardFunctions(ImGuiIO *io);
extern void ImGuiIO_ClearClipboardFunctions(ImGuiIO *io);

#ifdef __cplusplus
}
#endif

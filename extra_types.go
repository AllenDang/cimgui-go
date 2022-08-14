package cimgui

// #include "cimgui_wrapper.h"
import "C"

type ImWchar C.uint
type ImGuiID C.ImGuiID
type ImTextureID C.ImTextureID
type ImDrawIdx C.ImDrawIdx

type ImGuiInputTextCallback func(*ImGuiInputTextCallbackData) int
type ImGuiSizeCallback func(*ImGuiSizeCallbackData)

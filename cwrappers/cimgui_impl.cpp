#include "./imgui/imgui.h"
#ifdef IMGUI_ENABLE_FREETYPE
#include "./imgui/misc/freetype/imgui_freetype.h"
#endif
#include "./imgui/imgui_internal.h"
#include "cimgui.h"


#ifdef CIMGUI_USE_GLFW
#include "imgui_impl_glfw.h"
#endif

#ifdef CIMGUI_USE_OPENGL3
#include "imgui_impl_opengl3.h"
#endif

#ifdef CIMGUI_USE_OPENGL2
#include "imgui_impl_opengl2.h"
#endif

#ifdef CIMGUI_USE_SDL2
#include "imgui_impl_sdl2.h"
#endif

#include "cimgui_impl.h"

#ifdef CIMGUI_USE_VULKAN

CIMGUI_API ImGui_ImplVulkanH_Window* ImGui_ImplVulkanH_Window_ImGui_ImplVulkanH_Window()
{
	return IM_NEW(ImGui_ImplVulkanH_Window)();
}
CIMGUI_API void ImGui_ImplVulkanH_Window_Construct(ImGui_ImplVulkanH_Window* self)
{
	IM_PLACEMENT_NEW(self) ImGui_ImplVulkanH_Window();
}

#endif

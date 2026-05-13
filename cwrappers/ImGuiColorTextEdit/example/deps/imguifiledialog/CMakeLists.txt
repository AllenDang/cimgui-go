#	TextEditor - A syntax highlighting text editor for ImGui
#	Copyright (c) 2024-2026 Johan A. Goossens. All rights reserved.
#
#	This work is licensed under the terms of the MIT license.
#	For a copy, see <https://opensource.org/licenses/MIT>.

FetchContent_Populate(imguifd QUIET GIT_REPOSITORY "https://github.com/goossens/ImGuiFileDialog")

set(IMGUIFD_SOURCES
	${imguifd_SOURCE_DIR}/ImGuiFileDialog.h
	${imguifd_SOURCE_DIR}/ImGuiFileDialogConfig.h
	${imguifd_SOURCE_DIR}/ImGuiFileDialog.cpp)

add_library(imguifd STATIC ${IMGUIFD_SOURCES})

target_compile_definitions(imguifd PUBLIC CUSTOM_IMGUIFILEDIALOG_CONFIG="IgfdConfig.h")
target_include_directories(imguifd PUBLIC ${imguifd_SOURCE_DIR} ${CMAKE_CURRENT_LIST_DIR})
target_link_libraries(imguifd imgui)

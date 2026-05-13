#	TextEditor - A syntax highlighting text editor for ImGui
#	Copyright (c) 2024-2026 Johan A. Goossens. All rights reserved.
#
#	This work is licensed under the terms of the MIT license.
#	For a copy, see <https://opensource.org/licenses/MIT>.

set(BUILD_SHARED_LIBS OFF CACHE BOOL "" FORCE)
set(SDL_TEST_LIBRARY OFF CACHE BOOL "" FORCE)

FetchContent_Populate(sdl QUIET GIT_REPOSITORY "https://github.com/libsdl-org/SDL" GIT_TAG "release-3.4.4")
add_subdirectory(${sdl_SOURCE_DIR} ${sdl_BINARY_DIR} EXCLUDE_FROM_ALL)

if(CMAKE_GENERATOR STREQUAL "Xcode")
	set_target_properties(SDL3-static PROPERTIES XCODE_ATTRIBUTE_GCC_WARN_ABOUT_DEPRECATED_FUNCTIONS "NO")
	set_target_properties(SDL3-static PROPERTIES XCODE_ATTRIBUTE_GCC_WARN_64_TO_32_BIT_CONVERSION "NO")
	set_target_properties(SDL3-static PROPERTIES DISABLE_PRECOMPILE_HEADERS ON)
endif()

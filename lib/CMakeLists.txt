cmake_minimum_required(VERSION 3.5)

project(cimgui)

set (CMAKE_CXX_STANDARD 14)

if (EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/../cwrappers/imgui/imgui_tables.cpp)
	set(TABLES_SOURCE "../cwrappers/imgui/imgui_tables.cpp")
else()
	set(TABLES_SOURCE "")
endif()


#general settings
file(GLOB IMGUI_SOURCES
    ../cwrappers/cimgui.cpp
    ../cwrappers/imgui/imgui.cpp
    ../cwrappers/imgui/imgui_draw.cpp
    ../cwrappers/imgui/imgui_demo.cpp
    ../cwrappers/imgui/imgui_widgets.cpp
    ../cwrappers/implot/implot.cpp
    ../cwrappers/implot/implot_items.cpp
    ../cwrappers/implot/implot_demo.cpp
    ../cwrappers/cimplot.cpp
    ../cwrappers/imnodes/imnodes.cpp
    ../cwrappers/cimnodes.cpp
    ../cwrappers/cimmarkdown.cpp
    ../cwrappers/imgui_markdown/imgui_markdown.h
    ../cwrappers/ImGuizmo/ImGuizmo.cpp
    ../cwrappers/cimguizmo.cpp
    ../cwrappers/ImGuiColorTextEdit/TextEditor.cpp
    ../cwrappers/ImGuiColorTextEdit/LanguageDefinitions.cpp
    ../cwrappers/ImGuiColorTextEdit/UnitTests.cpp
    ../cwrappers/ImGuiColorTextEdit/ImGuiDebugPanel.cpp
    ../cwrappers/cimCTE.cpp
	${TABLES_SOURCE}
)

set(CMAKE_BUILD_TYPE "Release")
set(IMGUI_STATIC "on" CACHE STRING "Build as a static library")
set(IMGUI_FREETYPE "no" CACHE STRING "Build with freetype library")
set(IMGUI_LIBRARIES )
add_definitions(-DIMNODES_NAMESPACE=imnodes)
add_definitions(-DIMGUI_MARKDOWN_NAMESPACE=ImMarkdown)

add_definitions(-DIMGUI_USE_WCHAR32)

#glfw
set(GLFW3_PREFIX ${CMAKE_CURRENT_SOURCE_DIR}/../thirdparty/glfw)
set(GLFW3_INCLUDE_DIR ${GLFW3_PREFIX}/include)
list(APPEND IMGUI_SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/../cwrappers/imgui/backends/imgui_impl_glfw.cpp)
list(APPEND IMGUI_LIBRARIES glfw)

#SDL
set(SDL2_PREFIX ${CMAKE_CURRENT_SOURCE_DIR}/../thirdparty/SDL)
set(SDL2_INCLUDE_DIR ${SDL2_PREFIX}/include)
list(APPEND IMGUI_SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/../cwrappers/imgui/backends/imgui_impl_sdl2.cpp)
list(APPEND IMGUI_LIBRARIES sdl2)

#opengl3
list(APPEND IMGUI_SOURCES ${CMAKE_CURRENT_SOURCE_DIR}/../cwrappers/imgui/backends/imgui_impl_opengl3.cpp)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/../cwrappers/imgui/examples/libs/gl3w)
if(WIN32)
  list(APPEND IMGUI_LIBRARIES opengl32)
else(WIN32)#Unix
  list(APPEND IMGUI_LIBRARIES GL)
endif(WIN32)

if(IMGUI_FREETYPE)
	FIND_PACKAGE(freetype REQUIRED PATHS ${FREETYPE_PATH})
	list(APPEND IMGUI_LIBRARIES freetype)
	list(APPEND IMGUI_SOURCES imgui/misc/freetype/imgui_freetype.cpp)
	add_definitions("-DCIMGUI_FREETYPE=1")
endif(IMGUI_FREETYPE)

# add library and link
if (IMGUI_STATIC)
  add_library(cimgui STATIC ${IMGUI_SOURCES})
else (IMGUI_STATIC)
  add_library(cimgui SHARED ${IMGUI_SOURCES})
endif (IMGUI_STATIC)

target_compile_definitions(cimgui PUBLIC IMGUI_DISABLE_OBSOLETE_FUNCTIONS=1)
if (WIN32)
  list(APPEND IMGUI_LIBRARIES imm32)
endif (WIN32)

target_compile_definitions(cimgui PUBLIC IMGUI_IMPL_API=extern\t\"C\"\t)

target_include_directories(cimgui PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/../cwrappers/)
target_include_directories(cimgui PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/../cwrappers/imgui)
target_include_directories(cimgui PUBLIC ${GLFW3_INCLUDE_DIR})
target_include_directories(cimgui PUBLIC ${SDL2_INCLUDE_DIR})
set_target_properties(cimgui PROPERTIES PREFIX "")
target_link_libraries(cimgui ${IMGUI_LIBRARIES})
# this is for regex used by ImGuiColorTextEdit
target_include_directories(cimgui PRIVATE ../cwrappers/ImGuiColorTextEdit/vendor1/regex/include)

#install
install(TARGETS cimgui
    RUNTIME DESTINATION  .
    LIBRARY DESTINATION  .
    ARCHIVE DESTINATION  .
)

#test
set(CIMGUI_TEST "no" CACHE STRING "Enable compilation of a test unit based on imgui null")

if (CIMGUI_TEST)
  add_subdirectory(test)
endif ()

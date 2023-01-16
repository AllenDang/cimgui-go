cmake_minimum_required(VERSION 3.2)

project(cimgui)

set (CMAKE_CXX_STANDARD 11)

if (EXISTS ${CMAKE_CURRENT_SOURCE_DIR}/imgui/imgui_tables.cpp)
	set(TABLES_SOURCE "imgui/imgui_tables.cpp")
else()
	set(TABLES_SOURCE "")
endif()


#general settings
file(GLOB IMGUI_SOURCES
    cimgui.cpp
    imgui/imgui.cpp
    imgui/imgui_draw.cpp
    imgui/imgui_demo.cpp
    imgui/imgui_widgets.cpp
    ../cimplot/implot/implot.cpp
    ../cimplot/implot/implot_items.cpp
    ../cimplot/implot/implot_demo.cpp
    ../cimplot/cimplot.cpp
	${TABLES_SOURCE}
)

set(IMGUI_STATIC "no" CACHE STRING "Build as a static library")
set(IMGUI_FREETYPE "no" CACHE STRING "Build with freetype library")
set(IMGUI_LIBRARIES )

#glfw
set(GLFW3_PREFIX ${CMAKE_CURRENT_SOURCE_DIR}/../thirdparty/glfw)
set(GLFW3_INCLUDE_DIR ${GLFW3_PREFIX}/include)
list(APPEND IMGUI_SOURCES imgui/backends/imgui_impl_glfw.cpp)
list(APPEND IMGUI_LIBRARIES glfw)

#opengl3
list(APPEND IMGUI_SOURCES imgui/backends/imgui_impl_opengl3.cpp)
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/cimgui/imgui/examples/libs/gl3w)
if(WIN32)
  list(APPEND IMGUI_LIBRARIES opengl32)
else(WIN32)#Unix
  list(APPEND IMGUI_LIBRARIES GL)
endif(WIN32)

if(IMGUI_FREETYPE)
	FIND_PACKAGE(Freetype REQUIRED)
    INCLUDE_DIRECTORIES(${FREETYPE_INCLUDE_DIR_freetype2})
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

target_include_directories(cimgui PUBLIC ${CMAKE_CURRENT_SOURCE_DIR})
target_include_directories(cimgui PUBLIC ${CMAKE_CURRENT_SOURCE_DIR}/imgui)
target_include_directories(cimgui PUBLIC ${GLFW3_INCLUDE_DIR})
set_target_properties(cimgui PROPERTIES PREFIX "")
target_link_libraries(cimgui ${IMGUI_LIBRARIES})

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

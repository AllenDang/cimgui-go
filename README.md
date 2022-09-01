# cimgui-go

This project aims to generate go wrapper for Dear ImGui.

It comes with a default backend with GLFW 3.3 and OpenGL 3.2.

It works on macOS(arm64/x86) and windows(x64) now, idealy linux should works but I don't have a linux machine to test it. Check out `examples`, cd in and `go run .`.

## Current solution is:
1. Use cimgui's lua generator to generate function and struct definition as json.
2. Generate proper go code from the definition (via manual crafted go program `/cmd/codegen`).
3. Use the backend implementation from imgui (currently glfw and opengl3).
4. Use github workflow to compile cimgui and glfw to static lib and place them in /lib folder for further link. 

## Naming convention
For functions, 'Im/ImGui/ig' is trimmed.
'GetCursorPos' is renamed to 'GetDrawCursor', same with "SetCursor...".

## Function coverage
Currently most of the functions are generated, except memory related stuff (eg. memory allocator, storage management, etc...).
If you find any function is missing, report an issue.

## Generate binding
Install [just](https://just.systems/)

### Update imgui
1. Drop source code of imgui to `cimgui/imgui`.
2. Run `cd cimgui/generator; ./generator.sh`.
3. Run `just gencode_cimgui`.

### Update implot
1. Drop source code of implot to `cimplot/implot`.
2. Run `cd cimplot/generator; ./generator.sh`.
3. Run `just gencode_cimplot`.


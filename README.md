# cimgui-go [![GoDoc](https://pkg.go.dev/badge/github.com/AllenDang/cimgui-go?utm_source=godoc)](https://pkg.go.dev/mod/github.com/AllenDang/cimgui-go)

This project aims to generate go wrapper for Dear ImGui.

It comes with a default backend with GLFW 3.3 and OpenGL 3.2.

It works on macOS(arm64/x86), windows(x64), Arch Linux/KDE and Fedora Workstation 36, idealy other linux GUI should works but I don't have a linux machine to test it. Check out `examples`, cd in and `go run .`.

## Current solution is:
1. Use cimgui's lua generator to generate function and struct definition as json.
2. Generate proper go code from the definition (via manual crafted go program `/cmd/codegen`).
3. Use the backend implementation from imgui (currently glfw and opengl3).
4. Use github workflow to compile cimgui and glfw to static lib and place them in /lib folder for further link. 

## Naming convention

- For functions, 'Im/ImGui/ig' is trimmed.
- If function comes from `imgui_internal.h`, `Internal` prefix is added.
- Struct fields (if any exported) are prefixed with word `Field`

## Function coverage
Currently most of the functions are generated, except memory related stuff (eg. memory allocator, storage management, etc...).
If you find any function is missing, report an issue.

## Generate binding
Install [GNU make](https://www.gnu.org/software/make/manual/make.html) and run `make` to re-generate bunding.

## Update

To update to the latest version of dependencies, run `make update`.
After doing this, commit changes and navigate to GitHub.
In Actions tab, manually trigger workflows for each platform.

## How does it work?

- `cimgui/` directory holds C binding for C++ Dear ImGui libraries
- generator bases on `cimgui/{package_name}_templates` and generates all necessary GO/C code
- `libs/` contains pre-built shared libraries. `cimgui.go` includes and uses to decrease building time.

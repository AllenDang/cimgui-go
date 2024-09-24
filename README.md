[![Go Report Card](https://goreportcard.com/badge/github.com/AllenDang/cimgui-go)](https://goreportcard.com/report/github.com/AllenDang/cimgui-go)
![Build Status](https://github.com/AllenDang/cimgui-go/actions/workflows/build.yml/badge.svg)
[![GoDoc](https://pkg.go.dev/badge/github.com/AllenDang/cimgui-go?utm_source=godoc)](https://pkg.go.dev/mod/github.com/AllenDang/cimgui-go)

# cimgui-go 

This project aims to generate go wrapper for Dear ImGui.

It comes with a default backend with GLFW 3.3 and OpenGL 3.2.

It works on macOS(arm64/x86), windows(x64), Arch Linux/KDE and Fedora Workstation 36, idealy other linux GUI should works but I don't have a linux machine to test it. Check out `examples`, cd in and `go run .`.

## Current solution is:
1. Use cimgui's lua generator to generate function and struct definition as json.
2. Generate proper go code from the definition (via manual crafted go program `/cmd/codegen`).
3. Use the backend implementation from imgui.
4. Use github workflow to compile cimgui and glfw to static lib and place them in /lib folder for further link. 

## Supported Backends

In order to make it easy in use, cimgui-go implemented a few imgui backends.
To see more details about using a specific backend, take a look at the [examples](./examples).

We support the following backends:
- [GLFW](./examples/glfw). **Note:**: It is disabled by default, use `enable_cimgui_glfw` go build tag to enable it.
- [SDL2](./examples/sdl). **Note:**: It is disabled by default, use `enable_cimgui_sdl2` go build tag to enable it.
- [Ebitengine](./examples/ebiten) (`import "github.com/AllenDang/cimgui-go/ebitenbackend"`).

**Important**: Remember that various solution use different C libraries that can conflict with each other.
It is recommended to not enable e.g. GLFW and SDL backends at the same time as it may result in linker crash.

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
In Actions tab, manually trigger `Compile cimgui` workflows.

## How does it work?

- `cimgui/` directory holds C binding for C++ Dear ImGui libraries
- generator bases on `cimgui/{package_name}_templates` and generates all necessary GO/C code
- `libs/` contains pre-built shared libraries. `cimgui.go` includes and uses to decrease building time.

[![Go Report Card](https://goreportcard.com/badge/github.com/AllenDang/cimgui-go)](https://goreportcard.com/report/github.com/AllenDang/cimgui-go)
![Build Status](https://github.com/AllenDang/cimgui-go/actions/workflows/test.yaml/badge.svg)
![Linter Status](https://github.com/AllenDang/cimgui-go/actions/workflows/golangci-lint.yaml/badge.svg)
[![GoDoc](https://pkg.go.dev/badge/github.com/AllenDang/cimgui-go?utm_source=godoc)](https://pkg.go.dev/mod/github.com/AllenDang/cimgui-go)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#gui)

# cimgui-go 

This project aims to generate go wrapper for Dear ImGui.

It comes with [several default backends](supported-backends) implemented.

It works on macOS(arm64/x86), windows(x64), Arch Linux (Gnome/KDE) and Fedora Workstation 36+, idealy other linux GUI should works as well. Check out [`examples`](./examples): cd in and `go run .`.

## Setup

There are several dependencies you might need to install on your linux machine.
Take a look [here](https://github.com/allendang/giu#install)

## Current solution is:
1. Use [cimgui](https://github.com/cimgui/cimgui)'s lua generator to generate function and struct definition as json.
2. Generate proper go code from the definition ([via manual crafted go program](./cmd/codegen)).
3. Use [the backend implementation](#supported-backends) from imgui.
4. Use github workflow to compile cimgui, glfw and other C dependencies to static lib and place them in ./lib folder for further link. 

## Supported Backends

In order to make it easy in use, cimgui-go implemented a few imgui backends. All of them are placed in `backend/` sub-packages.
To see more details about usage of a specific backend, take a look at the [examples](./examples).

We support the following backends:
- [GLFW](./examples/glfw). (GLFW 3.3 + OpenGL)
- [SDL2](./examples/sdl). (SDL 2 + OpenGL)
- [Ebitengine](./examples/ebiten) (`import "github.com/AllenDang/cimgui-go/backend/ebitenbackend"`).

**Important**: Remember that various solution use different C libraries that can conflict with each other.
It is recommended to not enable e.g. GLFW and SDL backends at the same time as it may result in linker crash.

## Naming convention

- For functions, 'Im/ImGui/ig' is trimmed.
- `Get` prefix is also removed.
- If function comes from `imgui_internal.h`, `Internal` prefix is added.

## Pointers and Slices

Unfortunately, in C there is no way to ditinguish between a pointer and a slice.
We had to bring this unpleasantness to Go as well.
Our code defaults to pointers, but you can easily convert your slice to a pointer by simply `&(slice[0])`.
You can also use `imgui.SliceToPtr`.

## Callbacks

Please note that at the moment (November 2024) go (1.23) does not support passing annonymous functions to C via CGO.
We have it workarounded by pre-generating large amount of global functions and a pool.
For details see https://github.com/AllenDang/cimgui-go/issues/224 .
Just be aware of the limitation that you may run out of pre-generated pool and experience a crash.

## Function coverage
Currently most of the functions are generated, except memory related stuff (eg. memory allocator, storage management, etc...).
If you find any function is missing, report an issue.

# Generate binding
Install [GNU make](https://www.gnu.org/software/make/manual/make.html) and run `make` to re-generate bunding.

# Update

To update to the latest version of dependencies, run `make update`.
After doing this, commit changes and navigate to GitHub.
In Actions tab, manually trigger `Compile cimgui` workflows.

# How does it work?

- `cwrappers/` directory holds C binding for C++ Dear ImGui libraries
- generator bases on `cwrappers/{package_name}_templates` and generates all necessary GO/C code placing it in `{pkgname}/` directories in the root of cimgui-go
- `libs/` contains pre-built shared libraries. `cflags.go` includes and uses to decrease building time.

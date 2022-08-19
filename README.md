# cimgui-go

This project aims to generate go wrapper for Dear ImGui.

Current solution is:
1. Use cimgui's lua generator to generate function and struct definition as json.
2. Generate proper go code from the definition.
3. Use github workflow to compile cimgui and glfw to static lib and place them in /lib folder for further link. I'm not familiar with C compile and link tool, idealy we could compile using cgo from source code, but I have no luck to do it. If you are familiar with it, please give it a try.

Progress:
1. This project works on macOS(arm64/x86) and windows(x64) now, idealy linux should works but I don't have a linux machine to test it. Check out `examples`, cd in and `go run .`.


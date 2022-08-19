# cimgui-go

This project aims to generate go wrapper for Dear ImGui.

It comes with a default backend with GLFW 3.3 and OpenGL 3.2.

Current solution is:
1. Use cimgui's lua generator to generate function and struct definition as json.
2. Generate proper go code from the definition (via manual crafted go program `/cmd/codegen`).
3. Use the backend implementation from imgui (currently glfw and opengl3).
4. Use github workflow to compile cimgui and glfw to static lib and place them in /lib folder for further link. 

Progress:
1. This project works on macOS(arm64/x86) and windows(x64) now, idealy linux should works but I don't have a linux machine to test it. Check out `examples`, cd in and `go run .`.


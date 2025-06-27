## Backend Implementations

> [!note]
> do not confuse with `backend` package.
> This one contains **imgui-level** backend functionalities, such as:
> - NewFrame/EndFrame
> - Rendering
>
> Generally everyghing that is related to imgui interactions with the backend (e.g. plain glfw).

This package contains several sub-packages. Use only thes that you need.

> [!tip]
> Importing package from here has some "interactions" with linker and C code.

> [!warning]
> Code from here **does not** link against any shared object. This means that simply importing e.g. `glfw` imgui impl, will cause linker crash.
> Use this along with any glfw wrapper library such as `backend/glfwbackend` or `go-gl/glfw`.


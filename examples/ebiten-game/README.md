# Ebiten Game example

This shows, how to use cimgui-go in you Ebiten game.

## How to implement EbitenBackend

1. embed `ebitenbackend.EbitenBackend` in your `ebiten.Game` struct.
2. in your `Draw` function, call `(*ebitenbackend.EbitenBackend).Draw(screen)`.
3. in your `Layout` function call `(*ebitenbackend.EbitenBackend).Layout(w, h)`.
4. in your `Update` function, call `(*ebitenbackend.EbitenBackend).BeginFrame()`
5. Now you can use imgui widgets
6. When you are done, call `(*ebitenbackend.EbitenBackend).EndFrame()`
7. In `main` there are several key things you need to remember about:
    - even you're using `*ebitenbackend.EbitenBackend` in your struct, you must call `backend.CreateBackend` for texture manager compatibility (remember about ignoring `backend.CExposerError`).
    - you must run `(*ebitenbackend.EbitenBackend).CreateWindow`. You can change window title later with ebiten api.

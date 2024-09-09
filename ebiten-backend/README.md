## Ebiten backend for cimgui-go

originally created by [Gabriel Ochsenhofer](https://github.com/gabstv) as [ebiten-cimgui](https://github.com/gabstv/ebiten-imgui)
updated by [Brett Bronson](https://github.com/damntourists) as [cimgui-go-ebiten](https://github.com/damntourists/cimgui-go-ebiten)
then it was cloned and modified here by [@gucio321](https://github.com/gucio321) in terms of the [LICENSE](./LICENSE).

## Usage

As ebiten is not (thankfully) a C library, we can make it a
separated package :tada:.

You can use it as another backend for cimgui-go.

```go
        backend, _ = imgui.CreateBackend(ebitenbackend.NewEbitenBackend())
```

But when building your application, you (sadly) need to exclude the
other backends, as they will conflict with ebiten.

```bash
go build -tags ebiten
```

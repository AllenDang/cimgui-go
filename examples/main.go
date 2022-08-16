package main

import (
	"runtime"
	"time"

	"github.com/AllenDang/cimgui-go"
	"github.com/faiface/mainthread"
)

var (
	showDemoWindow bool
)

func init() {
	runtime.LockOSThread()
}

func main() {
	cimgui.CreateContext(0)

	io := cimgui.GetIO()
	io.SetBackendFlags(cimgui.ImGuiBackendFlags_RendererHasVtxOffset)
	io.SetIniFilename("")

	p, err := cimgui.NewGLFW(io, "Test cimgui-go", 1024, 768, 0)
	if err != nil {
		panic(err)
	}

	r, err := cimgui.NewOpenGL3(io, 1.0)
	if err != nil {
		panic(err)
	}

	fonts := io.GetFonts()

	cimgui.FontAtlas_AddFontDefault(fonts, 0)
	fontAtlas := fonts.GetTextureDataRGBA32()
	r.SetFontTexture(fontAtlas)

	mainthread.Run(func() {
		ticker := time.NewTicker(time.Second / time.Duration(p.GetTPS()))
		for !p.ShouldStop() {
			mainthread.Call(func() {
				p.ProcessEvents()

				p.NewFrame()
				r.PreRender([4]float32{0, 0, 0, 1})

				cimgui.NewFrame()

				// Do ui stuff
				cimgui.ShowDemoWindow(&showDemoWindow)

				cimgui.Render()

				r.Render(p.DisplaySize(), p.FramebufferSize(), cimgui.GetDrawData())
				p.PostRender()
			})

			<-ticker.C
		}

		mainthread.Call(func() {
			r.Dispose()
			p.Dispose()
			cimgui.DestroyContext(0)
		})
	})
}

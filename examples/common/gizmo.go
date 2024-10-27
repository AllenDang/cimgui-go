// Example of using imguizmo.
package common

import (
	"github.com/AllenDang/cimgui-go/backend"
	"github.com/AllenDang/cimgui-go/backend/glfwbackend"
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imguizmo"
)

var (
	currentBackend backend.Backend[glfwbackend.GLFWWindowFlags]

	// here are some 4x4 matries. Generally IDK what are they doing - figure it yourself :-)
	// Well, This one is parsed in an editor below.
	View = []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, -7, 1,
	}

	Projection = []float32{
		2.3787, 0, 0, 0,
		0, 3.1716, 0, 0,
		0, 0, -1.0002, -1,
		0, 0, -0.2, 0,
	}

	Ident = []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}

	MOmo = []float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0.5, 0.5, 0.5, 1,
	}

	zmoOP   = int32(imguizmo.TRANSLATE)
	zmoMODE = int32(imguizmo.LOCAL)

	Bounds = []float32{-0.5, -0.5, -0.5, 0.5, 0.5, 0.5}

	// GizmoControls can be enabled in editor and allows mouse control over gizmo
	GizmoControls bool
	UsingGizmo    bool
)

func ShowGizmoDemo() {
	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+400, basePos.Y+460), imgui.CondOnce, imgui.NewVec2(0, 0))
	if imgui.Begin("Gizmo editor") {
		imgui.Checkbox("Enable controls", &GizmoControls)
		imgui.RadioButtonIntPtr("trans", &zmoOP, int32(imguizmo.TRANSLATE))
		imgui.SameLine()
		imgui.RadioButtonIntPtr("rot", &zmoOP, int32(imguizmo.ROTATE))
		imgui.SameLine()
		imgui.RadioButtonIntPtr("scale", &zmoOP, int32(imguizmo.SCALE))
		imgui.SameLine()
		imgui.RadioButtonIntPtr("bounds", &zmoOP, int32(imguizmo.BOUNDS))
		imgui.RadioButtonIntPtr("local", &zmoMODE, int32(imguizmo.LOCAL))
		imgui.SameLine()
		imgui.RadioButtonIntPtr("world", &zmoMODE, int32(imguizmo.WORLD))

		imgui.Separator()
		imgui.Text("View:")
		// float matrixTranslation[3], matrixRotation[3], matrixScale[3];
		// ImGuizmo::DecomposeMatrixToComponents(gizmoMatrix.m16, matrixTranslation, matrixRotation, matrixScale);
		// ImGui::InputFloat3("Tr", matrixTranslation, 3);
		// ImGui::InputFloat3("Rt", matrixRotation, 3);
		// ImGui::InputFloat3("Sc", matrixScale, 3);
		// ImGuizmo::RecomposeMatrixFromComponents(matrixTranslation, matrixRotation, matrixScale, gizmoMatrix.m16);
		matrixTranslation, matrixRotation, matrixScale := make([]float32, 3), make([]float32, 3), make([]float32, 3)
		imguizmo.DecomposeMatrixToComponents(
			&(View[0]),
			&(matrixTranslation[0]),
			&(matrixRotation[0]),
			&(matrixScale[0]),
		)

		imgui.PushItemWidth(80)
		imgui.DragFloatV("TrX", &matrixTranslation[0], .1, 0.0, 0.0, "%.3f", 0)
		imgui.SameLine()
		imgui.DragFloatV("TrY", &matrixTranslation[1], .1, 0.0, 0.0, "%.3f", 0)
		imgui.SameLine()
		imgui.DragFloatV("TrZ", &matrixTranslation[2], .1, 0.0, 0.0, "%.3f", 0)
		imgui.PopItemWidth()

		imgui.Separator()
		imgui.Text("Projection:")
		imgui.PushItemWidth(80)
		imgui.DragFloatV("PrX", &matrixRotation[0], .1, 0.0, 0.0, "%.3f", 0)
		imgui.SameLine()
		imgui.DragFloatV("PrY", &matrixRotation[1], .1, 0.0, 0.0, "%.3f", 0)
		imgui.SameLine()
		imgui.DragFloatV("PrZ", &matrixRotation[2], .1, 0.0, 0.0, "%.3f", 0)
		imgui.PopItemWidth()

		imgui.Separator()
		imgui.Text("Scale:")
		imgui.PushItemWidth(80)
		imgui.DragFloatV("ScX", &matrixScale[0], .1, 0.0, 0.0, "%.3f", 0)
		imgui.SameLine()
		imgui.DragFloatV("ScY", &matrixScale[1], .1, 0.0, 0.0, "%.3f", 0)
		imgui.SameLine()
		imgui.DragFloatV("ScZ", &matrixScale[2], .1, 0.0, 0.0, "%.3f", 0)
		imgui.PopItemWidth()

		if GizmoControls && !imgui.IsWindowFocused() && !UsingGizmo {
			io := imgui.CurrentIO()
			matrixTranslation[2] += io.MouseWheel()
			if leftMouseButton := io.MouseDown()[0]; leftMouseButton {
				const scale = 0.01
				d := io.MouseDelta()
				matrixTranslation[0] += d.X * scale
				matrixTranslation[1] -= d.Y * scale
			}
			if rightMouseButton := io.MouseDown()[1]; rightMouseButton {
				const scale = 0.1
				d := io.MouseDelta()
				matrixRotation[1] += d.X * scale
				/*
					matrixRotation[0] = Clamp(matrixRotation[0]+d.Y*scale*(1-matrixRotation[1]/90), 0, 90)
					matrixRotation[2] = Clamp(matrixRotation[2]+d.Y*scale*matrixRotation[1]/90, 0, 90)
				*/
			}
		}

		imguizmo.RecomposeMatrixFromComponents(
			&(matrixTranslation[0]),
			&(matrixRotation[0]),
			&(matrixScale[0]),
			&(View[0]),
		)
	}
	imgui.End()

	gio := imgui.CurrentIO()
	imguizmo.SetRect(0, 0, gio.DisplaySize().X, gio.DisplaySize().Y)
	imguizmo.SetOrthographic(false)
	imguizmo.DrawGrid(
		&(View[0]),
		&(Projection[0]),
		&(Ident[0]),
		10)

	imguizmo.DrawCubes(
		&(View[0]),
		&(Projection[0]),
		&(MOmo[0]),
		1)

	imguizmo.ManipulateV(
		&(View[0]), &(Projection[0]),
		imguizmo.OPERATION(zmoOP),
		imguizmo.MODE(zmoMODE),
		&(MOmo[0]),
		nil, nil,
		func() *float32 {
			if imguizmo.OPERATION(zmoOP) == imguizmo.BOUNDS {
				return &(Bounds[0])
			}

			return nil
		}(),
		nil)

	imguizmo.ViewManipulateFloat(&(View[0]), 7, imgui.Vec2{0, 0}, imgui.Vec2{128, 128}, 0x01010101)
	UsingGizmo = imguizmo.IsUsingAny()
}

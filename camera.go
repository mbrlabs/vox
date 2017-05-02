package vox

import "github.com/mbrlabs/vox/glm"

type Camera struct {
	Combined   *glm.Mat4
	projection *glm.Mat4
	view       *glm.Mat4
	position   *glm.Vector3
	dirtyView  bool
}

func NewCamera(fov, ratio, near, far float32) *Camera {
	p := glm.NewMat4(false)
	p.Perspective(fov, ratio, near, far)

	v := glm.NewMat4(true)

	cam := &Camera{
		Combined:   glm.NewMat4(false),
		projection: p,
		dirtyView:  true,
		view:       v,
		position:   &glm.Vector3{X: 0, Y: 0, Z: 0},
	}
	cam.Update()
	return cam
}

func (cam *Camera) Move(x, y, z float32) {
	cam.position.Add(x, y, z)
	cam.dirtyView = true
}

func (cam *Camera) Update() {
	if cam.dirtyView {
		cam.view.Identity().Translate(-cam.position.X, -cam.position.Y, -cam.position.Z)

		cam.Combined.Set(cam.projection.Data)
		cam.Combined.Mul(cam.view)
	}
}

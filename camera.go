// Copyright (c) 2017 Marcus Brummer.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

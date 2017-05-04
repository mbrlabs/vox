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

var (
	tmpVec3 *glm.Vector3 = &glm.Vector3{0, 0, 0}
)

type Camera struct {
	Combined   *glm.Mat4
	projection *glm.Mat4
	view       *glm.Mat4

	position  *glm.Vector3
	direction *glm.Vector3
	up        *glm.Vector3
}

func NewCamera(fov, ratio, near, far float32) *Camera {
	p := glm.NewMat4(false)
	p.Perspective(fov, ratio, near, far)

	v := glm.NewMat4(true)

	cam := &Camera{
		Combined:   glm.NewMat4(false),
		projection: p,
		view:       v,
		position:   &glm.Vector3{X: 0, Y: 0, Z: 0},
		direction:  &glm.Vector3{X: 0, Y: 0, Z: -1},
		up:         &glm.Vector3{X: 0, Y: 1, Z: 0},
	}
	cam.Update()
	return cam
}

func (cam *Camera) Move(x, y, z float32) {
	cam.position.Add(x, y, z)
}

func (cam *Camera) Update() {
	t := tmpVec3.SetVector3(cam.position).AddVector3(cam.direction)
	cam.view.LookAt(cam.position, t, cam.up)

	cam.Combined.Set(cam.projection.Data)
	cam.Combined.Mul(cam.view)
}

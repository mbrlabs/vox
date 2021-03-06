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

type FpsCameraController struct {
	MouseSensivity float32 // degress per pixel
	Velocity       float32

	cam         *Camera
	pressedKeys map[Key]bool
	tmp         *glm.Vector3
}

func NewFpsController(cam *Camera) *FpsCameraController {
	return &FpsCameraController{
		MouseSensivity: 0.2,
		Velocity:       50,
		cam:            cam,
		pressedKeys:    make(map[Key]bool),
		tmp:            &glm.Vector3{},
	}
}

func (c *FpsCameraController) Update(delta float32) {
	progress := delta * c.Velocity

	if _, left := c.pressedKeys[KeyA]; left {
		c.tmp.SetVector3(c.cam.direction).Cross(c.cam.up).Norm().Scale(-progress)
		c.cam.Move(c.tmp.X, c.tmp.Y, c.tmp.Z)
	}
	if _, right := c.pressedKeys[KeyD]; right {
		c.tmp.SetVector3(c.cam.direction).Cross(c.cam.up).Norm().Scale(progress)
		c.cam.Move(c.tmp.X, c.tmp.Y, c.tmp.Z)
	}
	if _, forward := c.pressedKeys[KeyW]; forward {
		c.tmp.SetVector3(c.cam.direction).Norm().Scale(progress)
		c.cam.Move(c.tmp.X, c.tmp.Y, c.tmp.Z)
	}
	if _, back := c.pressedKeys[KeyS]; back {
		c.tmp.SetVector3(c.cam.direction).Norm().Scale(-progress)
		c.cam.Move(c.tmp.X, c.tmp.Y, c.tmp.Z)
	}
	if _, up := c.pressedKeys[KeyQ]; up {
		c.tmp.SetVector3(c.cam.up).Norm().Scale(progress)
		c.cam.Move(c.tmp.X, c.tmp.Y, c.tmp.Z)
	}
	if _, down := c.pressedKeys[KeyE]; down {
		c.tmp.SetVector3(c.cam.up).Norm().Scale(-progress)
		c.cam.Move(c.tmp.X, c.tmp.Y, c.tmp.Z)
	}

	c.cam.Update()
}

func (c *FpsCameraController) KeyDown(key Key) bool {
	c.pressedKeys[key] = true
	return false
}

func (c *FpsCameraController) KeyUp(key Key) bool {
	delete(c.pressedKeys, key)
	return false
}

func (c *FpsCameraController) KeyPressed(key Key) bool {
	return false
}

func (c *FpsCameraController) MouseMoved(x, y float64) bool {
	dx := -Vox.DeltaMouseX() * c.MouseSensivity
	dy := -Vox.DeltaMouseY() * c.MouseSensivity
	c.cam.direction.Rotate(c.cam.up, dx)
	c.tmp.SetVector3(c.cam.direction).Cross(c.cam.up).Norm()
	c.cam.direction.Rotate(c.tmp, dy)

	return false
}

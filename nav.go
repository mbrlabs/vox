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

import "github.com/go-gl/glfw/v3.2/glfw"

type FpsCameraController struct {
	Velocity float32

	cam         *Camera
	pressedKeys map[glfw.Key]bool
}

func NewFpsController(cam *Camera) *FpsCameraController {
	return &FpsCameraController{
		Velocity:    1,
		cam:         cam,
		pressedKeys: make(map[glfw.Key]bool),
	}
}

func (c *FpsCameraController) Update(delta float32) {
	// left
	if _, left := c.pressedKeys[glfw.KeyA]; left {
		c.cam.Move(-c.Velocity, 0, 0)
	}

	// right
	if _, right := c.pressedKeys[glfw.KeyD]; right {
		c.cam.Move(c.Velocity, 0, 0)
	}

	// forward
	if _, forward := c.pressedKeys[glfw.KeyW]; forward {
		c.cam.Move(0, 0, -c.Velocity)
	}

	// backward
	if _, back := c.pressedKeys[glfw.KeyS]; back {
		c.cam.Move(0, 0, c.Velocity)
	}

	// up
	if _, up := c.pressedKeys[glfw.KeyQ]; up {
		c.cam.Move(0, c.Velocity, 0)
	}

	// down
	if _, down := c.pressedKeys[glfw.KeyE]; down {
		c.cam.Move(0, -c.Velocity, 0)
	}

	c.cam.Update()
}

func (c *FpsCameraController) KeyDown(key glfw.Key) bool {
	c.pressedKeys[key] = true
	return false
}

func (c *FpsCameraController) KeyUp(key glfw.Key) bool {
	delete(c.pressedKeys, key)
	return false
}

func (c *FpsCameraController) KeyPressed(key glfw.Key) bool {
	return false
}

func (c *FpsCameraController) MouseMoved(x, y float64) bool {
	return false
}

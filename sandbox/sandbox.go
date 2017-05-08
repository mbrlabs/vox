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

package main

import (
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/mbrlabs/vox"
)

type Sandbox struct {
	fpsController *vox.FpsCameraController
	cam           *vox.Camera
	renderer      *vox.WorldRenderer
	world         *vox.World

	oldX, dx float32
}

func (s *Sandbox) Create() {
	vox.Vox.AddKeyListener(s)

	ratio := float32(windowWidth) / float32(windowHeight)
	s.cam = vox.NewCamera(70, ratio, 0.01, 1000)
	s.cam.Move(8, 15, 40)
	s.cam.Update()

	s.fpsController = vox.NewFpsController(s.cam)
	vox.Vox.AddMouseListener(s.fpsController)
	vox.Vox.AddKeyListener(s.fpsController)

	s.world = vox.NewWorld()
	s.world.BlockBank.AddType(&vox.BlockType{ID: 0x01, Color: vox.ColorRed.Copy()})
	s.world.BlockBank.AddType(&vox.BlockType{ID: 0x02, Color: vox.ColorGreen.Copy()})
	s.world.BlockBank.AddType(&vox.BlockType{ID: 0x03, Color: vox.ColorBlue.Copy()})
	s.world.BlockBank.AddType(&vox.BlockType{ID: 0x04, Color: vox.ColorTeal.Copy()})
	s.world.CreateChunk(0, 0, 0)
	s.world.CreateChunk(-1, 0, 0)
	s.world.CreateChunk(1, 0, 0)

	s.renderer = vox.NewWorldRenderer()

	gl.Enable(gl.DEPTH_TEST)
}

func (s *Sandbox) Dispose() {
	s.renderer.Dispose()
}

func (s *Sandbox) Update(delta float32) {
	s.fpsController.Update(delta)
}

func (s *Sandbox) Render(delta float32) {
	// clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.95, 0.95, 0.95, 0.0)

	// render world
	s.renderer.Render(s.cam, s.world)
}

func (s *Sandbox) Resize(width, height int) {

}

func (s *Sandbox) KeyDown(key vox.Key) bool {
	if key == vox.KeyEscape {
		vox.Vox.Exit()
	}
	return false
}

func (s *Sandbox) KeyUp(key vox.Key) bool {
	return false
}

func (s *Sandbox) KeyPressed(key vox.Key) bool {
	return false
}

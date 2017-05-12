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
	oldX, dx      float32

	blockBank *vox.BlockBank
	atlas     *vox.TextureAtlas

	fpsLogger *vox.FpsLogger
	spawns    int
}

func (s *Sandbox) Create() {
	vox.Vox.AddKeyListener(s)

	// load assets
	s.atlas = vox.NewTextureAtlas("assets/atlas.json", "assets/atlas.png")
	s.blockBank = vox.NewBlockBank()
	for _, t := range createBlockTypes(s.atlas) {
		s.blockBank.AddType(t)
	}

	ratio := float32(windowWidth) / float32(windowHeight)
	s.cam = vox.NewCamera(70, ratio, 0.01, 1000)
	s.cam.Move(8, 15, 40)
	s.cam.Update()

	s.fpsController = vox.NewFpsController(s.cam)
	vox.Vox.AddMouseListener(s.fpsController)
	vox.Vox.AddKeyListener(s.fpsController)

	s.world = vox.NewWorld(s.blockBank, &vox.CulledMesher{}, &vox.FlatGenerator{})

	// create huge flat 5x2x5 cube
	for x := 0; x < 15; x++ {
		for z := 0; z < 15; z++ {
			for y := 0; y < 3; y++ {
				s.world.GenerateNewChunk(x, -5+y, z)
			}
		}
	}

	// create interesting connected chunks
	s.world.GenerateNewChunk(0, -1, 0)
	s.world.GenerateNewChunk(0, 0, 0)
	s.world.GenerateNewChunk(0, 1, 0)
	s.world.GenerateNewChunk(0, 2, 0)
	s.world.GenerateNewChunk(1, 2, 0)
	s.world.GenerateNewChunk(2, 2, 0)
	s.world.GenerateNewChunk(2, 2, 1)
	s.world.GenerateNewChunk(2, 2, 2)

	s.renderer = vox.NewWorldRenderer()
	s.fpsLogger = &vox.FpsLogger{}

	gl.Enable(gl.DEPTH_TEST)
}

func (s *Sandbox) Dispose() {
	s.renderer.Dispose()
	s.atlas.Dispose()
}

func (s *Sandbox) Update(delta float32) {
	s.world.Update()

	s.fpsController.Update(delta)
	s.fpsLogger.Log(delta)
}

func (s *Sandbox) Render(delta float32) {
	// clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.95, 0.95, 0.95, 0.0)

	// render world
	s.atlas.Bind()
	s.renderer.Render(s.cam, s.world)
}

func (s *Sandbox) Resize(width, height int) {

}

func (s *Sandbox) KeyDown(key vox.Key) bool {
	if key == vox.KeyEscape {
		vox.Vox.Exit()
	} else if key == vox.KeyEnter {
		for x := 0; x < 10; x++ {
			for z := 0; z < 10; z++ {
				for y := 0; y < 3; y++ {
					s.world.GenerateNewChunk(9+x, s.spawns+y, z)
				}
			}
		}
		s.spawns += 3
	} else if key == vox.KeyDelete {
		for x := 0; x < 10; x++ {
			for z := 0; z < 10; z++ {
				for y := 0; y < 3; y++ {
					s.world.RemoveChunk(9+x, s.spawns+y-3, z)
				}
			}
		}
		s.spawns -= 3
	}

	return false
}

func (s *Sandbox) KeyUp(key vox.Key) bool {
	return false
}

func (s *Sandbox) KeyPressed(key vox.Key) bool {
	return false
}

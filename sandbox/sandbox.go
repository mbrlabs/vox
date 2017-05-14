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
	fpsController   *vox.FpsCameraController
	worldController *vox.InfiniteWorldController
	cam             *vox.Camera
	renderer        *vox.WorldRenderer
	world           *vox.World
	env             *vox.Environment
	oldX, dx        float32

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

	// build camera
	ratio := float32(windowWidth) / float32(windowHeight)
	s.cam = vox.NewCamera(70, ratio, 0.01, 1000)
	s.cam.Move(0, vox.ChunkHeight*2, 0)
	s.cam.Update()

	// build world
	s.world = vox.NewWorld(s.blockBank, &vox.CulledMesher{}, vox.NewSimplexGenerator(16726))

	// setup fps controller
	s.fpsController = vox.NewFpsController(s.cam)
	vox.Vox.AddMouseListener(s.fpsController)
	vox.Vox.AddKeyListener(s.fpsController)

	// misc
	s.renderer = vox.NewWorldRenderer()
	s.fpsLogger = &vox.FpsLogger{}
	s.worldController = vox.NewInifinteWorldController(s.cam, s.world)
	s.env = vox.NewEnvironment()

	gl.Enable(gl.DEPTH_TEST)
}

func (s *Sandbox) Dispose() {
	s.renderer.Dispose()
	s.atlas.Dispose()
}

func (s *Sandbox) Update(delta float32) {
	s.worldController.Update()
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
	s.renderer.Render(s.cam, s.world, s.env)
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

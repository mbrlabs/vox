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
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mbrlabs/vox"
	"runtime"
)

const (
	windowTitle  = "Chunk example"
	windowWidth  = 1024
	windowHeight = 768
)

type ChunkDemo struct {
	fpsController *vox.FpsCameraController
	cam           *vox.Camera
	renderer      *vox.WorldRenderer
	world         *vox.World

	oldX, dx float32
}

func (d *ChunkDemo) Create() {
	vox.Vox().AddKeyListener(d)

	ratio := float32(windowWidth) / float32(windowHeight)
	d.cam = vox.NewCamera(70, ratio, 0.01, 1000)
	d.cam.Move(0, 0, 50)
	d.cam.Update()

	d.fpsController = vox.NewFpsController(d.cam)
	vox.Vox().AddMouseListener(d.fpsController)
	vox.Vox().AddKeyListener(d.fpsController)

	d.world = vox.NewWorld()
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x01, Color: vox.ColorRed.Copy()})
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x02, Color: vox.ColorGreen.Copy()})
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x03, Color: vox.ColorBlue.Copy()})
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x04, Color: vox.ColorTeal.Copy()})
	d.world.GenerateDebugWorld()

	d.renderer = vox.NewWorldRenderer()

	gl.Enable(gl.DEPTH_TEST)
}

func (d *ChunkDemo) Dispose() {

}

func (d *ChunkDemo) Update(delta float32) {
	d.fpsController.Update(delta)
}

func (d *ChunkDemo) Render(delta float32) {
	// clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.95, 0.95, 0.95, 0.0)

	// render world
	d.renderer.Render(d.cam, d.world)
}

func (d *ChunkDemo) Resize(width, height int) {

}

func (d *ChunkDemo) KeyDown(key glfw.Key) bool {
	if key == glfw.KeyEscape {
		vox.Vox().Exit()
	}
	return false
}

func (d *ChunkDemo) KeyUp(key glfw.Key) bool {
	return false
}

func (d *ChunkDemo) KeyPressed(key glfw.Key) bool {
	return false
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

func init() {
	runtime.LockOSThread()
}

func main() {
	window := vox.NewWindow(&vox.WindowConfig{
		Height:     windowHeight,
		Width:      windowWidth,
		Title:      windowTitle,
		Resizable:  false,
		Fullscreen: false,
		Vsync:      true,
		//HiddenCursor: true,
	})

	window.Start(&ChunkDemo{})
}

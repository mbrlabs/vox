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
	"runtime"

	"math"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/mbrlabs/vox"
)

const (
	windowTitle  = "Chunk example"
	windowWidth  = 1024
	windowHeight = 768
)

type ChunkDemo struct {
	camera   *vox.Camera
	renderer *vox.WorldRenderer
	world    *vox.World

	oldX, dx float32
}

func (d *ChunkDemo) Create() {
	ratio := float32(windowWidth) / float32(windowHeight)
	d.camera = vox.NewCamera(70, ratio, 0.01, 1000)

	d.renderer = vox.NewWorldRenderer()
	vox.Vox().AddMouseListener(d)
	d.camera.Move(0, 0, 50)

	d.world = vox.NewWorld()
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x01, Color: vox.ColorRed.Copy()})
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x02, Color: vox.ColorGreen.Copy()})
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x03, Color: vox.ColorBlue.Copy()})
	d.world.BlockBank.AddType(&vox.BlockType{ID: 0x04, Color: vox.ColorTeal.Copy()})
	d.world.GenerateDebugWorld()

	gl.Enable(gl.DEPTH_TEST)
}

func (d *ChunkDemo) Dispose() {

}

func (d *ChunkDemo) Update(delta float32) {
	if math.Abs(float64(d.dx)) < 5 {
		d.camera.Move(d.dx, 0, 0)
	}

	d.camera.Update()
}

func (d *ChunkDemo) Render(delta float32) {
	// clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.95, 0.95, 0.95, 0.0)

	// render world
	d.renderer.Render(d.camera, d.world)
}

func (d *ChunkDemo) Resize(width, height int) {

}

func (d *ChunkDemo) MouseMoved(x, y float64) bool {
	d.dx = (float32(x) - d.oldX) * 0.1
	d.oldX = float32(x)

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
	})

	window.Start(&ChunkDemo{})
}

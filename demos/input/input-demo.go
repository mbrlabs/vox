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

	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"

	"github.com/mbrlabs/vox"
)

const (
	windowTitle  = "Input example"
	windowWidth  = 1024
	windowHeight = 768
)

type InputDemo struct {
}

func (d *InputDemo) Create() {
	vox := vox.Vox()
	vox.AddMouseListener(d)
	vox.AddKeyListener(d)
}

func (d *InputDemo) Dispose() {

}

func (d *InputDemo) Update(delta float32) {

}

func (d *InputDemo) Render(delta float32) {
	// clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.95, 0.95, 0.95, 0.0)

}

func (d *InputDemo) Resize(width, height int) {

}

func (d *InputDemo) MouseMoved(x, y float64) bool {
	fmt.Printf("Mouse position: %v,%v\n", x, y)
	return false
}

func (d *InputDemo) KeyDown(key glfw.Key) bool {
	fmt.Printf("Key down: %v\n", key)
	return false
}

func (d *InputDemo) KeyUp(key glfw.Key) bool {
	fmt.Printf("Key up: %v\n", key)
	return false
}

func (d *InputDemo) KeyPressed(key glfw.Key) bool {
	fmt.Printf("Key pressed: %v\n", key)
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

	window.Start(&InputDemo{})
}

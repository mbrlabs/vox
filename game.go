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

import (
	"fmt"
	"log"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Game interface {
	Disposable
	Create()
	Resize(width, height int)
	Render(delta float32)
	Update(delate float32)
}

type Window struct {
	glfwWindow *glfw.Window
}

type WindowConfig struct {
	Height     int
	Width      int
	Title      string
	Resizable  bool
	Fullscreen bool
	Vsync      bool
}

func NewWindow(config *WindowConfig) *Window {
	// setup glfw
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	// window hints
	if config.Resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// create window & make current
	window, err := glfw.CreateWindow(config.Width, config.Height, config.Title, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// vsync?
	if config.Vsync {
		glfw.SwapInterval(1)
	}

	// setup opengl
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Using OpenGL version", version)

	return &Window{
		glfwWindow: window,
	}
}

func (w *Window) Dispose() {
	glfw.Terminate()
}

func (w *Window) Start(game Game) {
	defer game.Dispose()
	game.Create()

	for !w.glfwWindow.ShouldClose() {
		// game update
		game.Update(1)
		game.Render(1)

		// glfw update
		w.glfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
}

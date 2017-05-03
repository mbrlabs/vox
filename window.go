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
	"time"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

// ----------------------------------------------------------------------------

// KeyListener todo
type KeyListener interface {
	KeyDown(keycode int) bool
	KeyUp(keycode int) bool
}

// ----------------------------------------------------------------------------

// MouseListener todo
type MouseListener interface {
	//MouseDown(x, y float64) bool
	//MouseUp(x, y float64) bool
	MouseMoved(x, y float64) bool
}

// ----------------------------------------------------------------------------

// WindowConfig todo
type WindowConfig struct {
	Height     int
	Width      int
	Title      string
	Resizable  bool
	Fullscreen bool
	Vsync      bool
}

// ----------------------------------------------------------------------------

// Window todo
type Window struct {
	glfwWindow *glfw.Window
	lastFrame  int64

	keyListeners   []KeyListener
	mouseListeners []MouseListener
}

// ----------------------------------------------------------------------------

// NewWindow todo
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

	// vsync
	if config.Vsync {
		glfw.SwapInterval(1)
	}

	// setup opengl
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("Using OpenGL version", version)

	win := &Window{
		glfwWindow: window,
	}
	setupVox(win)
	win.setupInputCallbacks()

	return win
}

// Dispose todo
func (w *Window) Dispose() {
	glfw.Terminate()
}

// Start todo
func (w *Window) Start(game Game) {
	defer game.Dispose()
	game.Create()
	w.lastFrame = time.Now().Unix()

	for !w.glfwWindow.ShouldClose() {
		// game update
		deltaNano := time.Now().UnixNano() - w.lastFrame
		deltaSeconds := float32(float64(deltaNano) / 1000000000.0)
		game.Update(deltaSeconds)
		game.Render(deltaSeconds)
		w.lastFrame = time.Now().UnixNano()

		// glfw update
		w.glfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
}

func (w *Window) setupInputCallbacks() {
	// mouse moved callback
	w.glfwWindow.SetCursorPosCallback(func(win *glfw.Window, xpos, ypos float64) {
		w.onMouseMoved(xpos, ypos)
	})
}

// AddKeyListener todo
func (w *Window) addKeyListener(listener KeyListener) {
	w.keyListeners = append(w.keyListeners, listener)
}

// AddMouseListener todo
func (w *Window) addMouseListener(listener MouseListener) {
	w.mouseListeners = append(w.mouseListeners, listener)
}

// Callback for glfw event
func (w *Window) onMouseMoved(xpos, ypos float64) {
	for _, listener := range w.mouseListeners {
		if listener.MouseMoved(xpos, ypos) {
			break
		}
	}
}

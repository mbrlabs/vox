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
	KeyDown(key glfw.Key) bool
	KeyUp(key glfw.Key) bool
	KeyPressed(key glfw.Key) bool
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
	glfwWindow    *glfw.Window
	lastFrame     int64
	lastMouseX    float32
	lastMouseY    float32
	exitRequested bool

	keyListeners   []KeyListener
	mouseListeners []MouseListener
	deltaTime      float32
	deltaX         float32
	deltaY         float32
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
		glfwWindow:    window,
		exitRequested: false,
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

	for !w.glfwWindow.ShouldClose() && !w.exitRequested {
		// game update
		deltaNano := time.Now().UnixNano() - w.lastFrame
		w.deltaTime = float32(float64(deltaNano) / 1000000000.0)
		game.Update(w.deltaTime)
		game.Render(w.deltaTime)
		w.lastFrame = time.Now().UnixNano()

		// glfw update
		w.glfwWindow.SwapBuffers()
		glfw.PollEvents()
	}
}

func (w *Window) setupInputCallbacks() {
	// mouse moved callback
	first := true
	w.glfwWindow.SetCursorPosCallback(func(win *glfw.Window, xpos, ypos float64) {
		if first {
			w.deltaX = 0
			w.deltaY = 0
		} else {
			w.deltaX = w.lastMouseX - float32(xpos)
			w.deltaY = w.lastMouseY - float32(ypos)
		}

		w.lastMouseX = float32(xpos)
		w.lastMouseY = float32(ypos)
		first = false

		for _, listener := range w.mouseListeners {
			if listener.MouseMoved(xpos, ypos) {
				break
			}
		}
	})

	// key callback
	w.glfwWindow.SetKeyCallback(func(win *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		for _, listener := range w.keyListeners {
			if action == glfw.Press {
				if listener.KeyDown(key) {
					break
				}
			} else if action == glfw.Release {
				if listener.KeyUp(key) {
					break
				}
			} else {
				if listener.KeyPressed(key) {
					break
				}
			}
		}

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

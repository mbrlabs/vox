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
	KeyDown(key Key) bool
	KeyUp(key Key) bool
	KeyPressed(key Key) bool
}

// ----------------------------------------------------------------------------

// MouseListener todo
type MouseListener interface {
	//MouseDown(x, y float64) bool
	//MouseUp(x, y float64) bool
	MouseMoved(x, y float64) bool
}

// ----------------------------------------------------------------------------

type Key uint16

const (
	KeyUnknown      Key = 0
	KeySpace        Key = 1
	KeyApostrophe   Key = 2
	KeyComma        Key = 3
	KeyMinus        Key = 4
	KeyPeriod       Key = 5
	KeySlash        Key = 6
	Key0            Key = 7
	Key1            Key = 8
	Key2            Key = 9
	Key3            Key = 10
	Key4            Key = 11
	Key5            Key = 12
	Key6            Key = 13
	Key7            Key = 14
	Key8            Key = 15
	Key9            Key = 16
	KeySemicolon    Key = 17
	KeyEqual        Key = 18
	KeyA            Key = 19
	KeyB            Key = 20
	KeyC            Key = 21
	KeyD            Key = 22
	KeyE            Key = 23
	KeyF            Key = 24
	KeyG            Key = 25
	KeyH            Key = 26
	KeyI            Key = 27
	KeyJ            Key = 28
	KeyK            Key = 29
	KeyL            Key = 30
	KeyM            Key = 31
	KeyN            Key = 32
	KeyO            Key = 33
	KeyP            Key = 34
	KeyQ            Key = 35
	KeyR            Key = 36
	KeyS            Key = 37
	KeyT            Key = 38
	KeyU            Key = 39
	KeyV            Key = 40
	KeyW            Key = 41
	KeyX            Key = 42
	KeyY            Key = 43
	KeyZ            Key = 44
	KeyLeftBracket  Key = 45
	KeyBackslash    Key = 46
	KeyRightBracket Key = 47
	KeyGraveAccent  Key = 48
	KeyWorld1       Key = 49
	KeyWorld2       Key = 50
	KeyEscape       Key = 51
	KeyEnter        Key = 52
	KeyTab          Key = 53
	KeyBackspace    Key = 54
	KeyInsert       Key = 55
	KeyDelete       Key = 56
	KeyRight        Key = 57
	KeyLeft         Key = 58
	KeyDown         Key = 59
	KeyUp           Key = 60
	KeyPageUp       Key = 61
	KeyPageDown     Key = 62
	KeyHome         Key = 63
	KeyEnd          Key = 64
	KeyCapsLock     Key = 65
	KeyScrollLock   Key = 66
	KeyNumLock      Key = 67
	KeyPrintScreen  Key = 68
	KeyPause        Key = 69
	KeyF1           Key = 70
	KeyF2           Key = 71
	KeyF3           Key = 72
	KeyF4           Key = 73
	KeyF5           Key = 74
	KeyF6           Key = 75
	KeyF7           Key = 76
	KeyF8           Key = 77
	KeyF9           Key = 78
	KeyF10          Key = 79
	KeyF11          Key = 80
	KeyF12          Key = 81
	KeyF13          Key = 82
	KeyF14          Key = 83
	KeyF15          Key = 84
	KeyF16          Key = 85
	KeyF17          Key = 86
	KeyF18          Key = 87
	KeyF19          Key = 88
	KeyF20          Key = 89
	KeyF21          Key = 90
	KeyF22          Key = 91
	KeyF23          Key = 92
	KeyF24          Key = 93
	KeyF25          Key = 94
	KeyKP0          Key = 95
	KeyKP1          Key = 96
	KeyKP2          Key = 97
	KeyKP3          Key = 98
	KeyKP4          Key = 99
	KeyKP5          Key = 100
	KeyKP6          Key = 101
	KeyKP7          Key = 102
	KeyKP8          Key = 103
	KeyKP9          Key = 104
	KeyKPDecimal    Key = 105
	KeyKPDivide     Key = 106
	KeyKPMultiply   Key = 107
	KeyKPSubtract   Key = 108
	KeyKPAdd        Key = 109
	KeyKPEnter      Key = 110
	KeyKPEqual      Key = 111
	KeyLeftShift    Key = 112
	KeyLeftControl  Key = 113
	KeyLeftAlt      Key = 114
	KeyLeftSuper    Key = 115
	KeyRightShift   Key = 116
	KeyRightControl Key = 117
	KeyRightAlt     Key = 118
	KeyRightSuper   Key = 119
	KeyMenu         Key = 120
)

// ----------------------------------------------------------------------------

// WindowConfig todo
type WindowConfig struct {
	Height       int
	Width        int
	Title        string
	Resizable    bool
	Fullscreen   bool
	Vsync        bool
	HiddenCursor bool
}

// ----------------------------------------------------------------------------

// Window todo
type Window struct {
	glfwWindow *glfw.Window

	lastFrame    int64
	fpsTime      float32
	fps          int
	frameCounter int

	exitRequested  bool
	keyListeners   []KeyListener
	mouseListeners []MouseListener
	deltaTime      float32
	deltaX         float32
	deltaY         float32
	lastMouseX     float32
	lastMouseY     float32

	keyMap [glfw.KeyLast + 1]Key
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

	if config.HiddenCursor {
		window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
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
	win.fillKeyMap()
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
		// update delta time
		deltaNano := time.Now().UnixNano() - w.lastFrame
		w.deltaTime = float32(float64(deltaNano) / 1000000000.0)
		w.lastFrame = time.Now().UnixNano()

		// update & render game
		game.Update(w.deltaTime)
		game.Render(w.deltaTime)

		// update fps
		w.frameCounter++
		w.fpsTime += w.deltaTime
		if w.fpsTime >= 1.0 {
			w.fps = w.frameCounter
			w.fpsTime = 0
			w.frameCounter = 0
		}

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
		vk := w.convertKey(key)
		for _, listener := range w.keyListeners {
			if action == glfw.Press {
				if listener.KeyDown(vk) {
					break
				}
			} else if action == glfw.Release {
				if listener.KeyUp(vk) {
					break
				}
			} else {
				if listener.KeyPressed(vk) {
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

func (w *Window) convertKey(key glfw.Key) Key {
	if key == glfw.KeyUnknown {
		return KeyUnknown
	}
	return w.keyMap[key]
}

func (w *Window) fillKeyMap() {
	w.keyMap[glfw.KeySpace] = KeySpace
	w.keyMap[glfw.KeyApostrophe] = KeyApostrophe
	w.keyMap[glfw.KeyComma] = KeyComma
	w.keyMap[glfw.KeyMinus] = KeyMinus
	w.keyMap[glfw.KeyPeriod] = KeyPeriod
	w.keyMap[glfw.KeySlash] = KeySlash
	w.keyMap[glfw.Key0] = Key0
	w.keyMap[glfw.Key1] = Key1
	w.keyMap[glfw.Key2] = Key2
	w.keyMap[glfw.Key3] = Key3
	w.keyMap[glfw.Key4] = Key4
	w.keyMap[glfw.Key5] = Key5
	w.keyMap[glfw.Key6] = Key6
	w.keyMap[glfw.Key7] = Key7
	w.keyMap[glfw.Key8] = Key8
	w.keyMap[glfw.Key9] = Key9
	w.keyMap[glfw.KeySemicolon] = KeySemicolon
	w.keyMap[glfw.KeyEqual] = KeyEqual
	w.keyMap[glfw.KeyA] = KeyA
	w.keyMap[glfw.KeyB] = KeyB
	w.keyMap[glfw.KeyC] = KeyC
	w.keyMap[glfw.KeyD] = KeyD
	w.keyMap[glfw.KeyE] = KeyE
	w.keyMap[glfw.KeyF] = KeyF
	w.keyMap[glfw.KeyG] = KeyG
	w.keyMap[glfw.KeyH] = KeyH
	w.keyMap[glfw.KeyI] = KeyI
	w.keyMap[glfw.KeyJ] = KeyJ
	w.keyMap[glfw.KeyK] = KeyK
	w.keyMap[glfw.KeyL] = KeyL
	w.keyMap[glfw.KeyM] = KeyM
	w.keyMap[glfw.KeyN] = KeyN
	w.keyMap[glfw.KeyO] = KeyO
	w.keyMap[glfw.KeyP] = KeyP
	w.keyMap[glfw.KeyQ] = KeyQ
	w.keyMap[glfw.KeyR] = KeyR
	w.keyMap[glfw.KeyS] = KeyS
	w.keyMap[glfw.KeyT] = KeyT
	w.keyMap[glfw.KeyU] = KeyU
	w.keyMap[glfw.KeyV] = KeyV
	w.keyMap[glfw.KeyW] = KeyW
	w.keyMap[glfw.KeyX] = KeyX
	w.keyMap[glfw.KeyY] = KeyY
	w.keyMap[glfw.KeyZ] = KeyZ
	w.keyMap[glfw.KeyLeftBracket] = KeyLeftBracket
	w.keyMap[glfw.KeyBackslash] = KeyBackslash
	w.keyMap[glfw.KeyRightBracket] = KeyRightBracket
	w.keyMap[glfw.KeyGraveAccent] = KeyGraveAccent
	w.keyMap[glfw.KeyWorld1] = KeyWorld1
	w.keyMap[glfw.KeyWorld2] = KeyWorld2
	w.keyMap[glfw.KeyEscape] = KeyEscape
	w.keyMap[glfw.KeyEnter] = KeyEnter
	w.keyMap[glfw.KeyTab] = KeyTab
	w.keyMap[glfw.KeyBackspace] = KeyBackspace
	w.keyMap[glfw.KeyInsert] = KeyInsert
	w.keyMap[glfw.KeyDelete] = KeyDelete
	w.keyMap[glfw.KeyRight] = KeyRight
	w.keyMap[glfw.KeyLeft] = KeyLeft
	w.keyMap[glfw.KeyDown] = KeyDown
	w.keyMap[glfw.KeyUp] = KeyUp
	w.keyMap[glfw.KeyPageUp] = KeyPageUp
	w.keyMap[glfw.KeyPageDown] = KeyPageDown
	w.keyMap[glfw.KeyHome] = KeyHome
	w.keyMap[glfw.KeyEnd] = KeyEnd
	w.keyMap[glfw.KeyCapsLock] = KeyCapsLock
	w.keyMap[glfw.KeyScrollLock] = KeyScrollLock
	w.keyMap[glfw.KeyNumLock] = KeyNumLock
	w.keyMap[glfw.KeyPrintScreen] = KeyPrintScreen
	w.keyMap[glfw.KeyPause] = KeyPause
	w.keyMap[glfw.KeyF1] = KeyF1
	w.keyMap[glfw.KeyF2] = KeyF2
	w.keyMap[glfw.KeyF3] = KeyF3
	w.keyMap[glfw.KeyF4] = KeyF4
	w.keyMap[glfw.KeyF5] = KeyF5
	w.keyMap[glfw.KeyF6] = KeyF6
	w.keyMap[glfw.KeyF7] = KeyF7
	w.keyMap[glfw.KeyF8] = KeyF8
	w.keyMap[glfw.KeyF9] = KeyF9
	w.keyMap[glfw.KeyF10] = KeyF10
	w.keyMap[glfw.KeyF11] = KeyF11
	w.keyMap[glfw.KeyF12] = KeyF12
	w.keyMap[glfw.KeyF13] = KeyF13
	w.keyMap[glfw.KeyF14] = KeyF14
	w.keyMap[glfw.KeyF15] = KeyF15
	w.keyMap[glfw.KeyF16] = KeyF16
	w.keyMap[glfw.KeyF17] = KeyF17
	w.keyMap[glfw.KeyF18] = KeyF18
	w.keyMap[glfw.KeyF19] = KeyF19
	w.keyMap[glfw.KeyF20] = KeyF20
	w.keyMap[glfw.KeyF21] = KeyF21
	w.keyMap[glfw.KeyF22] = KeyF22
	w.keyMap[glfw.KeyF23] = KeyF23
	w.keyMap[glfw.KeyF24] = KeyF24
	w.keyMap[glfw.KeyF25] = KeyF25
	w.keyMap[glfw.KeyKP0] = KeyKP0
	w.keyMap[glfw.KeyKP1] = KeyKP1
	w.keyMap[glfw.KeyKP2] = KeyKP2
	w.keyMap[glfw.KeyKP3] = KeyKP3
	w.keyMap[glfw.KeyKP4] = KeyKP4
	w.keyMap[glfw.KeyKP5] = KeyKP5
	w.keyMap[glfw.KeyKP6] = KeyKP6
	w.keyMap[glfw.KeyKP7] = KeyKP7
	w.keyMap[glfw.KeyKP8] = KeyKP8
	w.keyMap[glfw.KeyKP9] = KeyKP9
	w.keyMap[glfw.KeyKPDecimal] = KeyKPDecimal
	w.keyMap[glfw.KeyKPDivide] = KeyKPDivide
	w.keyMap[glfw.KeyKPMultiply] = KeyKPMultiply
	w.keyMap[glfw.KeyKPSubtract] = KeyKPSubtract
	w.keyMap[glfw.KeyKPAdd] = KeyKPAdd
	w.keyMap[glfw.KeyKPEnter] = KeyKPEnter
	w.keyMap[glfw.KeyKPEqual] = KeyKPEqual
	w.keyMap[glfw.KeyLeftShift] = KeyLeftShift
	w.keyMap[glfw.KeyLeftControl] = KeyLeftControl
	w.keyMap[glfw.KeyLeftAlt] = KeyLeftAlt
	w.keyMap[glfw.KeyLeftSuper] = KeyLeftSuper
	w.keyMap[glfw.KeyRightShift] = KeyRightShift
	w.keyMap[glfw.KeyRightControl] = KeyRightControl
	w.keyMap[glfw.KeyRightAlt] = KeyRightAlt
	w.keyMap[glfw.KeyRightSuper] = KeyRightSuper
	w.keyMap[glfw.KeyMenu] = KeyMenu
}

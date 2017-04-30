package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mbrlabs/gocraft/gocraft"
)

const (
	version      = "v.0.1.0"
	windowTitle  = "Gocraft " + version
	windowWidth  = 1024
	windowHeight = 768
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func setupWindow() *glfw.Window {
	// window hints
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// create window & make current
	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	return window
}

func setupOpenGL() {
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}

func main() {
	// setup glfw
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window := setupWindow()
	setupOpenGL()

	// test shaders
	gocraft.NewShader("shaders/world.frag.glsl", "shaders/world.vert.glsl")

	// game loop
	for !window.ShouldClose() {
		// clear window
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(0.5, 0.5, 0.5, 0.0)

		// glfw update
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

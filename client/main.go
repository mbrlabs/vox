package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/mbrlabs/gocraft"
	"github.com/mbrlabs/gocraft/glm"
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
	shader, err := gocraft.NewShader("shaders/world.vert.glsl", "shaders/world.frag.glsl")
	if err != nil {
		panic(err)
	}
	defer shader.Dispose()

	// cube positions
	verts := []float32{
		// front
		-1.0, -1.0, 1.0,
		1.0, -1.0, 1.0,
		1.0, 1.0, 1.0,
		-1.0, 1.0, 1.0,
		// back
		-1.0, -1.0, -1.0,
		1.0, -1.0, -1.0,
		1.0, 1.0, -1.0,
		-1.0, 1.0, -1.0,
	}
	// cube indices
	indices := []uint16{
		// front
		0, 1, 2,
		2, 3, 0,
		// top
		1, 5, 6,
		6, 2, 1,
		// back
		7, 6, 5,
		5, 4, 7,
		// bottom
		4, 0, 3,
		3, 7, 4,
		// left
		4, 5, 1,
		1, 0, 4,
		// right
		3, 2, 6,
		6, 7, 3,
	}
	uvs := []float32{1, 2}
	normals := []float32{1, 2}

	vao := gocraft.NewVao()
	defer vao.Dispose()
	vao.Load(verts, indices, uvs, normals)

	cam := gocraft.NewCamera(windowWidth, windowHeight)

	mvpUniform := gl.GetUniformLocation(shader.ID, gl.Str("u_mvp\x00"))

	model := glm.NewMat4(true)
	model.Translation(0.5, 0.5, -10)
	mvp := glm.NewMat4(true)

	gl.Enable(gl.DEPTH_TEST)

	// game loop
	for !window.ShouldClose() {
		// clear window
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(0.5, 1, 1, 0.0)

		model.Rotate(5, 0, -1, 0)
		cam.Update()

		mvp.Set(cam.Combined.Data)
		mvp.Mul(model)

		shader.Enable()
		vao.Bind()

		gl.UniformMatrix4fv(mvpUniform, 1, false, &mvp.Data[0])
		gl.DrawElements(gl.TRIANGLES, int32(len(indices)), gl.UNSIGNED_SHORT, gl.PtrOffset(0))

		vao.Unbind()
		shader.Disable()

		// glfw update
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

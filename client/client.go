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

const (
	WireframeVertexShader   = "shaders/wire.vert.glsl"
	WireframeFragmentShader = "shaders/wire.frag.glsl"

	WorldVertexShader   = "shaders/world.vert.glsl"
	WorldFragmentShader = "shaders/world.frag.glsl"
)

// ----------------------------------------------------------------------------
func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

// ----------------------------------------------------------------------------
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

// ----------------------------------------------------------------------------
func setupOpenGL() {
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)
}

// ----------------------------------------------------------------------------
func createCube() *gocraft.Vao {
	// cube positions
	verts := []float32{
		// front
		-0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,
		0.5, 0.5, 0.5,
		-0.5, 0.5, 0.5,
		// back
		-0.5, -0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, 0.5, -0.5,
		-0.5, 0.5, -0.5,
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
	vao.Load(verts, indices, uvs, normals)
	return vao
}

// ----------------------------------------------------------------------------
func createShaders() (*gocraft.Shader, *gocraft.Shader) {
	// world shader
	attribs := []gocraft.VertexAttribute{
		gocraft.VertexAttribute{Position: gocraft.AttribIndexPositions, Name: "a_pos"},
		gocraft.VertexAttribute{Position: gocraft.AttribIndexUvs, Name: "a_uvs"},
		gocraft.VertexAttribute{Position: gocraft.AttribIndexNormals, Name: "a_norm"},
	}
	worldShader, err := gocraft.NewShader(WorldVertexShader, WorldFragmentShader, attribs)
	if err != nil {
		panic(err)
	}

	// wireframe shader
	attribs = []gocraft.VertexAttribute{
		gocraft.VertexAttribute{Position: gocraft.AttribIndexPositions, Name: "a_pos"},
	}
	wireShader, err := gocraft.NewShader(WireframeVertexShader, WireframeFragmentShader, attribs)
	if err != nil {
		panic(err)
	}

	return worldShader, wireShader
}

// ----------------------------------------------------------------------------
func createBlockTypes() map[uint8]*gocraft.BlockType {
	defs := make(map[uint8]*gocraft.BlockType)
	defs[0x01] = &gocraft.BlockType{Color: gocraft.ColorRed.Copy()}   // red
	defs[0x02] = &gocraft.BlockType{Color: gocraft.ColorGreen.Copy()} // green
	defs[0x03] = &gocraft.BlockType{Color: gocraft.ColorBlue.Copy()}  // blue
	defs[0x04] = &gocraft.BlockType{Color: gocraft.ColorTeal.Copy()}  // teal
	return defs
}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------
func main() {
	// setup glfw
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	window := setupWindow()
	setupOpenGL()

	// shaders
	worldShader, wireShader := createShaders()
	defer worldShader.Dispose()
	defer wireShader.Dispose()

	// block definitions
	blocks := createBlockTypes()

	// cube mesh
	cube := createCube()
	defer cube.Dispose()
	model := glm.NewMat4(true)
	model.Translation(0, 0.5, -5)

	// camera
	ratio := float32(windowWidth) / float32(windowHeight)
	cam := gocraft.NewCamera(70, ratio, 0.01, 1000)

	// uniforms
	worldMvpUniform := gl.GetUniformLocation(worldShader.ID, gl.Str("u_mvp\x00"))
	worldColorUniform := gl.GetUniformLocation(worldShader.ID, gl.Str("u_color\x00"))
	wireMvpUniform := gl.GetUniformLocation(wireShader.ID, gl.Str("u_mvp\x00"))
	mvp := glm.NewMat4(true)

	// voxel data
	voxel := gocraft.Block(0x04) // 0x04 -> teal

	// game loop
	gl.Enable(gl.DEPTH_TEST)
	for !window.ShouldClose() {
		// clear window
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.ClearColor(0.95, 0.95, 0.95, 0.0)

		model.Rotate(2, 0, -1, 0)
		cam.Update()
		mvp.Set(cam.Combined.Data)
		mvp.Mul(model)

		cube.Bind()

		// draw solid
		color := blocks[voxel.BlockType()].Color
		worldShader.Enable()
		gl.UniformMatrix4fv(worldMvpUniform, 1, false, &mvp.Data[0])
		gl.Uniform3f(worldColorUniform, color.R, color.G, color.B)
		gl.DrawElements(gl.TRIANGLES, cube.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
		worldShader.Disable()

		// draw wireframe
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		wireShader.Enable()
		gl.UniformMatrix4fv(wireMvpUniform, 1, false, &mvp.Data[0])
		gl.DrawElements(gl.TRIANGLES, cube.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
		wireShader.Disable()
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

		cube.Unbind()

		// glfw update
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

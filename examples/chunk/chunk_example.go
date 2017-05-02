package main

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/mbrlabs/gocraft"
	"github.com/mbrlabs/gocraft/glm"
)

const (
	windowTitle  = "Chunk example"
	windowWidth  = 1024
	windowHeight = 768
)

const (
	WireframeVertexShader   = "../shaders/wire.vert"
	WireframeFragmentShader = "../shaders/wire.frag"

	WorldVertexShader   = "../shaders/world.vert"
	WorldFragmentShader = "../shaders/world.frag"
)

// ----------------------------------------------------------------------------
func createChunkMesh() *gocraft.Vao {
	mesher := gocraft.StupidMesher{}
	chunk := gocraft.NewChunk()
	mesh := mesher.Generate(chunk)

	vao := gocraft.NewVao()
	vao.Load(mesh.Positions, mesh.Indices, []float32{1, 2})
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

type ChunkDemo struct {
	blockTypes map[uint8]*gocraft.BlockType

	modelMatrix *glm.Mat4
	camera      *gocraft.Camera
	chunk       *gocraft.Vao

	mvp *glm.Mat4

	worldShader     *gocraft.Shader
	wireShader      *gocraft.Shader
	worldMvpUniform int32
	wireMvpUniform  int32
}

func (d *ChunkDemo) Create() {
	d.worldShader, d.wireShader = createShaders()
	d.blockTypes = createBlockTypes()
	d.chunk = createChunkMesh()

	d.worldMvpUniform = gl.GetUniformLocation(d.worldShader.ID, gl.Str("u_mvp\x00"))
	d.wireMvpUniform = gl.GetUniformLocation(d.wireShader.ID, gl.Str("u_mvp\x00"))

	ratio := float32(windowWidth) / float32(windowHeight)
	d.camera = gocraft.NewCamera(70, ratio, 0.01, 1000)

	d.modelMatrix = glm.NewMat4(true)
	d.modelMatrix.Translation(0, 0, -50)

	d.mvp = glm.NewMat4(true)

	gl.Enable(gl.DEPTH_TEST)
}

func (d *ChunkDemo) Dispose() {
	d.wireShader.Dispose()
	d.worldShader.Dispose()
	d.chunk.Dispose()
}

func (d *ChunkDemo) Update(delta float32) {
	d.modelMatrix.Rotate(2, 0, -1, 0)
	d.camera.Update()
	d.mvp.Set(d.camera.Combined.Data)
	d.mvp.Mul(d.modelMatrix)
}

func (d *ChunkDemo) Render(delta float32) {
	// clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.95, 0.95, 0.95, 0.0)

	d.chunk.Bind()

	// draw solid
	if true {
		d.worldShader.Enable()
		gl.UniformMatrix4fv(d.worldMvpUniform, 1, false, &d.mvp.Data[0])
		gl.DrawElements(gl.TRIANGLES, d.chunk.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
		d.worldShader.Disable()
	}

	// draw wireframe
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	d.wireShader.Enable()
	gl.UniformMatrix4fv(d.wireMvpUniform, 1, false, &d.mvp.Data[0])
	gl.DrawElements(gl.TRIANGLES, d.chunk.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
	d.wireShader.Disable()
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	d.chunk.Unbind()
}

func (d *ChunkDemo) Resize(width, height int) {

}

// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

func init() {
	runtime.LockOSThread()
}

func main() {
	window := gocraft.NewWindow(&gocraft.WindowConfig{
		Height:     windowHeight,
		Width:      windowWidth,
		Title:      windowTitle,
		Resizable:  false,
		Fullscreen: false,
		Vsync:      true,
	})

	window.Start(&ChunkDemo{})
}

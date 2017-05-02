package main

import (
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/mbrlabs/vox"
	"github.com/mbrlabs/vox/glm"
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
func createChunkMesh() *vox.Vao {
	mesher := vox.StupidMesher{}
	chunk := vox.NewChunk()
	mesh := mesher.Generate(chunk)

	vao := vox.NewVao()
	vao.Load(mesh.Positions, mesh.Indices, []float32{1, 2})
	return vao
}

// ----------------------------------------------------------------------------
func createShaders() (*vox.Shader, *vox.Shader) {
	// world shader
	attribs := []vox.VertexAttribute{
		vox.VertexAttribute{Position: vox.AttribIndexPositions, Name: "a_pos"},
		vox.VertexAttribute{Position: vox.AttribIndexUvs, Name: "a_uvs"},
		vox.VertexAttribute{Position: vox.AttribIndexNormals, Name: "a_norm"},
	}
	worldShader, err := vox.NewShader(WorldVertexShader, WorldFragmentShader, attribs)
	if err != nil {
		panic(err)
	}

	// wireframe shader
	attribs = []vox.VertexAttribute{
		vox.VertexAttribute{Position: vox.AttribIndexPositions, Name: "a_pos"},
	}
	wireShader, err := vox.NewShader(WireframeVertexShader, WireframeFragmentShader, attribs)
	if err != nil {
		panic(err)
	}

	return worldShader, wireShader
}

// ----------------------------------------------------------------------------
func createBlockTypes() map[uint8]*vox.BlockType {
	defs := make(map[uint8]*vox.BlockType)
	defs[0x01] = &vox.BlockType{Color: vox.ColorRed.Copy()}   // red
	defs[0x02] = &vox.BlockType{Color: vox.ColorGreen.Copy()} // green
	defs[0x03] = &vox.BlockType{Color: vox.ColorBlue.Copy()}  // blue
	defs[0x04] = &vox.BlockType{Color: vox.ColorTeal.Copy()}  // teal
	return defs
}

type ChunkDemo struct {
	blockTypes map[uint8]*vox.BlockType

	modelMatrix *glm.Mat4
	camera      *vox.Camera
	chunk       *vox.Vao

	mvp *glm.Mat4

	worldShader     *vox.Shader
	wireShader      *vox.Shader
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
	d.camera = vox.NewCamera(70, ratio, 0.01, 1000)

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

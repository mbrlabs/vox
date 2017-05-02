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

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/mbrlabs/vox"
	"github.com/mbrlabs/vox/glm"
)

const (
	windowTitle  = "Cube example"
	windowWidth  = 1024
	windowHeight = 768
)

const (
	WireframeVertexShader   = "shaders/wire.vert"
	WireframeFragmentShader = "shaders/wire.frag"

	WorldVertexShader   = "shaders/world.vert"
	WorldFragmentShader = "shaders/world.frag"
)

// ----------------------------------------------------------------------------
func createCube() *vox.Vao {
	// cube positions
	verts := []float32{
		// front
		-0.5, -0.5, 0.5, // 0
		0.5, -0.5, 0.5, // 1
		0.5, 0.5, 0.5, // 2
		-0.5, 0.5, 0.5, // 3
		// back
		-0.5, -0.5, -0.5, // 4
		0.5, -0.5, -0.5, // 5
		0.5, 0.5, -0.5, // 6
		-0.5, 0.5, -0.5, // 7
	}
	// cube indices
	indices := []uint16{
		// front
		0, 1, 2,
		2, 3, 0,
		// back
		5, 4, 7,
		7, 6, 5,
		// top
		3, 2, 6,
		6, 7, 3,
		// bottom
		0, 1, 5,
		5, 4, 0,
		// left
		4, 0, 3,
		3, 7, 4,
		// right
		1, 5, 6,
		6, 2, 1,
	}
	normals := []float32{1, 2}

	vao := vox.NewVao()
	vao.Load(verts, indices, normals)
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

type CubeDemo struct {
	modelMatrix *glm.Mat4
	camera      *vox.Camera
	cube        *vox.Vao

	mvp *glm.Mat4

	worldShader     *vox.Shader
	wireShader      *vox.Shader
	worldMvpUniform int32
	wireMvpUniform  int32
}

func (d *CubeDemo) Create() {
	d.worldShader, d.wireShader = createShaders()
	d.cube = createCube()

	d.worldMvpUniform = gl.GetUniformLocation(d.worldShader.ID, gl.Str("u_mvp\x00"))
	d.wireMvpUniform = gl.GetUniformLocation(d.wireShader.ID, gl.Str("u_mvp\x00"))

	ratio := float32(windowWidth) / float32(windowHeight)
	d.camera = vox.NewCamera(70, ratio, 0.01, 1000)

	d.modelMatrix = glm.NewMat4(true)
	d.modelMatrix.Translation(0, 0.5, -5)

	d.mvp = glm.NewMat4(true)

	gl.Enable(gl.DEPTH_TEST)
}

func (d *CubeDemo) Dispose() {
	d.wireShader.Dispose()
	d.worldShader.Dispose()
	d.cube.Dispose()
}

func (d *CubeDemo) Update(delta float32) {
	d.modelMatrix.Rotate(2, 0, -1, 0)
	d.camera.Update()

	d.mvp.Set(d.camera.Combined.Data)
	d.mvp.Mul(d.modelMatrix)
}

func (d *CubeDemo) Render(delta float32) {
	// clear window
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.ClearColor(0.95, 0.95, 0.95, 0.0)

	d.cube.Bind()

	// draw solid
	d.worldShader.Enable()
	gl.UniformMatrix4fv(d.worldMvpUniform, 1, false, &d.mvp.Data[0])
	gl.DrawElements(gl.TRIANGLES, d.cube.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
	d.worldShader.Disable()

	// draw wireframe
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
	d.wireShader.Enable()
	gl.UniformMatrix4fv(d.wireMvpUniform, 1, false, &d.mvp.Data[0])
	gl.DrawElements(gl.TRIANGLES, d.cube.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
	d.wireShader.Disable()
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	d.cube.Unbind()
}

func (d *CubeDemo) Resize(width, height int) {

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

	window.Start(&CubeDemo{})
}

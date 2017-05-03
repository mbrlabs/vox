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

import "github.com/go-gl/gl/v3.3-core/gl"

const (
	AttribIndexPositions = 0
	//AttribIndexUvs       = 1
	AttribIndexColor   = 1
	AttribIndexNormals = 2
)

type Vao struct {
	id             uint32
	positionBuffer uint32
	indexBuffer    uint32
	colorBuffer    uint32
	normalBuffer   uint32

	IndexCount int32
}

func NewVao() *Vao {
	vao := &Vao{}
	gl.GenVertexArrays(1, &vao.id)
	gl.GenBuffers(1, &vao.positionBuffer)
	gl.GenBuffers(1, &vao.indexBuffer)
	gl.GenBuffers(1, &vao.colorBuffer)
	gl.GenBuffers(1, &vao.normalBuffer)

	return vao
}

func (v *Vao) Load(positions []float32, indices []uint16, colors []float32, normals []float32) {
	gl.BindVertexArray(v.id)

	// indices
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.indexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, gl.Ptr(indices), gl.STATIC_DRAW)

	// positions
	gl.BindBuffer(gl.ARRAY_BUFFER, v.positionBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions)*4, gl.Ptr(positions), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexPositions, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	// colors
	gl.BindBuffer(gl.ARRAY_BUFFER, v.colorBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(colors)*4, gl.Ptr(colors), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexColor, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	// uvs
	//gl.BindBuffer(gl.ARRAY_BUFFER, v.uvBuffer)
	//gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*4, gl.Ptr(uvs), gl.STATIC_DRAW)
	//gl.VertexAttribPointer(AttribIndexUvs, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	// normals
	gl.BindBuffer(gl.ARRAY_BUFFER, v.normalBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(normals)*4, gl.Ptr(normals), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexNormals, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	v.IndexCount = int32(len(indices))
}

func (v *Vao) Bind() {
	gl.BindVertexArray(v.id)
	gl.EnableVertexAttribArray(AttribIndexPositions)
	gl.EnableVertexAttribArray(AttribIndexColor)
	//gl.EnableVertexAttribArray(AttribIndexUvs)
	gl.EnableVertexAttribArray(AttribIndexNormals)
}

func (v *Vao) Unbind() {
	gl.DisableVertexAttribArray(AttribIndexNormals)
	gl.EnableVertexAttribArray(AttribIndexColor)
	//gl.DisableVertexAttribArray(AttribIndexUvs)
	gl.DisableVertexAttribArray(AttribIndexPositions)
	gl.BindVertexArray(0)
}

func (v *Vao) Dispose() {
	// TODO delete buffers as well?
	gl.DeleteVertexArrays(1, &v.id)
}

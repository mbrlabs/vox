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
	"github.com/go-gl/gl/v3.3-core/gl"
)

const (
	AttribIndexPositions = 0
	AttribIndexColor     = 1
	AttribIndexNormals   = 2
)

var (
	mustCreateChunkIndexBuffer bool = true
	cunckIndexBuffer           uint32
)

func createChunkIndexBuffer() {
	// allocate max possible buffer for the given chunk size
	var verts uint16
	faces := ChunkXYZ * 6

	indices := make([]uint16, 0)
	for i := 0; i < faces; i++ {
		verts += 4
		indices = append(indices,
			verts-4, verts-3, verts-2,
			verts-2, verts-1, verts-4,
		)
	}

	gl.GenBuffers(1, &cunckIndexBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, cunckIndexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, gl.Ptr(indices), gl.STATIC_DRAW)
}

type MeshData struct {
	Positions  []float32
	Colors     []float32
	IndexCount int
}

type Mesh struct {
	vao            uint32
	positionBuffer uint32
	colorBuffer    uint32

	IndexCount int32
}

func NewMesh() *Mesh {
	mesh := &Mesh{}
	gl.GenVertexArrays(1, &mesh.vao)
	gl.GenBuffers(1, &mesh.positionBuffer)
	gl.GenBuffers(1, &mesh.colorBuffer)

	return mesh
}

func (m *Mesh) Load(data *MeshData) {
	// generate global index buffer if not already done
	if mustCreateChunkIndexBuffer {
		createChunkIndexBuffer()
		mustCreateChunkIndexBuffer = false
	}

	positions := data.Positions
	colors := data.Colors

	gl.BindVertexArray(m.vao)

	// bind global index buffer
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, cunckIndexBuffer)

	// positions

	gl.BindBuffer(gl.ARRAY_BUFFER, m.positionBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions)*4, gl.Ptr(positions), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexPositions, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	// colors
	gl.BindBuffer(gl.ARRAY_BUFFER, m.colorBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(colors)*4, gl.Ptr(colors), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexColor, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	m.IndexCount = int32(len(data.Positions))
}

func (m *Mesh) Bind() {
	gl.BindVertexArray(m.vao)
	gl.EnableVertexAttribArray(AttribIndexPositions)
	gl.EnableVertexAttribArray(AttribIndexColor)
}

func (m *Mesh) Unbind() {
	gl.EnableVertexAttribArray(AttribIndexColor)
	gl.DisableVertexAttribArray(AttribIndexPositions)
	gl.BindVertexArray(0)
}

func (m *Mesh) Dispose() {
	// TODO delete buffers as well?
	gl.DeleteVertexArrays(1, &m.vao)
}

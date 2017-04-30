package gocraft

import "github.com/go-gl/gl/v3.3-core/gl"

const (
	AttribIndexPositions = 0
	AttribIndexUvs       = 1
	AttribIndexNormals   = 2
)

type Vao struct {
	id             uint32
	positionBuffer uint32
	indexBuffer    uint32
	uvBuffer       uint32
	normalBuffer   uint32
}

func NewVao() *Vao {
	vao := &Vao{}
	gl.GenVertexArrays(1, &vao.id)
	gl.GenBuffers(1, &vao.positionBuffer)
	gl.GenBuffers(1, &vao.indexBuffer)
	gl.GenBuffers(1, &vao.uvBuffer)
	gl.GenBuffers(1, &vao.normalBuffer)

	return vao
}

func (v *Vao) Load(positions []float32, indices []uint16, uvs []float32, normals []float32) {
	gl.BindVertexArray(v.id)

	// indices
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, v.indexBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*2, gl.Ptr(indices), gl.STATIC_DRAW)

	// positions
	gl.BindBuffer(gl.ARRAY_BUFFER, v.positionBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(positions)*4, gl.Ptr(positions), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexPositions, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	// uvs
	gl.BindBuffer(gl.ARRAY_BUFFER, v.uvBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*4, gl.Ptr(uvs), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexUvs, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	// normals
	gl.BindBuffer(gl.ARRAY_BUFFER, v.normalBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, len(uvs)*4, gl.Ptr(normals), gl.STATIC_DRAW)
	gl.VertexAttribPointer(AttribIndexNormals, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func (v *Vao) Dispose() {
	// TODO delete buffers as well?
	gl.DeleteVertexArrays(1, &v.id)
}

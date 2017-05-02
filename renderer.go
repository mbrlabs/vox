package vox

import "github.com/go-gl/gl/v3.3-core/gl"

type WorldRenderer struct {
	Disposable

	solidShader     *Shader
	uniformSolidMvp int32

	wireShader     *Shader
	uniformWireMvp int32
}

func NewWorldRenderer() *WorldRenderer {
	// solid shader
	attribs := []VertexAttribute{
		VertexAttribute{Position: AttribIndexPositions, Name: "a_pos"},
		VertexAttribute{Position: AttribIndexUvs, Name: "a_uvs"},
		VertexAttribute{Position: AttribIndexNormals, Name: "a_norm"},
	}
	ss, err := NewShader(WorldVertexShader, WorldFragmentShader, attribs)
	if err != nil {
		panic(err)
	}

	// wireframe shader
	attribs = []VertexAttribute{
		VertexAttribute{Position: AttribIndexPositions, Name: "a_pos"},
	}
	ws, err := NewShader(WireframeVertexShader, WireframeFragmentShader, attribs)
	if err != nil {
		panic(err)
	}

	return &WorldRenderer{
		solidShader:     ss,
		wireShader:      ws,
		uniformSolidMvp: gl.GetUniformLocation(ss.ID, gl.Str("u_mvp\x00")),
		uniformWireMvp:  gl.GetUniformLocation(ws.ID, gl.Str("u_mvp\x00")),
	}
}

func (r *WorldRenderer) Dispose() {
	r.solidShader.Dispose()
	r.wireShader.Dispose()
}

func (r *WorldRenderer) Render(cam *Camera, world *World) {
	//bank := world.BlockBank

	for _, chunk := range world.Chunks {
		chunk.Mesh.Bind()

		// solid render pass
		r.solidShader.Enable()
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		gl.UniformMatrix4fv(r.uniformSolidMvp, 1, false, &cam.Combined.Data[0])
		gl.DrawElements(gl.TRIANGLES, chunk.Mesh.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))

		// wireframe render
		r.wireShader.Enable()
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		gl.UniformMatrix4fv(r.uniformWireMvp, 1, false, &cam.Combined.Data[0])
		gl.DrawElements(gl.TRIANGLES, chunk.Mesh.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
	}
}

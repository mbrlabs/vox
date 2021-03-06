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

const worldVert = `
#version 330 

uniform mat4 u_mvp;
uniform vec3 u_sun_direction;

in vec3 a_pos;
in vec3 a_norm;
in vec2 a_uv;

out vec2 texCoords;
out float diffuse;

void main() {
	texCoords = a_uv;
	diffuse = dot(a_norm, u_sun_direction);
	diffuse = clamp(diffuse, 0.2, 1);
    gl_Position = u_mvp * vec4(a_pos, 1.0);
}
`

const worldFrag = `
#version 330

uniform sampler2D tex;
uniform vec3 u_sun_color;

in vec2 texCoords;
in float diffuse;

out vec4 outColor;

void main() {
	vec4 light = vec4(diffuse * u_sun_color, 1);
	outColor = light * texture(tex, texCoords);
}
`

const wireVert = `
#version 330 

const float SCALE_FACTOR = 1.003;

uniform mat4 u_mvp;

in vec3 a_pos;

void main() {
    vec3 scaledVertex = a_pos * SCALE_FACTOR;
    gl_Position = u_mvp * vec4(scaledVertex, 1.0);
}
`

const wireFrag = `
#version 330

out vec4 outColor;

void main() {
    outColor = vec4(0.0, 0.0, 0.0, 1.0);
}
`

type WorldRenderer struct {
	Disposable

	solidShader         *Shader
	uniformSolidMvp     int32
	uniformSunDirection int32
	uniformSunColor     int32

	wireShader     *Shader
	uniformWireMvp int32
}

func NewWorldRenderer() *WorldRenderer {
	// solid shader
	attribs := []VertexAttribute{
		{Position: AttribIndexPositions, Name: "a_pos"},
		{Position: AttribIndexUvs, Name: "a_uv"},
		{Position: AttribIndexNormals, Name: "a_norm"},
	}
	ss, err := NewShader(worldVert, worldFrag, attribs)
	if err != nil {
		panic(err)
	}

	// wireframe shader
	attribs = []VertexAttribute{
		VertexAttribute{Position: AttribIndexPositions, Name: "a_pos"},
	}
	ws, err := NewShader(wireVert, wireFrag, attribs)
	if err != nil {
		panic(err)
	}

	return &WorldRenderer{
		solidShader:         ss,
		wireShader:          ws,
		uniformWireMvp:      gl.GetUniformLocation(ws.ID, gl.Str("u_mvp\x00")),
		uniformSolidMvp:     gl.GetUniformLocation(ss.ID, gl.Str("u_mvp\x00")),
		uniformSunDirection: gl.GetUniformLocation(ss.ID, gl.Str("u_sun_direction\x00")),
		uniformSunColor:     gl.GetUniformLocation(ss.ID, gl.Str("u_sun_color\x00")),
	}
}

func (r *WorldRenderer) Dispose() {
	r.solidShader.Dispose()
	r.wireShader.Dispose()
}

func (r *WorldRenderer) Render(cam *Camera, world *World, env *Environment) {
	for _, chunk := range world.Chunks {
		// can happen if chunk is completly sourrounded by other chunks and not a single triange would be drawn
		if chunk.Mesh == nil {
			continue
		}

		sunDir := env.Sun.Direction
		sunColor := env.Sun.Color

		chunk.Mesh.Bind()

		// solid render pass
		r.solidShader.Enable()
		gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)
		gl.UniformMatrix4fv(r.uniformSolidMvp, 1, false, &cam.Combined.Data[0])
		gl.UniformMatrix4fv(r.uniformSolidMvp, 1, false, &cam.Combined.Data[0])
		gl.Uniform3f(r.uniformSunColor, sunColor.R, sunColor.G, sunColor.B)
		gl.Uniform3f(r.uniformSunDirection, sunDir.X, sunDir.Y, sunDir.Z)

		gl.DrawElements(gl.TRIANGLES, chunk.Mesh.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))

		// wireframe render
		// r.wireShader.Enable()
		// gl.PolygonMode(gl.FRONT_AND_BACK, gl.LINE)
		// gl.UniformMatrix4fv(r.uniformWireMvp, 1, false, &cam.Combined.Data[0])
		// gl.DrawElements(gl.TRIANGLES, chunk.Mesh.IndexCount, gl.UNSIGNED_SHORT, gl.PtrOffset(0))
	}
}

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
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

const (
	WireframeVertexShader   = "shaders/wire.vert"
	WireframeFragmentShader = "shaders/wire.frag"

	WorldVertexShader   = "shaders/world.vert"
	WorldFragmentShader = "shaders/world.frag"
)

type VertexAttribute struct {
	Position uint32
	Name     string
}

type Shader struct {
	ID       uint32
	vertPath string
	fragPath string
}

func compileShader(path string, shaderType uint32) (uint32, error) {
	// read sources
	source, err := ioutil.ReadFile(path)
	if err != nil {
		return 0, err
	}

	shader := gl.CreateShader(shaderType)
	csources, free := gl.Strs(string(source) + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	// check for compile errors
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("Failed to compile %v: %v", path, log)
	}

	return shader, nil
}

func NewShader(vertexPath, fragmentPath string, attribs []VertexAttribute) (*Shader, error) {
	// compile vertex shader
	vertexShader, err := compileShader(vertexPath, gl.VERTEX_SHADER)
	if err != nil {
		return nil, err
	}

	// compile fragment shader
	fragmentShader, err := compileShader(fragmentPath, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, err
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	// bind vbos
	for _, attrib := range attribs {
		gl.BindAttribLocation(program, attrib.Position, gl.Str(attrib.Name+"\x00"))
	}

	gl.LinkProgram(program)
	//gl.ValidateProgram(program)

	// check for link errors
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("Failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Shader{
		ID:       program,
		vertPath: vertexPath,
		fragPath: vertexPath,
	}, nil
}

func (s *Shader) Enable() {
	gl.UseProgram(s.ID)
}

func (s *Shader) Disable() {
	gl.UseProgram(0)
}

func (s *Shader) Dispose() {
	s.Disable()
	gl.DeleteProgram(s.ID)
}

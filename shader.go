package vox

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
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

	//gl.BindAttribLocation(program, AttribIndexPositions, gl.Str("a_pos\x00"))
	//gl.BindAttribLocation(program, AttribIndexUvs, gl.Str("a_uvs\x00"))
	//gl.BindAttribLocation(program, AttribIndexNormals, gl.Str("a_norm\x00"))

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

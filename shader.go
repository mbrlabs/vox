package gocraft

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

type Shader struct {
	id       uint32
	vertPath string
	fragPath string
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
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

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

func NewShader(vertexPath, fragmentPath string) *Shader {
	// read vertex source
	vertexSource, err := ioutil.ReadFile(vertexPath)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	// read fragment source
	fragmentSource, err := ioutil.ReadFile(fragmentPath)
	if err != nil {
		fmt.Print(err)
		return nil
	}

	// compile vertex shader
	vertexShader, err := compileShader(string(vertexSource), gl.VERTEX_SHADER)
	if err != nil {
		return nil
	}

	// compile fragment shader
	fragmentShader, err := compileShader(string(fragmentSource), gl.FRAGMENT_SHADER)
	if err != nil {
		return nil
	}

	// create program
	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	// check for link errors
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		fmt.Printf("failed to link program: %v", log)
		return nil
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Shader{
		id:       program,
		vertPath: vertexPath,
		fragPath: vertexPath,
	}
}

func (s *Shader) Enable() {
	gl.UseProgram(s.id)
}

func (s *Shader) Disable() {
	gl.UseProgram(0)
}

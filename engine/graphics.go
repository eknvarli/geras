package engine

import (
	"fmt"

	"github.com/go-gl/gl/v3.3-core/gl"
)

var (
	quadVAO       uint32
	quadVBO       uint32
	shaderProgram uint32
	colorUniform  int32
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

var winWidth, winHeight int

func InitGraphics(width, height int) {
	winWidth = width
	winHeight = height

	vertexShaderSource := `#version 330 core
layout(location = 0) in vec2 aPos;
void main() {
    gl_Position = vec4(aPos, 0.0, 1.0);
}`
	fragmentShaderSource := `#version 330 core
out vec4 FragColor;
uniform vec3 uColor;
void main() {
    FragColor = vec4(uColor, 1.0);
}`

	shaderProgram, _ = newShaderProgram(vertexShaderSource, fragmentShaderSource)
	colorUniform = gl.GetUniformLocation(shaderProgram, gl.Str("uColor\x00"))

	gl.GenVertexArrays(1, &quadVAO)
	gl.GenBuffers(1, &quadVBO)
	gl.BindVertexArray(quadVAO)
	gl.BindBuffer(gl.ARRAY_BUFFER, quadVBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*2*4, nil, gl.DYNAMIC_DRAW)
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 2, gl.FLOAT, false, 2*4, nil)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func DrawQuad(x, y, w, h float32, color [3]float32) {
	gl.UseProgram(shaderProgram)
	gl.Uniform3f(colorUniform, color[0], color[1], color[2])

	sx := 2.0 / float32(winWidth)
	sy := 2.0 / float32(winHeight)

	verts := []float32{
		x*sx - 1, 1 - y*sy,
		(x+w)*sx - 1, 1 - y*sy,
		(x+w)*sx - 1, 1 - (y+h)*sy,
		x*sx - 1, 1 - (y+h)*sy,
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, quadVBO)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, len(verts)*4, gl.Ptr(verts))
	gl.BindVertexArray(quadVAO)
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, 4)
	gl.BindVertexArray(0)
}

func newShaderProgram(vertexSrc, fragmentSrc string) (uint32, error) {
	vertexShader, err := compileShader(vertexSrc, gl.VERTEX_SHADER)
	if err != nil {
		return 0, err
	}
	fragmentShader, err := compileShader(fragmentSrc, gl.FRAGMENT_SHADER)
	if err != nil {
		return 0, err
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetProgramInfoLog(program, logLength, nil, &log[0])
		return 0, fmt.Errorf("failed to link program: %s", log)
	}
	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return program, nil
}

func compileShader(src string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	sources, free := gl.Strs(src + "\x00")
	gl.ShaderSource(shader, 1, sources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := make([]byte, logLength)
		gl.GetShaderInfoLog(shader, logLength, nil, &log[0])
		return 0, fmt.Errorf("failed to compile shader: %s", log)
	}
	return shader, nil
}

func DrawText(x, y float32, text string, color [3]float32) {
	for i := 0; i < len(text); i++ {
		cx := x + float32(i)*12
		cy := y
		DrawQuad(cx, cy, 10, 14, color)
	}
}

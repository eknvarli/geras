package engine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func beginRender() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func endRender() {
	window.SwapBuffers()
}

// TODO: add simple sprite shader, quad VBO, texture unit management

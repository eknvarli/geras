package engine

import (
	"github.com/go-gl/gl/v3.3-core/gl"
)

func initGL() error {
	if err := gl.Init(); err != nil {
		return err
	}
	gl.ClearColor(0.1, 0.1, 0.12, 1.0)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	return nil
}

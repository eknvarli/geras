package engine

import (
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

var window *glfw.Window

func init() {
	runtime.LockOSThread()
}

func initWindow(w, h int, title string) error {
	if err := glfw.Init(); err != nil {
		return err
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)

	win, err := glfw.CreateWindow(w, h, title, nil, nil)
	if err != nil {
		return err
	}
	window = win
	window.MakeContextCurrent()
	return nil
}

func terminateWindow() {
	if window != nil {
		window.Destroy()
	}
	glfw.Terminate()
}

func WindowShouldClose() bool { return window.ShouldClose() }

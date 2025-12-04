package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	keyState       = make(map[glfw.Key]glfw.Action)
	mouseX, mouseY float64
	mouseButtons   = make(map[glfw.MouseButton]glfw.Action)
)

func pollInput() {
	for k := glfw.KeySpace; k <= glfw.KeyLast; k++ {
		keyState[k] = window.GetKey(k)
	}

	mouseX, mouseY = window.GetCursorPos()
	for b := glfw.MouseButtonLeft; b <= glfw.MouseButtonLast; b++ {
		mouseButtons[b] = window.GetMouseButton(b)
	}
}

func IsKeyDown(k glfw.Key) bool           { return keyState[k] == glfw.Press }
func GetMousePos() (float64, float64)     { return mouseX, mouseY }
func IsMouseDown(b glfw.MouseButton) bool { return mouseButtons[b] == glfw.Press }

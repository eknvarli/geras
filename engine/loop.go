package engine

import (
	"time"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func gameLoop() {
	last := time.Now()
	for !WindowShouldClose() {
		now := time.Now()
		dt := float32(now.Sub(last).Seconds())
		last = now

		pollInput()
		if currentGame != nil {
			currentGame.Update(dt)
		}

		beginRender()
		if currentGame != nil {
			currentGame.Draw()
		}
		endRender()

		glfw.PollEvents()
	}
}

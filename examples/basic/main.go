package main

import (
	"github.com/eknvarli/geras/engine"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type MyGame struct {
	x, y   float32
	vx, vy float32
	size   float32
}

func (g *MyGame) Update(dt float32) {
	// ok tuşları ile hareket
	speed := float32(200)
	if engine.IsKeyDown(glfw.KeyRight) {
		g.x += speed * dt
	}
	if engine.IsKeyDown(glfw.KeyLeft) {
		g.x -= speed * dt
	}
	if engine.IsKeyDown(glfw.KeyUp) {
		g.y -= speed * dt
	}
	if engine.IsKeyDown(glfw.KeyDown) {
		g.y += speed * dt
	}

	// pencere sınırları
	if g.x < 0 {
		g.x = 0
	}
	if g.y < 0 {
		g.y = 0
	}
	if g.x+g.size > 800 {
		g.x = 800 - g.size
	}
	if g.y+g.size > 600 {
		g.y = 600 - g.size
	}
}

func (g *MyGame) Draw() {
	// player kare
	engine.DrawQuad(g.x, g.y, g.size, g.size, [3]float32{1, 0, 0})
	// ekrana yazı
	engine.DrawText(300, 50, "HELLO GERAS", [3]float32{0, 1, 0})
}

func main() {
	game := &MyGame{x: 400, y: 300, size: 50}
	cfg := engine.Config{Width: 800, Height: 600, Title: "Geras Demo"}
	if err := engine.Run(game, cfg); err != nil {
		panic(err)
	}
}

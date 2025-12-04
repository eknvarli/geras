# Yet another Go game engine.

Geras is a lightweight, cross-platform 2D game engine written in Go. It provides basic rendering, input handling, sound support, and a simple game loop. Designed for simplicity, it allows developers to create 2D desktop games quickly.

## Features

* **Cross-platform**: Works on Windows, Linux, and macOS.
* **Rendering Engine**: OpenGL 3.3 core profile for 2D graphics.
* **Input Engine**: Keyboard and mouse support via GLFW.
* **Sound Engine**: Basic sound playback.
* **Game Loop**: Delta time-based update for smooth gameplay.
* **Simple Text Rendering**: Draw simple colored quads to simulate text.

## Installation

Clone the repository and use it as a Go module:

```bash
git clone https://github.com/eknvarli/geras.git
cd geras
```

In your game project:

```go
go get github.com/eknvarli/geras
```

## Usage

### Basic Example

```go
package main

import (
	"github.com/eknvarli/geras/engine"
	"github.com/go-gl/glfw/v3.3/glfw"
)

type MyGame struct {
	x, y float32
	vx, vy float32
	size float32
}

func (g *MyGame) Update(dt float32) {
	if engine.IsKeyDown(glfw.KeyRight) {
		g.x += 200 * dt
	}
	// handle other keys...
}

func (g *MyGame) Draw() {
	engine.DrawQuad(g.x, g.y, g.size, g.size, [3]float32{1, 0, 0})
	engine.DrawText(300, 50, "HELLO GERAS", [3]float32{0, 1, 0})
}

func main() {
	game := &MyGame{x: 400, y: 300, size: 50}
	cfg := engine.Config{Width: 800, Height: 600, Title: "Geras Demo"}
	engine.Run(game, cfg)
}
```

## Project Structure

```
geras/
├── engine/         # Core engine modules
│   ├── app.go      # Run, Config, termination
│   ├── graphics.go # Rendering, DrawQuad, DrawText
│   ├── input.go    # Input handling
│   ├── sound.go    # Sound engine
│   └── ...
├── examples/       # Sample games
│   └── basic/      # Simple moving square demo
├── go.mod
└── README.md       # This file
```

## Contributing

Feel free to fork, open issues, or submit pull requests to improve Geras.

## License

GPLv3.0 License

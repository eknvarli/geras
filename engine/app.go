package engine

import (
	"log"
)

type Game interface {
	Update(dt float32)
	Draw()
}

var currentGame Game

func Run(g Game, cfg Config) error {
	currentGame = g
	if err := initWindow(cfg.Width, cfg.Height, cfg.Title); err != nil {
		return err
	}
	if err := initGL(); err != nil {
		return err
	}

	InitGraphics(cfg.Width, cfg.Height)

	if err := SoundInit(); err != nil {
		log.Println("warning: sound init failed:", err)
	}
	gameLoop()
	Terminate()
	return nil
}

func Terminate() {
	SoundTerminate()
	terminateWindow()
}

type Config struct {
	Width  int
	Height int
	Title  string
}

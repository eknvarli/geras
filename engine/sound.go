package engine

import (
	"bytes"
	"io"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

var audioCtx *oto.Context

func SoundInit() error {
	var err error
	audioCtx, err = oto.NewContext(44100, 2, 2, 8192)
	return err
}

func PlayMP3Data(data []byte) error {
	if audioCtx == nil {
		return nil
	}

	r := bytes.NewReader(data)
	decoder, err := mp3.NewDecoder(r)
	if err != nil {
		return err
	}

	player := audioCtx.NewPlayer()
	go func() {
		_, _ = io.Copy(player, decoder)
		player.Close()
	}()

	return nil
}

func SoundTerminate() {
	if audioCtx != nil {
		audioCtx.Close()
	}
}

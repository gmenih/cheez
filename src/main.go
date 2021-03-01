package main

import (
	"gmenih341/cheez/src/cheez/game"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Cheez",
		Bounds: pixel.R(0, 0, 1024, 768),

		VSync: true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	c := game.NewGame(win)
	t := time.Now()

	for !win.Closed() {
		n := time.Now()

		c.Update(int(t.Sub(n)))

		c.Draw()

		win.Update()

		t = n
	}
}

func main() {
	pixelgl.Run(run)
}

package main

import (
	"time"

	"github.com/rinchsan/lifegame"
)

func main() {
	const (
		alive      = "@@"
		dead       = "__"
		height     = 50
		width      = 50
		aliveRatio = 0.3
		interval   = 500 * time.Millisecond
	)

	game := lifegame.NewGame(height, width, alive, dead, aliveRatio, interval)
	game.Start()
}

package lifegame

import (
	"time"
)

type Game struct {
	board          Board
	initAliveRatio float64
	interval       time.Duration
}

func NewGame(height, width int, aliveStr, deadStr string, initAliveRatio float64, interval time.Duration) Game {
	return Game{
		board:          NewBoard(height, width, aliveStr, deadStr),
		initAliveRatio: initAliveRatio,
		interval:       interval,
	}
}

func (g Game) Start() {
	g.Reset()
	g.Print()
	time.Sleep(g.interval)
	for {
		g.Next()
		g.Print()
		time.Sleep(g.interval)
	}
}

func (g *Game) Reset() {
	g.board.Reset(g.initAliveRatio)
}

func (g *Game) Next() {
	g.board = g.board.Next()
}

func (g Game) Print() {
	g.board.Print()
}

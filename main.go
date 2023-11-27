package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	const (
		markAlive        = "★ "
		markDead         = "**"
		boardHeight      = 50
		boardWidth       = 50
		aliveProbability = 0.1
		interval         = 500 * time.Millisecond
	)

	game := newGame(boardHeight, boardWidth, markAlive, markDead, aliveProbability, interval)
	game.start()
}

type game struct {
	board            board
	aliveProbability float64
	interval         time.Duration
	sep              string
}

func newGame(boardHeight, boardWidth int, markAlive, markDead string, aliveProbability float64, interval time.Duration) game {
	return game{
		board:            newBoard(boardHeight, boardWidth, markAlive, markDead),
		aliveProbability: aliveProbability,
		interval:         interval,
		sep:              strings.Repeat("-", len(markDead)),
	}
}

func (g game) start() {
	g.reset()
	g.print()
	time.Sleep(g.interval)
	for {
		g.next()
		g.print()
		time.Sleep(g.interval)
	}
}

func (g *game) reset() {
	g.board.reset(g.aliveProbability)
}

func (g *game) next() {
	g.board = g.board.next()
}

func (g game) print() {
	g.board.print()
	fmt.Println(strings.Repeat(g.sep, g.board.width()))
}

type board struct {
	lines     [][]block
	markAlive string
	markDead  string
}

func newBoard(height, width int, markAlive, markDead string) board {
	lines := make([][]block, height)
	for i := range lines {
		lines[i] = make([]block, width)
	}
	for i := range lines {
		for j := range lines[i] {
			lines[i][j] = newBlock(false)
		}
	}
	return board{
		lines:     lines,
		markAlive: markAlive,
		markDead:  markDead,
	}
}

func (b board) height() int {
	return len(b.lines)
}

func (b board) width() int {
	if len(b.lines) == 0 {
		return 0
	}
	return len(b.lines[0])
}

func (b board) next() board {
	next := newBoard(len(b.lines), len(b.lines[0]), b.markAlive, b.markDead)
	for i := range b.lines {
		for j := range b.lines[i] {
			aliveNext := b.aliveNext(i, j)
			next.setAlive(i, j, aliveNext)
		}
	}
	return next
}

func (b *board) setAlive(i, j int, alive bool) {
	b.lines[i][j].setAlive(alive)
}

func (b board) aliveNext(i, j int) bool {
	var aliveCount int
	if i != 0 && j != 0 {
		if b.lines[i-1][j-1].alive {
			aliveCount++
		}
	}
	if j != 0 {
		if b.lines[i][j-1].alive {
			aliveCount++
		}
	}
	if i != len(b.lines)-1 && j != 0 {
		if b.lines[i+1][j-1].alive {
			aliveCount++
		}
	}
	if i != 0 {
		if b.lines[i-1][j].alive {
			aliveCount++
		}
	}
	if i != len(b.lines)-1 {
		if b.lines[i+1][j].alive {
			aliveCount++
		}
	}
	if i != 0 && j != len(b.lines[i])-1 {
		if b.lines[i-1][j+1].alive {
			aliveCount++
		}
	}
	if j != len(b.lines[i])-1 {
		if b.lines[i][j+1].alive {
			aliveCount++
		}
	}
	if i != len(b.lines)-1 && j != len(b.lines[i])-1 {
		if b.lines[i+1][j+1].alive {
			aliveCount++
		}
	}
	return aliveCount == 3 || (b.lines[i][j].alive && aliveCount == 2)
}

func (b board) print() {
	for _, line := range b.lines {
		for _, block := range line {
			block.print(b.markAlive, b.markDead)
		}
		fmt.Println("")
	}
}

func (b *board) reset(aliveProbability float64) {
	for i := range b.lines {
		for j := range b.lines[i] {
			b.lines[i][j].reset(aliveProbability)
		}
	}
}

type block struct {
	alive bool
}

func newBlock(alive bool) block {
	return block{
		alive: alive,
	}
}

func (b *block) setAlive(alive bool) {
	b.alive = alive
}

func (b *block) reset(aliveProbability float64) {
	r := rand.Float64()
	b.alive = aliveProbability >= r
}

func (b block) print(markAlive, markDead string) {
	if b.alive {
		fmt.Print(markAlive)
	} else {
		fmt.Print(markDead)
	}
}

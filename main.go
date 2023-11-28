package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	const (
		markAlive        = "@@"
		markDead         = "__"
		boardHeight      = 50
		boardWidth       = 50
		aliveProbability = 0.2
		interval         = 500 * time.Millisecond
	)

	game := NewGame(boardHeight, boardWidth, markAlive, markDead, aliveProbability, interval)
	game.Start()
}

type Game struct {
	board            Board
	aliveProbability float64
	interval         time.Duration
}

func NewGame(boardHeight, boardWidth int, markAlive, markDead string, aliveProbability float64, interval time.Duration) Game {
	return Game{
		board:            NewBoard(boardHeight, boardWidth, markAlive, markDead),
		aliveProbability: aliveProbability,
		interval:         interval,
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
	g.board.Reset(g.aliveProbability)
}

func (g *Game) Next() {
	g.board = g.board.Next()
}

func (g Game) Print() {
	g.board.Print()
}

type Board struct {
	lines     [][]Block
	markAlive string
	markDead  string
}

func NewBoard(height, width int, markAlive, markDead string) Board {
	lines := make([][]Block, height)
	for i := range lines {
		lines[i] = make([]Block, width)
	}
	for i := range lines {
		for j := range lines[i] {
			lines[i][j] = NewBlock(false)
		}
	}
	return Board{
		lines:     lines,
		markAlive: markAlive,
		markDead:  markDead,
	}
}

func (b Board) Next() Board {
	next := NewBoard(len(b.lines), len(b.lines[0]), b.markAlive, b.markDead)
	for i := range b.lines {
		for j := range b.lines[i] {
			aliveNext := b.AliveNext(i, j)
			next.SetAlive(i, j, aliveNext)
		}
	}
	return next
}

func (b *Board) SetAlive(i, j int, alive bool) {
	b.lines[i][j].SetAlive(alive)
}

func (b Board) AliveNext(i, j int) bool {
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

func (b Board) Print() {
	if len(b.lines) == 0 {
		return
	}
	sep := strings.Repeat("-", len(b.markAlive))
	fmt.Println("*" + strings.Repeat(sep, len(b.lines[0])+1) + "*")
	for _, line := range b.lines {
		fmt.Print("| ")
		for _, block := range line {
			block.Print(b.markAlive, b.markDead)
		}
		fmt.Println(" |")
	}
	fmt.Println("*" + strings.Repeat(sep, len(b.lines[0])+1) + "*")
}

func (b *Board) Reset(aliveProbability float64) {
	for i := range b.lines {
		for j := range b.lines[i] {
			b.lines[i][j].Reset(aliveProbability)
		}
	}
}

type Block struct {
	alive bool
}

func NewBlock(alive bool) Block {
	return Block{
		alive: alive,
	}
}

func (b *Block) SetAlive(alive bool) {
	b.alive = alive
}

func (b *Block) Reset(aliveProbability float64) {
	r := rand.Float64()
	b.alive = aliveProbability >= r
}

func (b Block) Print(markAlive, markDead string) {
	if b.alive {
		fmt.Print(markAlive)
	} else {
		fmt.Print(markDead)
	}
}

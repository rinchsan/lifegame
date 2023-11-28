package lifegame

import (
	"fmt"
	"strings"
)

type Board struct {
	lines    [][]Block
	aliveStr string
	deadStr  string
	sep      string
}

func NewBoard(height, width int, aliveStr, deadStr string) Board {
	lines := make([][]Block, height)
	for i := range lines {
		lines[i] = make([]Block, width)
	}
	for i := range lines {
		for j := range lines[i] {
			lines[i][j] = NewBlock(false, aliveStr, deadStr)
		}
	}
	return Board{
		lines:    lines,
		aliveStr: aliveStr,
		deadStr:  deadStr,
		sep:      "*" + strings.Repeat(strings.Repeat("-", len(aliveStr)), width+1) + "*",
	}
}

func (b Board) Next() Board {
	if len(b.lines) == 0 {
		return Board{}
	}
	next := NewBoard(len(b.lines), len(b.lines[0]), b.aliveStr, b.deadStr)
	for i := range b.lines {
		for j := range b.lines[i] {
			alive := b.AliveNext(i, j)
			next.SetAlive(i, j, alive)
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
	fmt.Println(b.sep)
	for _, line := range b.lines {
		fmt.Print("| ")
		for _, block := range line {
			block.Print()
		}
		fmt.Println(" |")
	}
	fmt.Println(b.sep)
}

func (b *Board) Reset(initAliveRatio float64) {
	for i := range b.lines {
		for j := range b.lines[i] {
			b.lines[i][j].Reset(initAliveRatio)
		}
	}
}

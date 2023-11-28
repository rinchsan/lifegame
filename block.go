package lifegame

import (
	"fmt"
	"math/rand"
)

type Block struct {
	alive    bool
	aliveStr string
	deadStr  string
}

func NewBlock(alive bool, aliveStr, deadStr string) Block {
	return Block{
		alive:    alive,
		aliveStr: aliveStr,
		deadStr:  deadStr,
	}
}

func (b *Block) SetAlive(alive bool) {
	b.alive = alive
}

func (b *Block) Reset(probability float64) {
	b.alive = probability >= rand.Float64()
}

func (b Block) Print() {
	if b.alive {
		fmt.Print(b.aliveStr)
	} else {
		fmt.Print(b.deadStr)
	}
}

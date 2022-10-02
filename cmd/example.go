package main

import (
	"fmt"
	"github.com/a-skua/lifegame-go"
	"time"
)

const (
	L = lifegame.Live
	D = lifegame.Die
)

func show(table [][]lifegame.State) {
	for _, row := range table {
		for _, state := range row {
			if state.IsLive() {
				fmt.Print("■ ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func main() {
	lifegame := lifegame.New(lifegame.NewCell(6, 6, []lifegame.State{
		D, D, D, D, D, D,
		D, L, L, D, D, D,
		D, L, L, D, D, D,
		D, D, D, L, L, D,
		D, D, D, L, L, D,
		D, D, D, D, D, D,
	}))

	for {
		show(lifegame.Table())
		fmt.Println()
		lifegame.Next()
		time.Sleep(500 * time.Millisecond)
	}
}

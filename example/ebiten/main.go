package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/a-skua/lifegame-go"
	"github.com/hajimehoshi/ebiten/v2"
)

func writeTable(pixel []byte, table [][]lifegame.State) {
	pixel = pixel[:0]
	for _, row := range table {
		for _, state := range row {
			if state.IsLive() {
				pixel = append(pixel, 0xff, 0xff, 0xff, 0xff)
			} else {
				pixel = append(pixel, 0x00, 0x00, 0x00, 0x00)
			}
		}
	}
}

type Game struct {
	state  *lifegame.Lifegame
	width  int
	height int
	pixel  []byte
	ticker *time.Ticker
}

func NewGame() *Game {
	const width, height = 128, 96

	states := make([]lifegame.State, 0, width*height)
	for i := 0; i < width*height; i++ {
		var s lifegame.State
		if n := rand.Intn(2); n > 0 {
			s = lifegame.Live
		}

		states = append(states, s)
	}

	pixel := make([]byte, width*height*4)
	state := lifegame.New(lifegame.NewCell(width, height, states))
	writeTable(pixel, state.Table())

	return &Game{
		state:  state,
		width:  width,
		height: height,
		pixel:  pixel,
		ticker: time.NewTicker(100 * time.Millisecond),
	}
}

func (g *Game) Update() error {
	select {
	case <-g.ticker.C:
	default:
		return nil
	}
	defer g.state.Next()

	writeTable(g.pixel, g.state.Table())
	return nil
}

func (g *Game) Layout(w, h int) (int, int) {
	return g.width, g.height
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.WritePixels(g.pixel)
}

func main() {
	g := NewGame()

	ebiten.SetWindowSize(320, 240)
	ebiten.SetWindowTitle("Life Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

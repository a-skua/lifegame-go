package main

import (
	"log"
	"time"

	"github.com/a-skua/lifegame-go"
	"github.com/hajimehoshi/ebiten/v2"
)

func writeTable(pixel []byte, table [][]lifegame.State) {
	pixel = pixel[:0]
	for _, row := range table {
		for _, state := range row {
			if state.IsLive() {
				pixel = append(pixel, 0xff, 0xff, 0xff, 0x00)
			} else {
				pixel = append(pixel, 0x00, 0x00, 0x00, 0x00)
			}
		}
	}
}

const (
	L = lifegame.Live
	D = lifegame.Die
)

type Game struct {
	state  *lifegame.Lifegame
	width  int
	height int
	pixel  []byte
	ticker *time.Ticker
}

func NewGame() *Game {
	const width, height = 6, 6

	pixel := make([]byte, width*height*4)
	state := lifegame.New(lifegame.NewCell(width, height, []lifegame.State{
		D, D, D, D, D, D,
		D, L, L, D, D, D,
		D, L, L, D, D, D,
		D, D, D, L, L, D,
		D, D, D, L, L, D,
		D, D, D, D, D, D,
	}))
	writeTable(pixel, state.Table())

	return &Game{
		state:  state,
		width:  width,
		height: height,
		pixel:  pixel,
		ticker: time.NewTicker(500 * time.Millisecond),
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

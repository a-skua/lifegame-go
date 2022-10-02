package lifegame

type X int

func (x X) Int() int {
	return int(x)
}

type Y int

func (y Y) Int() int {
	return int(y)
}

type State bool

const (
	Live State = true
	Die  State = false
)

func (s State) IsLive() bool {
	return s == Live
}

func (s State) isDie() bool {
	return s == Die
}

type Cell struct {
	x      X
	y      Y
	states []State
}

func NewCell(x X, y Y, states []State) *Cell {
	return &Cell{x, y, states}
}

func (c *Cell) inrangeX(x X) bool {
	return x >= 0 && x < c.x
}

func (c *Cell) outrangeX(x X) bool {
	return !c.inrangeX(x)
}

func (c *Cell) inrangeY(y Y) bool {
	return y >= 0 && y < c.y
}

func (c *Cell) outrangeY(y Y) bool {
	return !c.inrangeY(y)
}

// Y\X| 0 | 1 | 2
//
//	0 | L   L   D
//	1 | L   D   D
//	2 | D   D   D
func (c *Cell) current(x X, y Y) State {
	if c.outrangeX(x) || c.outrangeY(y) {
		return Die
	}

	return c.states[c.x.Int()*y.Int()+x.Int()]
}

// Y\X| 0 | 1 | 2
//
//	0 | L   L   D
//	1 | L   D   D
//	2 | D   D   D
func (c *Cell) future(x X, y Y) State {
	aroundStates := c.aroundStates(x, y)

	countLive := 0
	for _, state := range aroundStates {
		if state.IsLive() {
			countLive += 1
		}
	}

	switch {
	case c.current(x, y).isDie() && countLive == 3,
		c.current(x, y).IsLive() && (countLive == 2 || countLive == 3):
		return Live
	default:
		return Die
	}
}

func (c *Cell) aroundStates(x X, y Y) []State {
	return []State{
		c.current(x-1, y-1),
		c.current(x, y-1),
		c.current(x+1, y-1),
		c.current(x+1, y),
		c.current(x+1, y+1),
		c.current(x, y+1),
		c.current(x-1, y+1),
		c.current(x-1, y),
	}
}

type Lifegame struct {
	cell *Cell
}

func New(cell *Cell) *Lifegame {
	return &Lifegame{cell}
}

func (lg *Lifegame) Next() {
	length := lg.cell.x.Int() * lg.cell.y.Int()

	futureStates := make([]State, 0, length)
	for y := Y(0); y < lg.cell.y; y++ {
		for x := X(0); x < lg.cell.x; x++ {
			futureStates = append(futureStates, lg.cell.future(x, y))
		}
	}

	lg.cell = NewCell(lg.cell.x, lg.cell.y, futureStates)
}

// return current states
func (lg *Lifegame) Table() [][]State {
	var (
		start = 0
		end   = lg.cell.x.Int()
	)

	table := make([][]State, 0, lg.cell.y.Int())
	for i := 0; i < lg.cell.y.Int(); i++ {
		table = append(table, lg.cell.states[start:end])
		start += lg.cell.x.Int()
		end += lg.cell.x.Int()
	}

	return table
}

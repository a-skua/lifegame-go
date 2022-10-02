package lifegame

import (
	"reflect"
	"testing"
)

func TestCell_current(t *testing.T) {
	type args struct {
		x X
		y Y
	}

	type test struct {
		name string
		cell *Cell
		args
		want State
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.current(tt.x, tt.y)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "get first",
			cell: NewCell(3, 3, []State{
				Live, Die, Die,
				Die, Die, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 0,
				y: 0,
			},
			want: Live,
		},
		{
			name: "get last",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Die, Die,
				Die, Die, Live,
			}),
			args: args{
				x: 2,
				y: 2,
			},
			want: Live,
		},
		{
			name: "get middle",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: Live,
		},
		{
			name: "outrange",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 3,
				y: 0,
			},
			want: Die,
		},
		{
			name: "outrange",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: -1,
				y: 0,
			},
			want: Die,
		},
		{
			name: "outrange",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 0,
				y: 3,
			},
			want: Die,
		},
		{
			name: "outrange",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 0,
				y: -1,
			},
			want: Die,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCell_future(t *testing.T) {
	type args struct {
		x X
		y Y
	}

	type test struct {
		name string
		cell *Cell
		args
		want State
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.future(tt.x, tt.y)
			if tt.want != got {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "die to live",
			cell: NewCell(3, 3, []State{
				Live, Live, Die,
				Live, Die, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: Live,
		},
		{
			name: "live to live",
			cell: NewCell(3, 3, []State{
				Live, Live, Die,
				Live, Die, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 0,
				y: 1,
			},
			want: Live,
		},
		{
			name: "live to live",
			cell: NewCell(3, 3, []State{
				Live, Live, Die,
				Live, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 0,
			},
			want: Live,
		},
		{
			name: "live to die",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Live, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: Die,
		},
		{
			name: "live to die",
			cell: NewCell(3, 3, []State{
				Live, Live, Live,
				Live, Live, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: Die,
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestCell_aroundStates(t *testing.T) {
	type args struct {
		x X
		y Y
	}

	type test struct {
		name string
		cell *Cell
		args
		want []State
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cell.aroundStates(tt.x, tt.y)
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Live, Die, Die,
				Die, Die, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Live, Die, Die, Die, Die, Die, Die, Die},
		},
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Die, Live, Die,
				Die, Die, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Die, Live, Die, Die, Die, Die, Die, Die},
		},
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Die, Die, Live,
				Die, Die, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Die, Die, Live, Die, Die, Die, Die, Die},
		},
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Die, Live,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Die, Die, Die, Live, Die, Die, Die, Die},
		},
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Die, Die,
				Die, Die, Live,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Die, Die, Die, Die, Live, Die, Die, Die},
		},
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Die, Die,
				Die, Live, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Die, Die, Die, Die, Die, Live, Die, Die},
		},
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Die, Die, Die,
				Live, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Die, Die, Die, Die, Die, Die, Live, Die},
		},
		{
			name: "true",
			cell: NewCell(3, 3, []State{
				Die, Die, Die,
				Live, Die, Die,
				Die, Die, Die,
			}),
			args: args{
				x: 1,
				y: 1,
			},
			want: []State{Die, Die, Die, Die, Die, Die, Die, Live},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestLifegame_Next(t *testing.T) {
	type test struct {
		name     string
		lifegame *Lifegame
		want     *Lifegame
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			tt.lifegame.Next()

			if !reflect.DeepEqual(tt.want, tt.lifegame) {
				t.Fatalf("want=%v, got=%v.", tt.want, tt.lifegame)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			lifegame: New(NewCell(3, 3, []State{
				Die, Die, Die,
				Live, Live, Live,
				Die, Die, Die,
			})),
			want: New(NewCell(3, 3, []State{
				Die, Live, Die,
				Die, Live, Die,
				Die, Live, Die,
			})),
		},
		{
			name: "true",
			lifegame: New(NewCell(3, 3, []State{
				Die, Live, Die,
				Die, Live, Die,
				Die, Live, Die,
			})),
			want: New(NewCell(3, 3, []State{
				Die, Die, Die,
				Live, Live, Live,
				Die, Die, Die,
			})),
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

func TestLifegame_Table(t *testing.T) {
	type test struct {
		name     string
		lifegame *Lifegame
		want     [][]State
	}

	do := func(tt *test) {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.lifegame.Table()
			if !reflect.DeepEqual(tt.want, got) {
				t.Fatalf("want=%v, got=%v.", tt.want, got)
			}
		})
	}

	tests := []*test{
		{
			name: "true",
			lifegame: New(NewCell(3, 3, []State{
				Die, Live, Die,
				Die, Live, Die,
				Die, Live, Die,
			})),
			want: [][]State{
				{Die, Live, Die},
				{Die, Live, Die},
				{Die, Live, Die},
			},
		},
		{
			name: "true",
			lifegame: New(NewCell(3, 3, []State{
				Die, Die, Die,
				Live, Live, Live,
				Die, Die, Die,
			})),
			want: [][]State{
				{Die, Die, Die},
				{Live, Live, Live},
				{Die, Die, Die},
			},
		},
	}

	for _, tt := range tests {
		do(tt)
	}
}

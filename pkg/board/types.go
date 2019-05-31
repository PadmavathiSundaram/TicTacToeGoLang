package board

import (
	"bytes"
	"fmt"
	"sort"
)

const boardStyle = "\n───┼───┼───"
const strikeThrough = "\u0336"

type board struct {
	Cells         map[int]string
	enabledCells  []int
	UserInterface Userinterface
}

func NewBoard() *board {
	indexedCells := make(map[int]string, 9)

	for x := 0; x < 9; x++ {
		indexedCells[x] = ""
	}

	b := board{Cells: indexedCells, UserInterface: NewConsoleInput()}
	return &b
}

func (b *board) DisplayWinner(winningCells []int) {
	for _, index := range winningCells {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, strikeThrough)
		fmt.Fprintf(&buf, b.Cells[index])
		fmt.Fprintf(&buf, strikeThrough)
		b.Cells[index] = buf.String()
	}
	b.Display()
}

func (b *board) Display() {

	// To store the keys in slice in sorted order
	var keys []int
	for k := range b.Cells {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var buf bytes.Buffer
	// To perform the opertion you want
	for _, k := range keys {
		i := b.Cells[k]
		if i == "" {
			fmt.Fprintf(&buf, "(%d)|", k)
		} else {
			fmt.Fprintf(&buf, " %s |", i)
		}
		if k == 2 || k == 5 {
			buf.Truncate(buf.Len() - 1)
			fmt.Fprintln(&buf, boardStyle)
		}
	}
	buf.Truncate(buf.Len() - 1)

	fmt.Println(buf.String())

}

func (b *board) ValidateIndex(x int) bool {
	if x >= 9 {
		fmt.Printf("Invalid index %d \n", x)
		return false
	}
	cell := b.Cells[x]
	if cell != "" {
		fmt.Printf("Cell %d Already taken by %s\n", x, cell)
		return false
	}
	return true
}

func (b *board) Click(index int, player string) *board {
	b.Cells[index] = player
	b.enabledCells = append(b.enabledCells, index)
	return b
}

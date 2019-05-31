package board

import (
	"fmt"

	"github.com/tictactoe/pkg/screen"
)

type board struct {
	Cells         map[int]string
	EnabledCells  []int
	UserInterface screen.Userinterface
}

func NewBoard() *board {
	indexedCells := make(map[int]string, 9)

	for x := 0; x < 9; x++ {
		indexedCells[x] = ""
	}

	b := board{Cells: indexedCells, UserInterface: screen.NewConsoleInput(indexedCells)}
	return &b
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
	b.EnabledCells = append(b.EnabledCells, index)
	return b
}

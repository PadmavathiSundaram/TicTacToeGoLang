package board

import (
	"fmt"
	"sort"
)

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

func (b *board) Display() {

	// To store the keys in slice in sorted order
	var keys []int
	for k := range b.Cells {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		i := b.Cells[k]
		if i == "" {
			fmt.Printf("(%d)", k)
		} else {
			fmt.Printf("%s", i)
		}
		fmt.Print("\t\t")
		if k == 2 || k == 5 || k == 8 {
			fmt.Println("\n\t")
		}
	}
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

package screen

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

const boardStyle = "\n───┼───┼───"
const strikeThrough = "\u0336"

type Userinterface interface {
	FetchInput() (int, error)
	DisplayWinner(winningCells []int)
	Display()
}

type consoleInput struct {
	reader *bufio.Reader
	cells  map[int]string
}

func NewConsoleInput(cells map[int]string) Userinterface {
	reader := bufio.NewReader(os.Stdin)
	return &consoleInput{reader, cells}
}

func (c *consoleInput) FetchInput() (int, error) {

	fmt.Println("Select the index you want to occupy: ")
	index, err := c.reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	var x int
	if _, err := fmt.Sscan(index, &x); err != nil {
		return 0, err
	}

	return x, nil
}

func (c *consoleInput) DisplayWinner(winningCells []int) {
	for _, index := range winningCells {
		var buf bytes.Buffer
		fmt.Fprintf(&buf, strikeThrough)
		fmt.Fprintf(&buf, c.cells[index])
		fmt.Fprintf(&buf, strikeThrough)
		c.cells[index] = buf.String()
	}
	c.Display()
}

func (c *consoleInput) Display() {

	// To store the keys in slice in sorted order
	var keys []int
	for k := range c.cells {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var buf bytes.Buffer
	// To perform the opertion you want
	for _, k := range keys {
		i := c.cells[k]
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

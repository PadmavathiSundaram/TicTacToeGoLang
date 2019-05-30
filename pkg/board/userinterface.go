package board

import (
	"bufio"
	"fmt"
	"os"
)

type Userinterface interface {
	fetchInput() (int, error)
}

type consoleInput struct {
	reader *bufio.Reader
}

func NewConsoleInput() Userinterface {
	reader := bufio.NewReader(os.Stdin)
	return &consoleInput{reader}
}

func (c *consoleInput) fetchInput() (int, error) {

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

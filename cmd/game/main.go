package main

import (
	br "github.com/tictactoe/pkg/board"
)

func main() {
	b := br.NewBoard()
	br.Play(b)
}

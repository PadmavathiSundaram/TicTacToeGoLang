package board

import (
	"fmt"
)

const winningText = "Player %s is the winner.Game Over\n"
const drawText = "Its a Draw..Game Over"
const X = "X"
const Y = "Y"

var winningList = [][]int{{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6}}

func Play(b *board) {
	b.Display()
	player1 := true

	for !isGameOver(b) {
		nextTurn(b, player1)
		player1 = !player1
	}
}

func isGameOver(b *board) bool {

	if len(b.enabledCells) >= 5 {
		if winner := WrapperToDisplayExecutionTime(b, "isThereAWinner", isThereAWinner); winner {
			return true
		}

		if len(b.enabledCells) >= 9 {
			fmt.Println(drawText)
			return true
		}
	}
	return false
}

func isThereAWinner(b *board) bool {
	winningCells := make([]int, 3)
	for k := range winningList {
		matchYCount := 0
		matchXCount := 0

		winningCells = winningCells[0:0]
		for row := range winningList[k] {
			cellData := b.Cells[winningList[k][row]]
			if Contains(b.enabledCells, winningList[k][row]) {
				if cellData == X {
					matchXCount++
				} else {
					matchYCount++
				}
				winningCells = append(winningCells, winningList[k][row])
			}

			if matchXCount == 3 {
				fmt.Printf(winningText, X)
				b.DisplayWinner(winningCells)
				return true
			}

			if matchYCount == 3 {
				fmt.Printf(winningText, Y)
				b.DisplayWinner(winningCells)
				return true
			}
		}
	}

	return false
}

func nextTurn(b *board, flag bool) *board {

	var player = X
	if !flag {
		player = Y
	}

	fmt.Printf("Player %s's turn", player)

	index, err := b.UserInterface.fetchInput()
	if err != nil {
		fmt.Println("Invalid Input.Please enter a valid number between 0-9")
		return nextTurn(b, flag)
	}

	if !b.ValidateIndex(index) {
		return nextTurn(b, flag)
	}

	b = b.Click(index, player)
	b.Display()
	return b
}

package board

import (
	"fmt"
	"testing"

	"github.com/tictactoe/pkg/util"
	"gotest.tools/assert"
)

func TestIsGameOver(t *testing.T) {
	b := NewBoard()
	b.Cells[2] = X
	b.Cells[4] = X
	b.Cells[3] = X
	b.Cells[8] = X
	b.EnabledCells = append(b.EnabledCells, 0, 1, 2, 3, 4, 5, 6, 7, 8)

	winningYBoard := NewBoard()
	winningYBoard.EnabledCells = append(winningYBoard.EnabledCells, 0, 1, 2, 3, 4, 5, 6, 7, 8)

	winningXBoard := NewBoard()
	winningXBoard.Cells[0] = X
	winningXBoard.Cells[1] = X
	winningXBoard.Cells[2] = X
	winningXBoard.EnabledCells = append(winningXBoard.EnabledCells, 0, 1, 2, 3, 4, 5, 6, 7, 8)

	tests := []struct {
		name           string
		board          *board
		expectedResult bool
	}{
		{"New Game", new(board), false},
		{"winner Y found", winningYBoard, true},
		{"winner X found", winningXBoard, true},
		{"Draw case", b, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isGameOver(test.board)
			assert.Equal(t, result, test.expectedResult, "isGameOver returned unexpected result")

		})
	}
}

type mockUI struct {
	index int
}

func (m *mockUI) Display() {
}

func (m *mockUI) DisplayWinner(winningCells []int) {
}

func (m *mockUI) FetchInput() (int, error) {
	if m.index == 5 {
		m.index++
		return 0, fmt.Errorf("Bad data")
	}
	if m.index == 2 {
		m.index++
		return 1, nil
	}

	if m.index == 3 {
		m.index++
		return 9, nil
	}
	m.index++

	return m.index, nil
}

func TestNextTurn(t *testing.T) {

	b := NewBoard()
	b.UserInterface = &mockUI{index: 0}

	invalidBoard := NewBoard()
	invalidBoard.UserInterface = &mockUI{index: 5}

	tests := []struct {
		name           string
		board          *board
		isPlayerX      bool
		expectedResult int
	}{
		{"Player X turn", b, true, 1},
		{"Player Y turn", b, false, 2},
		{"Player X turn - invalid index", b, true, 1},
		{"Invalid data turn", invalidBoard, false, 7},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := nextTurn(test.board, test.isPlayerX)
			assert.Assert(t, util.Contains(result.EnabledCells, test.expectedResult), "isGameOver returned unexpected result")
		})
	}
}

func TestPlay(t *testing.T) {
	b := NewBoard()
	b.EnabledCells = append(b.EnabledCells, 0,
		2, 3, 4)
	b.UserInterface = &mockUI{index: 0}

	Play(b)
	assert.Assert(t, util.Contains(b.EnabledCells, 1), "isGameOver returned unexpected result")
}

func TestWrapperToDisplayExecutionTime(t *testing.T) {

	mockProcess := func(b *board) bool {
		return true
	}
	winner := WrapperToDisplayExecutionTime(nil, "mockProcess", mockProcess)
	assert.Equal(t, winner, true)

}

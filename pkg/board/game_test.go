package board

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
)

func TestIsGameOver(t *testing.T) {
	b := NewBoard()
	b.Cells[2] = X
	b.Cells[4] = X
	b.Cells[3] = X
	b.Cells[8] = X
	b.enabledCells = append(b.enabledCells, 0, 1, 2, 3, 4, 5, 6, 7, 8)

	winningYBoard := new(board)
	winningYBoard.enabledCells = append(winningYBoard.enabledCells, 0, 1, 2, 3, 4, 5, 6, 7, 8)

	winningXBoard := NewBoard()
	winningXBoard.Cells[0] = X
	winningXBoard.Cells[1] = X
	winningXBoard.Cells[2] = X
	winningXBoard.enabledCells = append(winningXBoard.enabledCells, 0, 1, 2, 3, 4, 5, 6, 7, 8)

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

func (m *mockUI) fetchInput() (int, error) {
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
			assert.Assert(t, Contains(result.enabledCells, test.expectedResult), "isGameOver returned unexpected result")
		})
	}
}

func TestPlay(t *testing.T) {
	b := NewBoard()
	b.enabledCells = append(b.enabledCells, 0, 2, 3, 4)
	b.UserInterface = &mockUI{index: 0}

	Play(b)
	assert.Assert(t, Contains(b.enabledCells, 1), "isGameOver returned unexpected result")

}

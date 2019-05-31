package screen

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestFetchInput(t *testing.T) {
	indexedCells := make(map[int]string, 9)

	eofUI := NewConsoleInput(indexedCells)

	s := strings.NewReader("2\n")
	reader := bufio.NewReader(s)
	ui := &consoleInput{reader, indexedCells}

	invalidInput := strings.NewReader("abcd\n")
	invalidReader := bufio.NewReader(invalidInput)
	invalidUI := &consoleInput{invalidReader, indexedCells}

	tests := []struct {
		name          string
		ui            Userinterface
		expectedIndex int
		err           string
	}{
		{"EOF scenario", eofUI, 0, "EOF"},
		{"Valid index", ui, 2, ""},
		{"Valid index", invalidUI, 0, "expected integer"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			index, err := test.ui.FetchInput()
			assert.Equal(t, index, test.expectedIndex, "Mismatched scanned value")
			if err != nil {
				assert.Equal(t, err.Error(), test.err, "Missing expected error")
			}
		})
	}
}

func TestDisplayWinner(t *testing.T) {
	indexedCells := make(map[int]string, 9)
	indexedCells[3] = ""
	s := strings.NewReader("3\n")
	reader := bufio.NewReader(s)
	ui := &consoleInput{reader, indexedCells}

	var buf bytes.Buffer
	fmt.Fprintf(&buf, strikeThrough)
	fmt.Fprintf(&buf, "")
	fmt.Fprintf(&buf, strikeThrough)

	ui.DisplayWinner([]int{0, 1, 2})

	assert.Equal(t, ui.cells[0], buf.String())

}

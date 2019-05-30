package board

import (
	"bufio"
	"gotest.tools/assert"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	eofUI := NewConsoleInput()

	s := strings.NewReader("2\n")
	reader := bufio.NewReader(s)
	ui := &consoleInput{reader}

	invalidInput := strings.NewReader("abcd\n")
	invalidReader := bufio.NewReader(invalidInput)
	invalidUI := &consoleInput{invalidReader}

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
			index, err := test.ui.fetchInput()
			assert.Equal(t, index, test.expectedIndex, "Mismatched scanned value")
			if err != nil {
				assert.Equal(t, err.Error(), test.err, "Missing expected error")
			}
		})
	}
}

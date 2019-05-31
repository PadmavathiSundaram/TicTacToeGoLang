package util

import (
	"fmt"

	"gotest.tools/assert"
	"testing"
)

func TestConatins(t *testing.T) {
	testScenarios := []struct {
		Description    string
		DataList       []int
		SearchString   int
		ExpectedResult bool
	}{
		{"Match Found", []int{0, 1, 2}, 2, true},
		{"Match not found", []int{3, 4, 5}, 2, false},
		{"Invalid data", []int{}, 0, false},
	}
	for _, td := range testScenarios {
		t.Run(fmt.Sprintf("%s - %v contains %v : %v",
			td.Description, td.DataList, td.SearchString, td.ExpectedResult), func(t *testing.T) {
			result := Contains(td.DataList, td.SearchString)
			assert.Equal(t, result, td.ExpectedResult)
		})
	}
}

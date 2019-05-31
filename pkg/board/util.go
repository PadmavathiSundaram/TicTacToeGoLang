package board

import (
	"fmt"
	"time"
)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func WrapperToDisplayExecutionTime(b *board, processName string, process func(b *board) bool) bool {
	startTime := time.Now()
	winner := process(b)
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Printf("Total time taken by %s loop is %v nano seconds \n", processName, diff.Nanoseconds())
	return winner
}

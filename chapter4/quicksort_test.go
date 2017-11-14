package chapter4

import (
	"fmt"
	"testing"
)

func Test_AdvanceQuickSort(t *testing.T) {
	sortArray := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	quicksortAdvance(sortArray, 0, len(sortArray)-1)

	for _, value := range sortArray {
		fmt.Printf("%d\t", value)
	}
}

func Test_QuickSort(t *testing.T) {
	sortArray := []int{41, 24, 76, 11, 45, 64, 21, 69, 19, 36}
	quicksort(sortArray, 0, len(sortArray)-1)

	for _, value := range sortArray {
		fmt.Printf("%d\t", value)
	}
}

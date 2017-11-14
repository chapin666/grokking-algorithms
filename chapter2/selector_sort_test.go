package chapter2

import (
	"fmt"
	"testing"
)

func Test_SelectorSort(t *testing.T) {
	datas := []int{4, 7, 2, 67, 45, 34, 2, 5, 6}
	datas = selectorSort(datas)
	fmt.Println(datas)
}

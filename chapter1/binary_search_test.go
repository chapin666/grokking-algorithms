package chapter1

import (
	"fmt"
	"testing"
)

func TestBinarySearchRecursion(t *testing.T) {
	key := 10
	datas := []int{1, 2, 3, 5, 8, 10, 24, 27, 40}
	index := binarySearchRecursion(datas, 0, len(datas), key)
	fmt.Println("TestBinarySearchRecursion")
	fmt.Println("the position is: ", index)
}

func TestBinarySearchOrder(t *testing.T) {
	key := 10
	datas := []int{1, 2, 3, 5, 8, 10, 24, 27, 40}
	index := binarySearchOrder(datas, 0, len(datas), key)
	fmt.Println("TestBinarySearchOrder")
	fmt.Println("the position is: ", index)
}

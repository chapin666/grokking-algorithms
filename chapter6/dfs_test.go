package chapter6

import (
	"fmt"
	"testing"
)

func Test_Search(t *testing.T) {
	s := search("you")
	fmt.Println(*s)
}

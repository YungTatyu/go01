package main

import (
	"fmt"
	"piscine"
)

func main() {
	s := []int{2, 1, 8, 7, 3, 5}
	piscine.SortIntegerTable(s)
	fmt.Println(s)
}

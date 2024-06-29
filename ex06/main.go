package main

import (
	"fmt"
	"piscine"
)

func main() {
	var (
		a int = 4
		b int = 2
	)
	piscine.Swap(&a, &b)
	fmt.Println(a, b)
}

package main

import (
	"fmt"
	"piscine"
)

func main() {
	var num int = 100
	var ptrNum *int = &num
	var ptrPtr **int = &ptrNum
	piscine.UltimatePointOne(&ptrPtr)
	fmt.Println(num)
}

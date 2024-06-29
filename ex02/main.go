package main

import (
	"fmt"
	"piscine"
)

func main() {
	var div int
	var mod int
	piscine.DivMod(0, 2, &div, &mod)
	fmt.Println("div=", div)
	fmt.Println("mod=", mod)
}

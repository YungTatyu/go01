package main

import (
	"fmt"
	"piscine"
)

func main() {
	var div int = 14
	var mod int = 2
	piscine.UltimateDivMod(&div, &mod)
	fmt.Println("div=", div)
	fmt.Println("mod=", mod)
}

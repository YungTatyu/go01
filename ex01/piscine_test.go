package main

import (
	"piscine"
	"testing"
)

func TestUltimatePointOne(t *testing.T) {
	var num int = 100
	var ptrNum *int = &num
	var ptrPtr **int = &ptrNum
	piscine.UltimatePointOne(&ptrPtr)
	if num != 1 {
		t.Errorf("expected 1, but got %d", num)
	}
}

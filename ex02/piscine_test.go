package main

import (
	"piscine"
	"testing"
)

func expect(actual int, expect int, t *testing.T) {
	if actual != expect {
		t.Errorf("expected %d, but got %d", expect, actual)
	}
}

func expectPanic(f func(), t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestNoRemain(t *testing.T) {
	var div int
	var mod int
	piscine.DivMod(10, 2, &div, &mod)
	expect(div, 5, t)
	expect(mod, 0, t)
}

func TestRemain(t *testing.T) {
	var div int
	var mod int
	piscine.DivMod(10, 3, &div, &mod)
	expect(div, 3, t)
	expect(mod, 1, t)
}

func TestDividedByZero(t *testing.T) {
	var div int
	var mod int
	expectPanic(func() { piscine.DivMod(10, 0, &div, &mod) }, t)
}

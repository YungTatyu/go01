package test

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
	var div int = 10
	var mod int = 2
	piscine.UltimateDivMod(&div, &mod)
	expect(div, 5, t)
	expect(mod, 0, t)
}

func TestRemain(t *testing.T) {
	var div int = 10
	var mod int = 3
	piscine.UltimateDivMod(&div, &mod)
	expect(div, 3, t)
	expect(mod, 1, t)
}

func TestNegative(t *testing.T) {
	var div int = -5
	var mod int = 2
	piscine.UltimateDivMod(&div, &mod)
	expect(div, -2, t)
	expect(mod, -1, t)
}

func TestDevideZero(t *testing.T) {
	var div int = 0
	var mod int = 2
	piscine.UltimateDivMod(&div, &mod)
	expect(div, 0, t)
	expect(mod, 0, t)
}

func TestDividedByZero(t *testing.T) {
	var div int = 2
	var mod int = 0
	expectPanic(func() { piscine.UltimateDivMod(&div, &mod) }, t)
}

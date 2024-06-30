package test

import (
	"piscine"
	"testing"
)

func expect(actual, expect int, t *testing.T) {
	if actual != expect {
		t.Errorf("expected%d, but got %d", expect, actual)
	}
}

func TestBasicAtoi(t *testing.T) {
	expect(piscine.BasicAtoi(""), 0, t)
	expect(piscine.BasicAtoi("test"), 0, t)
	expect(piscine.BasicAtoi("42"), 42, t)
	expect(piscine.BasicAtoi("000000000000000000042"), 42, t)
	expect(piscine.BasicAtoi("0000000000000000000"), 0, t)
	expect(piscine.BasicAtoi("012 345"), 0, t)
	expect(piscine.BasicAtoi("9223372036854775807"), 9223372036854775807, t)
	expect(piscine.BasicAtoi("9223372036854775808"), -9223372036854775808, t)
}

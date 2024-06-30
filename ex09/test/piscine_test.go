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

func TestBasicAtoi2(t *testing.T) {
	expect(piscine.BasicAtoi2(""), 0, t)
	expect(piscine.BasicAtoi2("test"), 0, t)
	expect(piscine.BasicAtoi2("42"), 42, t)
	expect(piscine.BasicAtoi2("000000000000000000042"), 42, t)
	expect(piscine.BasicAtoi2("0000000000000000000"), 0, t)
	expect(piscine.BasicAtoi2("012 345"), 0, t)
	expect(piscine.BasicAtoi2("9223372036854775807"), 9223372036854775807, t)
	expect(piscine.BasicAtoi2("9223372036854775808"), -9223372036854775808, t)
}

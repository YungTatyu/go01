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

func TestAtoi(t *testing.T) {
	expect(piscine.Atoi(""), 0, t)
	expect(piscine.Atoi("test"), 0, t)
	expect(piscine.Atoi("Hello World!"), 0, t)
	expect(piscine.Atoi("42"), 42, t)
	expect(piscine.Atoi("000000000000000000042"), 42, t)
	expect(piscine.Atoi("0000000000000000000"), 0, t)
	expect(piscine.Atoi("012 345"), 0, t)
	expect(piscine.Atoi("9223372036854775807"), 9223372036854775807, t)
	expect(piscine.Atoi("9223372036854775808"), -9223372036854775808, t)
	expect(piscine.Atoi("-9223372036854775808"), -9223372036854775808, t)
	expect(piscine.Atoi("-9223372036854775809"), 9223372036854775807, t)
	expect(piscine.Atoi("+42"), 42, t)
	expect(piscine.Atoi("-42"), -42, t)
	expect(piscine.Atoi("+0"), 0, t)
	expect(piscine.Atoi("-0"), 0, t)
	expect(piscine.Atoi("++100"), 0, t)
	expect(piscine.Atoi("--100"), 0, t)
}

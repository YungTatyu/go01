package test

import (
	"piscine"
	"testing"
)

func expect(actual, expect string, t *testing.T) {
	if actual != expect {
		t.Errorf("expected%s, but got %s", expect, actual)
	}
}

func TestStrRev(t *testing.T) {
	expect(piscine.StrRev("Hello World!"), "!dlroW olleH", t)
	expect(piscine.StrRev("i"), "i", t)
	expect(piscine.StrRev(""), "", t)
	expect(piscine.StrRev("  "), "  ", t)
}

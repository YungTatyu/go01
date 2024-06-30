package test

import (
	"piscine"
	"reflect"
	"testing"
)

func expect(nums, expect []int, t *testing.T) {
	piscine.SortIntegerTable(nums)
	if !reflect.DeepEqual(nums, expect) {
		t.Errorf("expected%d, but got %d", expect, nums)
	}
}

func TestSortIntegerTable(t *testing.T) {
	expect([]int{5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5}, t)
	expect([]int{0, 1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4, 5}, t)
	expect([]int{0, 9, 3, 100, 1, 0}, []int{0, 0, 1, 3, 9, 100}, t)
}

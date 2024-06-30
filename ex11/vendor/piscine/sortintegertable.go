package piscine

func quicksort(nums []int, leftmostIndex, righmostIndex int) {
	if leftmostIndex < righmostIndex {
		pivotIndex := partition(nums, leftmostIndex, righmostIndex)
		// partitionで分けられた左右のpartをそれぞれsortする
		quicksort(nums, leftmostIndex, pivotIndex-1)
		quicksort(nums, pivotIndex+1, righmostIndex)
	}
}

func partition(nums []int, leftmostIndex, rightmostIndex int) int {
	pivot := nums[rightmostIndex] // ピボットを右端とする
	// pivotより値が大きい要素のindex
	firstGreterIndex := leftmostIndex - 1

	for i := leftmostIndex; i < rightmostIndex; i++ {
		if nums[i] <= pivot {
			firstGreterIndex++
			nums[firstGreterIndex], nums[i] = nums[i], nums[firstGreterIndex]
		}
	}
	firstGreterIndex++ // pivotと一番左にあるpivotより大きい値をswap
	nums[firstGreterIndex], nums[rightmostIndex] = nums[rightmostIndex], nums[firstGreterIndex]
	return firstGreterIndex
}

func SortIntegerTable(table []int) {
	quicksort(table, 0, len(table)-1)
}

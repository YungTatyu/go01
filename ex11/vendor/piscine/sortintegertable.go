package piscine

func quicksort(nums []int) {
	quicksortRecursive(nums, 0, len(nums)-1)
}

func quicksortRecursive(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition(nums, low, high)
		quicksortRecursive(nums, low, pivotIndex-1)
		quicksortRecursive(nums, pivotIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high] // ピボットを右端とする
	i := low - 1

	for j := low; j < high; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[high] = nums[high], nums[i+1]
	return i + 1
}

func SortIntegerTable(table []int) {
	quicksort(table)
}

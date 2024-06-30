package piscine

func quicksort(nums []int, low, high int) {
	if low < high {
		pivotIndex := partition(nums, low, high)
		quicksort(nums, low, pivotIndex-1)
		quicksort(nums, pivotIndex+1, high)
	}
}

func partition(nums []int, low, high int) int {
	pivot := nums[high] // ピボットを右端とする
	cur_low := low - 1

	for i := low; i < high; i++ {
		if nums[i] <= pivot {
			cur_low++
			nums[cur_low], nums[i] = nums[i], nums[cur_low]
		}
	}

	nums[cur_low+1], nums[high] = nums[high], nums[cur_low+1]
	return cur_low + 1
}

func SortIntegerTable(table []int) {
	quicksort(table, 0, len(table)-1)
}

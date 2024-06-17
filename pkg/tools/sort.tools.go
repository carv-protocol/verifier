package tools

func BinarySearch[T any](nums []T, target T, compare func(T, T) bool) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if compare(nums[mid], target) {
			return mid
		} else if compare(target, nums[mid]) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

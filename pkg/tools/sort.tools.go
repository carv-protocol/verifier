package tools

type LessFunc[T any] func(a, b T) bool

func BinarySearch[T any](slice []T, val T, less LessFunc[T]) int {
	low := 0
	high := len(slice) - 1
	for low <= high {
		mid := (low + high) / 2
		if less(slice[mid], val) {
			low = mid + 1
		} else if less(val, slice[mid]) {
			high = mid - 1
		} else {
			return mid // 找到目标值
		}
	}
	return -1 // 未找到目标值
}

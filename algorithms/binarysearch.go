package algorithms

/*
	Requires a sorted array as input a
*/

func BinarySearch[T any](key T, a []T, cmp Comparator[T]) T {

	low := 0
	high := len(a) - 1

	for low <= high {
		mid := (low + high) / 2
		comparison := cmp(a[mid], key)
		if comparison == 0 {
			return a[mid]
		}
		if comparison < 0 {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	var zero T
	return zero
}

package algorithms

type Comparator[T any] func(a, b T) int

func Less[T any](a T, w T, cmp Comparator[T]) bool {
	return cmp(a, w) < 0
}

func Swap[T any](a []T, i, j int) {
	temp := a[i]
	a[i] = a[j]
	a[j] = temp
}

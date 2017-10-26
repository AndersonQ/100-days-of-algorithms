package sort

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, lo, hi int) {
	if lo < hi {
		p := partition(arr, lo, hi)
		quickSort(arr, lo, p)
		quickSort(arr, p+1, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo - 1
	for j := lo; j < hi-1; j++ {
		if arr[j] < pivot {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	if arr[hi] < arr[i+1] {
		arr[i+1], arr[hi] = arr[hi], arr[i+1]
	}
	return i + 1
}

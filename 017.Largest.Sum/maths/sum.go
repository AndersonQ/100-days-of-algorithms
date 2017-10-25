package maths

func LargestSum(arr []int) int {
	sum := 0
	max := arr[0]
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			sum += arr[j]
			if sum > max {
				max = sum
			}
		}
		sum = 0
	}

	return max
}

package homework

func reverse(input []int64) (result []int64) {
	n := len(input)
	result = make([]int64, n)
	for i := 0; i < n; i++ {
		result[n-1-i] = input[i]
	}
	return result
}

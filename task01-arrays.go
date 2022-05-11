package homework

const N = 15

func average(input [N]float32) (result float32) {
	for _, v := range input {
		result += v
	}
	return result / N
}

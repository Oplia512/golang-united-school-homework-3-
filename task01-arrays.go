package homework

func average(input [15]float32) (result float32) {

	for _, number := range input {
		result = result + number
	}

	result = result / (float32(len(input)))
	return result
}

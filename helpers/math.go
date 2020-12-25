package helpers

func Multiply(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := 1
	for _, n := range numbers {
		result *= n
	}
	return result
}

func Sum(numbers []int) int {
	result := 0
	for _, x := range numbers {
		result += x
	}
	return result
}

func FindMinAndMax(numbers []int) (int, int) {
	min, max := numbers[0], numbers[0]
	for _, x := range numbers {
		if x < min {
			min = x
		}
		if x > max {
			max = x
		}
	}
	return min, max
}

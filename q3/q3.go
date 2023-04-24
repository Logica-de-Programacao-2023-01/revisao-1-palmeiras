package q3

import (
	"errors"
	"math"
)

func FindMinMaxAverage(numbers []int) (int, int, float64, error) {
	if len(numbers) == 0 {
		return 0, 0, 0, errors.New("lista vazia")
	}

	if len(numbers) == 1 {
		return numbers[0], numbers[0], float64(numbers[0]), nil
	}

	var max = math.MinInt64
	var min = math.MaxInt64
	var sum = 0

	for _, number := range numbers {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
		sum += number
	}

	// Para calcular a média, desconsideramos o maior e o menor valor
	average := float64(sum-max-min) / float64(len(numbers)-2)

	return max, min, average, nil
}

package quiz

import "errors"

func Sum(scores []int) int {
	sumScore := 0

	for _, score := range scores {
		sumScore += score
	}
	return sumScore
}

func Calculate(a, b int, mathOperation string) (result float64, err error) {
	switch mathOperation {
	case "+":
		result = float64(a) + float64(b)

	case "-":
		result = float64(a) - float64(b)

	case "*":
		result = float64(a) * float64(b)

	case "/":
		result = float64(a) / float64(b)

	default:
		result = 0
		err = errors.New("unknown Operation")
	}

	return
}

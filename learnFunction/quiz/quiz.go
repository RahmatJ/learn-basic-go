package quiz

func sum(scores []int) int {
	sumScore := 0

	for _, score := range scores {
		sumScore += score
	}
	return sumScore
}

//TODO: make calculate function

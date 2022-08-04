package main

import (
	"fmt"
)

func main() {
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}
	//	1.
	//	hitung rata rata
	mean := calculateMean(scores)
	fmt.Println("Rerata: ", mean)

	//	2.
	//	ambil nilai yang >=90 ke var baru
	var goodScores []int
	goodScores = filterNumbers(scores)
	fmt.Println("filtered: ", goodScores)
}

func calculateMean(numbers [8]int) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += float64(number)
	}
	var mean = sum / float64(len(numbers))
	return mean
}

func filterNumbers(numbers [8]int) []int {
	var filtered []int

	for _, number := range numbers {
		if number >= 90 {
			filtered = append(filtered, number)
		}
	}

	return filtered
}

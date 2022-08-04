package goKompleks

import "fmt"

func GoKompleksSliceOfMap() {
	students := []map[string]string{
		{"name": "Agung", "score": "A"},
		{"name": "Link", "score": "B"},
		{"name": "Mario", "score": "C"},
	}

	//print
	fmt.Println(students)

	// foreach
	for _, student := range students {
		fmt.Println("name: ", student["name"], ", score: ", student["score"])
	}
}

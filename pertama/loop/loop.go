package loop

import "fmt"

func LoopFor() {
	for i := 1; i <= 100; i++ {
		fmt.Println("Saya Sedang Belajar GO: ", i)
	}
}

func LoopWhile() {
	i := 1
	for i <= 100 {
		fmt.Println("While: Saya Sedang Belajar GO: ", i)
		i++
	}
}

func LoopString() {
	title := "Golang the best language"
	for index, letter := range title {
		fmt.Println("index: ", index, " letter: ", string(letter))
	}

	// _ digunakan untuk menyimpan variable yang tidak terpakai, tapi tidak bisa diakses juga
	for _, letter := range title {
		fmt.Println("withUnderscoreLetter: ", string(letter))
	}
}

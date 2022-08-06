package main

import (
	"fmt"
	"learnFunction/quiz"
)

func main() {
	printMyResult()
	printMyResultWithInput("saya")
	printMyResultWithInput("sedang belajar")
	printMyResultWithInput("golang")
	sentence := printMyResultWithInputAndOutput("belajar golang")
	fmt.Println(sentence)
	addResult := add(10, 10)
	fmt.Println("Add Result:", addResult)
	//	kalau ada fungsi dengan multiple return maka bisa dilakukan seperti berikut
	//	jika hanya ingin menggunakan salah satu, ganti dengan _ return yang tidak terpakai
	luas, keliling := calculate(10, 20)
	fmt.Println("Luas:", luas, ", Keliling:", keliling)

	luas, keliling = calculatePredefined(10, 20)
	fmt.Println("Predefined: Luas:", luas, ", Keliling:", keliling)

	//	TODO: run Sum and Calculate
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	sum := quiz.Sum(numbers)
	fmt.Println("sum:", sum)
	calculate, err := quiz.Calculate(10, 5, "+")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`calculate 10 + 5:`, calculate)
	calculate, err = quiz.Calculate(10, 5, "-")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`calculate 10 - 5:`, calculate)
	calculate, err = quiz.Calculate(10, 5, "*")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`calculate 10 * 5:`, calculate)
	calculate, err = quiz.Calculate(10, 5, "/")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`calculate 10 / 5:`, calculate)

	calculate, err = quiz.Calculate(10, 5, "s")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(`calculate 10 - 5:`, calculate)
}

//simple function tanpa input
func printMyResult() {
	fmt.Println("Saya sedang belajar Go")
}

//simple function dengan input
func printMyResultWithInput(sentence string) {
	fmt.Println(sentence)
}

// simple function dengan input dan output
func printMyResultWithInputAndOutput(sentence string) string {
	newSentence := sentence + "."
	return newSentence
}

//kalau ada 2 input dengan tipe data sama bisa disederhanakan
// fungsi add berarti memiliki inputan number dan numberTwo dengan tipe data int
func add(number, numberTwo int) int {
	return number + numberTwo
}

//multiple return
//menghitung luas dan keliling persegi panjang
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

//predefined return
//variable return di definisikan di return type, sehingga waktu return tidak perlu
//menulis variable apa yang di return
func calculatePredefined(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}

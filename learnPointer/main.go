package main

import (
	"fmt"
	"learnPointer/pointerStruct"
	"learnPointer/quiz"
)

func main() {
	numberA := 5
	//tanda & akan menghasilkan alamat memory dari suatu variable
	//pada kasus ini numberB bakal menghasilkan alamat dari numberA
	//istilah referencing
	numberB := &numberA

	//fmt.Println(numberA)
	//fmt.Println(numberB)
	//	cara untuk mengakses isi alamat dari suatu variable bisa menggunakan tanda *
	// *numberB akan menghasilkan nilai variable 5
	// istilah dereference
	//fmt.Println(*numberB)

	//	karena numberB mengacu ke alamat memory yang sama dengan numberA
	//	maka kalau ada perubahan nilai variable dari numberB maka numberA akan ikut berubah
	*numberB = 10
	//fmt.Println(*numberB)
	//fmt.Println(numberA)

	//	define with var
	var numberC int = 11
	var numberD *int = &numberC

	fmt.Println(numberC)
	fmt.Println(*numberD)

	numberA = 1
	numberC = 2

	//fmt.Println(numberC)
	//fmt.Println(*numberD)
	//fmt.Println(*numberB)
	//fmt.Println(numberA)

	//	demonstrasi penggunaan pointer
	numberE := 5
	fmt.Println("Alamat memory awal :", &numberE)
	fmt.Println("Nilai awal :", numberE)
	change(&numberE, 10)
	fmt.Println("Alamat memory akhir :", &numberE)
	fmt.Println("Nilai akhir :", numberE)

	//	learn pointer in struct
	pointerStruct.PointerInStruct()

	//	run quiz
	quiz.RunQuiz()
}

//change mencoba mendefinisikan penggunaan pointer
func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory dalam fungsi :", old)
	fmt.Println("In function:", *old)
}

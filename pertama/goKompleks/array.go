package goKompleks

import "fmt"

func GoKompleksArray() {
	// default value dari array adalah menyesuaikan tipe data utamanya
	// misal int => 0, string => empty string ""
	var languages [5]string
	fmt.Println(languages)

	// cara assign value ke array di go
	// index terbesar = index terakhir, diluar itu bakal out of bound
	languages[0] = "GO"
	languages[1] = "Ruby"
	languages[2] = "JavaScript"
	languages[3] = "C#"
	languages[4] = "Python"

	// solusi ke 2 array: dengan one liner
	languages = [5]string{"Ruby", "C#", "Java", "PHP", "JavaScript"}
	fmt.Println(languages)

	// menghitung panjang array
	length := len(languages)
	fmt.Println(length)

	// ... berfungsi mengisi panjang array sesuai yang dimasukkan
	carType := [...]string{"Jeep", "Sedan", "Truk"}
	fmt.Println(len(carType))

	// cara mengakses array dengan menggunakan for each
	for i, lang := range languages {
		fmt.Println("Index: ", i, ", Value: ", lang)
	}
}

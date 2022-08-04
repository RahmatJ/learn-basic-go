package goKompleks

import "fmt"

func GoKompleksSlice() {
	// inisialisasi slice mirip dengan array, tapi tidak diberi panjang spesifik
	var gamingConsoles []string

	// cara untuk menambah elemen baru adalah dengan cara append data baru ke data existing
	gamingConsoles = append(gamingConsoles, "PlayStation 4")
	gamingConsoles = append(gamingConsoles, "Nintendo Switch")

	fmt.Println(gamingConsoles)

	// cara inisialisasi slice
	carType := []string{"Sedan", "Jeep"}

	fmt.Println(carType)

	// cara akses elemen dalam slice sama dengan array

	for _, data := range gamingConsoles {
		fmt.Println(data)
	}
}

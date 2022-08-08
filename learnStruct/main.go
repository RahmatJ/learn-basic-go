package main

import "fmt"

//mirip dengan struktur table user
type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
	IsActive  bool
}

func main() {
	//basic initiate user
	//jika kosong, maka diisi dengan default value tiap tipe data
	user := User{}
	//cara pengaksesan dan pengubahan pada attribute User struct
	user.ID = 1
	user.FirstName = "Reisalin"
	user.LastName = "Stout"
	user.email = "Ryza@koei.tecmo.com"
	user.IsActive = true

	//	other way to initate
	// by defining struct key and its value
	user1 := User{
		ID:        2,
		FirstName: "Sophie",
		LastName:  "Neuenmueller",
		email:     "sophie@atelier.com",
		IsActive:  true,
	}

	//	another way to initiate
	// without defineing struct key, but must be oredered sequentially
	user2 := User{
		3,
		"Ayesha",
		"Sophie",
		"ayesha@atelier.com",
		true,
	}

	display1 := displayUser(user1)
	display := displayUser(user)
	display2 := displayUser(user2)
	fmt.Println(display)
	fmt.Println(display1)
	fmt.Println(display2)
}

func displayUser(user User) string {
	//Sprintf mengembalikan string yang telah diformat
	//TODO: explore more about this
	result := fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.LastName)
	return result
}

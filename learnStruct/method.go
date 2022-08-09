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

//method adalah fungsi yang dimiliki suatu struct
func (user User) display() string {
	//Sprintf mengembalikan string yang telah diformat
	//TODO: explore more about this
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.email)
}

func displayUser(user User) string {
	//Sprintf mengembalikan string yang telah diformat
	//TODO: explore more about this
	return fmt.Sprintf("Name: %s %s, Email: %s", user.FirstName, user.LastName, user.email)
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func (group Group) display() {
	//printf akan mencetak formatted string
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Printf("Member Count: %d\n", len(group.Users))
	//	loop for display all users
	fmt.Println("Users: ")
	for _, user := range group.Users {
		fmt.Println(displayUser(user))
	}
}

func displayGroup(group Group) {
	//printf akan mencetak formatted string
	fmt.Printf("Name: %s\n", group.Name)
	fmt.Printf("Member Count: %d\n", len(group.Users))
	//	loop for display all users
	fmt.Println("Users: ")
	for _, user := range group.Users {
		fmt.Println(displayUser(user))
	}
}

package conditional

import "fmt"

func ConditionalIf() {
	age := 10

	if age > 10 {
		fmt.Println("Boleh bermain game")
	} else {
		fmt.Println("Belum boleh bermain game")
	}
}

func ConditionalElseIf() {
	score := 80
	//inisialisasi variable dengan var
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)
}

func ConditionalSwitch() {
	number := 5
	switch number {
	case 1:
		fmt.Println("Satu")
	case 2:
		fmt.Println("Dua")
	case 3:
		fmt.Println("Tiga")
	default:
		fmt.Println("DEFAULT")
	}
}

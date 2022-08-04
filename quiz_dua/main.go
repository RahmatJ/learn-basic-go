package main

import "fmt"

func main() {
	text := "Golang The Best Language"
	//	print only even index
	for index, letter := range text {
		if index%2 == 0 {
			fmt.Println("index: ", index, " letter: ", string(letter))
		}
	}
	//	print only vocal
	for _, letter := range text {
		if string(letter) == "a" || string(letter) == "i" || string(letter) == "u" || string(letter) == "e" || string(letter) == "o" {
			fmt.Println("Letter: ", string(letter))
		}
	}
}

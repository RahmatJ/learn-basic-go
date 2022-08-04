package main

import "fmt"

func main() {
	printMyResult()
	printMyResultWithInput("saya")
	printMyResultWithInput("sedang belajar")
	printMyResultWithInput("golang")
}

//simple function tanpa input
func printMyResult() {
	fmt.Println("Saya sedang belajar Go")
}

//simple function dengan input
func printMyResultWithInput(sentence string) {
	fmt.Println(sentence)
}

// TODO: simple function dengan input dan output

package quiz

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

//1. Buat method AddGame untuk menambahkan games ke struct Gamer
func (gamer *Gamer) AddGame(name string) {
	gamer.Games = append(gamer.Games, name)
}

func RunQuiz() {
	gamer := Gamer{Name: "Zelda"}

	//tambah game disini
	gamer.AddGame("Mario Bros")
	gamer.AddGame("Atelier Iris")
	gamer.AddGame("Atelier Iris 2")
	gamer.AddGame("Atelier Iris 3")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}

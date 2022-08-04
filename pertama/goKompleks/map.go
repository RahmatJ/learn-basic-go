package goKompleks

import "fmt"

func GoKompleksMap() {
	// membuat map dengan key string, dan value int
	var myMap map[string]int
	//initialize map, defaultnya nil
	myMap = map[string]int{}

	myMap["ruby"] = 9
	myMap["go"] = 9
	myMap["typeScript"] = 10

	fmt.Println(myMap)
	// cara pengaksesan value
	fmt.Println(myMap["go"])

	// cara inisialisasi langsung
	myLang := map[string]string{
		"ruby": "is beautiful",
		"go":   "is super fast",
	}

	fmt.Println(myLang)

	// looping map
	for key, value := range myMap {
		fmt.Println("Key: ", key, ", value: ", value)
	}

	fmt.Println("======================================")

	// delete key and value dari map
	delete(myMap, "typeScript")

	for key, value := range myMap {
		fmt.Println("Key: ", key, ", value: ", value)
	}

	// check ada tidaknya key di suatu map
	_, isAvailable := myMap["python"]
	fmt.Println(isAvailable)
}

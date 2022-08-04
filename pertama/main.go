package main

import (
	"fmt"
	"pertama/calculation"
	"pertama/conditional"
	"pertama/goKompleks"
	"pertama/loop"
)

func main() {
	fmt.Println("Hello dude!")
	//tidak perlu import Function kalau berada dalam 1 package, meskipun berbeda file
	sentence := TestAja()
	fmt.Println(sentence)

	result := calculation.Add(9, 9)
	fmt.Println(result)

	//	conditional
	conditional.ConditionalIf()
	conditional.ConditionalElseIf()
	conditional.ConditionalSwitch()

	//	Loop
	loop.LoopFor()
	loop.LoopWhile()
	loop.LoopString()

	//	array
	goKompleks.GoKompleksArray()

	// slice
	goKompleks.GoKompleksSlice()

	// map
	goKompleks.GoKompleksMap()

	//slice of map
	goKompleks.GoKompleksSliceOfMap()
}

package pointerStruct

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float32
}

//method version of graduate
func (student *Student) graduate() {
	student.Name = student.Name + ", S.T"
}

func PointerInStruct() {
	student := Student{1, "Rahmat", 3.8}

	fmt.Println(student.Name)

	student.graduate()

	fmt.Println(student.Name)
}

func graduate(student *Student) {
	student.Name = student.Name + ", S.T"
}

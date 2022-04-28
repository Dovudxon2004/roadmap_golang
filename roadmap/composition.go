package roadmap

import "fmt"

type Dog struct {
	color string
	breed string
	age   int
}

type Cat struct {
	Dog
}

func (d Dog) speak() {
	fmt.Println("Woof")
}

func (c Cat) speak() {
	fmt.Println("Meow")
}

func Composition() {
	chapa := Dog{"orange", "simple", 4}
	chase := Cat{Dog{"black", "Swedish", 3}}
	chapa.speak()
	chase.Dog.speak()
}

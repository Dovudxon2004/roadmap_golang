package main

import (
	"fmt"

	"github.com/Dovudxon2004/roadmap_golang/roadmap"
)

func main() {
	NextInt := roadmap.FibonacciWithClosure()
	for i := 0; i < 5; i++ {
		fmt.Println(NextInt())
	}
}

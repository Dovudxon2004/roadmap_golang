package roadmap

import (
	"fmt"
)

type car struct {
	color    string
	brand    string
	topSpeed int
	isNew    bool
}

func CarIndustry() {
	gentra := car{
		color: "black",

		topSpeed: 200,
		isNew:    true,
	}

	if gentra.brand == "chevrolet" {
		fmt.Println("Monopoly")
	} else {
		fmt.Println("you're good")
	}
}

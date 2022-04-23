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

func SliceOfStructs() {
	sl := []struct {
		name string
		cost int
	}{
		{"gentra", 150},
		{"malibu", 300},
		{"trailblazer", 600},
		{"tahoe", 900},
	}
	for _, v := range sl {
		if v.cost > 500 {
			fmt.Println(v.name)
		}
	}
}

package roadmap

import (
	"fmt"
	"math"
)

func HelloString() {

	pow := func(x, n, lim float64) float64 {
		if v := math.Pow(x, n); v < lim {
			return v
		}
		return lim
	}
	fmt.Println(pow(2, 5, 31))
	fmt.Println(pow(2, 5, 33))
}

package roadmap

import (
	"math"
)

func FindRoot(n float64) float64 {
	m := n / 2
	for math.Abs((m*m - n)) > 0.01 {
		m -= (m*m - n) / (2 * m)
	}
	return m
}

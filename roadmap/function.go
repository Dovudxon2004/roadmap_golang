package roadmap

func Increment() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}

}

func Fibonacci(x int) int {
	if x == 1 || x == 2 {
		return 1
	}

	return Fibonacci(x-2) + Fibonacci(x-1)
}

func FibonacciWithClosure() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

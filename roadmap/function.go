package roadmap

import "fmt"

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
func FizzBuzz(n int) []string {
	arr := make([]string, n)
	for i := range arr {
		if (i+1)%3 == 0 || (i+1)%5 == 0 {
			arr[i] = "FizzBuzz"
		} else if (i+1)%3 == 0 {
			arr[i] = "Fizz"
		} else if (i+1)%5 == 0 {
			arr[i] = "Buzz"
		} else {
			arr[i] = fmt.Sprintf("%v", i+1)
		}
	}
	return arr
}

package roadmap

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func order(c chan int) {
	n := 0
	for i := 1; i <= 10; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("order #%v\n", i)
		n = i
	}
	c <- n
	wg.Done()
}
func refund(c chan int) {
	n := 0
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond * 1000)
		fmt.Printf("refund #%v\n", i)
		n = i
	}
	c <- n
	wg.Done()
}

func Concurrency() {
	c := make(chan int)
	wg.Add(2)
	go order(c)
	go refund(c)
	x, y := <-c, <-c
	wg.Wait()
	fmt.Printf("Number of orders: %v, and Number of refunds %v\n", x, y)
}

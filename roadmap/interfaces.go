package roadmap

import (
	"errors"
	"fmt"
	"math"
)

//interface for two objects both with method area(). From Youtube, Tech with Tim, Interfaces
type circle struct {
	radius float32
}
type rect struct {
	width  int
	height int
}

func (c circle) area() float32 {
	return float32(c.radius * c.radius * math.Pi)
}
func (r rect) area() float32 {
	return float32(r.height * r.width)
}

type areaFinder interface {
	area() float32
}

func Interfaces1() {
	c1 := circle{5}
	r1 := rect{4, 5}
	x := func(a areaFinder) float32 {
		y := a.area()
		return float32(y)
	}
	slice := []areaFinder{c1, r1}
	for _, v := range slice {
		fmt.Println(x(v))
	}
}

//Interface for a bank account. From Youtube, Code with RYan, Advanced Golang: Channels,Context and interfaces.
//  This is a pretty good channel

type BankAccount interface {
	GetBalance() int
	Deposit(i int)
	Withdraw(i int) error
}

// WellsFargo account
type WellsFargo struct {
	Balance int
}

func NewWellsFargo() *WellsFargo {
	return &WellsFargo{}
}
func (wf *WellsFargo) GetBalance() int {
	return wf.Balance
}
func (wf *WellsFargo) Deposit(i int) {
	wf.Balance += i
}
func (wf *WellsFargo) Withdraw(i int) error {
	NewBalance := wf.Balance - i
	if NewBalance < 0 {
		return errors.New("Insufficient funds")
	}
	wf.Balance -= i
	return nil
}

// Bitcoin account
type Bitcoin struct {
	Balance int
	fee     int
}

func NewBitcoin() *Bitcoin {
	return &Bitcoin{
		fee: 100,
	}
}

func (btc *Bitcoin) GetBalance() int {
	return btc.Balance
}

func (btc *Bitcoin) Deposit(i int) {
	btc.Balance += i
}

func (btc *Bitcoin) Withdraw(i int) error {
	NewBalance := btc.Balance - i - btc.fee

	if NewBalance < 0 {
		return errors.New("insufficient funds")
	}

	btc.Balance -= i + btc.fee

	return nil
}

// playing with the interface
func Interfaces2() {

	Accounts := []BankAccount{NewWellsFargo(), NewBitcoin()}

	for _, v := range Accounts {
		v.Deposit(500)

		if err := v.Withdraw(100); err != nil {
			fmt.Printf("Err: %d\n", err)
		}

		fmt.Println(v.GetBalance())
	}
}

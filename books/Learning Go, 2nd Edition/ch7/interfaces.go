package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Score int
type HighScore Score
type Employee Person

func main() {
	fmt.Println("Pointer Receiver")
	pointerReceiver()

	fmt.Println("Code Your Methods for nil Instances")
	checkforNil()

	fmt.Println("Methods Are Functions")
	methodsAreFunctions()

	fmt.Println("Type Declarations Aren’t Inheritance")
	Inheritance()

}

func Inheritance() {
	type Score int
	type HighScore Score
	type Employee Person

	var i int = 300
	var s Score = 100
	var hs HighScore = 200

	s = Score(i)
	hs = HighScore(s)
	fmt.Print(hs)
}

type Adder struct {
	start int
}

func methodsAreFunctions() {
	myAdder := Adder{start: 10}
	fmt.Println(myAdder.AddTo(5))

	f1 := myAdder.AddTo
	fmt.Println(f1(10))

	f2 := Adder.AddTo
	fmt.Println(f2(myAdder, 15)) //25

}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

// need to check
func checkforNil() {
	var it *IntTree
	it = it.Insert(5)
	fmt.Println(it.Contains(2))
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2)) // true
	fmt.Println(it.Contains(12))
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func pointerReceiver() {
	var c Counter
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	// Be aware that the rules for passing values to functions still apply.
	doUpdateWrong(c)
	fmt.Println("Update Wrong :", c.String())

	doUpdateRight(&c)
	fmt.Println("Update Right :", c.String())
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	// Go considers both pointer and value receiver methods to be in the method set for a pointer instance. For a value instance, only the value receiver methods are in the method set.
	fmt.Println("in doUpdateRight:", c.String())
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

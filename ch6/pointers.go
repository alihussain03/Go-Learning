package main

import (
	"fmt"
)

func main() {

	fmt.Println("First Pointer. Explain & and *")
	firstPointerExample()

	fmt.Println("Second Pointer. Explain more referecing operator")
	secondPointerExample()

	pointerUsingNewkeyword()

	fmt.Println("Pointers Indicate Mutable Parameters")
	mutable()

	ex1()
	ex2()
}

func firstPointerExample() {
	var x int32 = 10
	var y bool = true
	pointerX := &x       // store address of x
	pointerY := &y       ///store address of y
	var pointerZ *string // store nothing

	fmt.Println("PointerX", pointerX)
	fmt.Println("PointerY", pointerY)
	fmt.Println("PointerZ", pointerZ)

	fmt.Println(*pointerX) // prints 10
	z := 5 + *pointerX
	fmt.Println("PointerZ", z)
}

func secondPointerExample() {
	var x *int
	fmt.Println(x == nil) // prints true
	//fmt.Println(*x)       // it is nil and will go for panic

	x1 := 10
	var pointerToX *int = &x1
	fmt.Println(*pointerToX)

	var x2 = new(int)      // new operator pointer with zero value
	fmt.Println(x2 == nil) // prints false
	fmt.Println(*x2)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}

	//	need to come back to this later
	p := person{
		FirstName:  "Pat",
		MiddleName: makePointer("Perry"), // This line won't compile
		LastName:   "Peterson",
	}

	fmt.Println(*p.MiddleName)
}

func mutable() {
	// if a pointer is passed to a function, the function gets a copy of the pointer. This 		 still points to the original data, which means that the original data can be modified by the called function.
	var x *int
	failedUpdate(x)
	fmt.Println(x) // prints nil

	// update function
	y := 20
	update(&y)
	fmt.Println(y)
}

func failedUpdate(g *int) {
	x := 10
	g = &x
}

func update(px *int) {
	*px = 20
}

func makePointer[T any](t T) *T {
	return &t
}

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func ex1() {
	per1 := MakePerson("Ali", "Hussain", 28)
	fmt.Println(per1.FirstName)

	per2 := MakePersonPointer("Ali", "Hussain", 28)
	fmt.Println(per2.FirstName)

}

func MakePerson(FirstName string, LastName string, age int) Person {
	return Person{FirstName: FirstName, LastName: LastName, Age: age}
}

func MakePersonPointer(FirstName string, LastName string, age int) *Person {
	return &Person{FirstName: FirstName, LastName: LastName, Age: age}
}

func ex2() {
	var strSlice = []string{"hello", "Ali", "Hussain"}
	fmt.Println("My slice before call update Slice", strSlice)
	UpdateSlice(strSlice, "World")
	fmt.Println("My slice after call update Slice", strSlice)

	fmt.Println("My slice before call grow Slice", strSlice)
	growSlice(strSlice, "World")
	fmt.Println("My slice after call grow Slice", strSlice)
}

func UpdateSlice(mySlice []string, str string) {
	mySlice[len(mySlice)-1] = mySlice[len(mySlice)-1] + " " + str
	//fmt.Println(mySlice)
}

func growSlice(mySlice []string, str string) {
	mySlice = append(mySlice, str)
	//fmt.Println(mySlice)
}

func pointerUsingNewkeyword() {
	var x = new(int)
	fmt.Println(x == nil) // prints false
	fmt.Println(*x)       // prints 0
}

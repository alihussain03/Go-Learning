package main

import "fmt"

var (
	v0 int
	v1 float64
	v2 bool
	v3 string
)

func main() {
	fmt.Println("Basic Data types")
	basicDataTypes()

	fmt.Println("\nPointers")
	pointers()

	fmt.Println("\nFunctions")
	functions()

	fmt.Println("\nifElse")
	ifElse()

	fmt.Println("\nError Handling")
	area, error := areaUsingErrorHanlding(2, 4)
	if error == nil {
		fmt.Print(area)
	} else {
		fmt.Print(error)
	}

}

/* Basic Data types */
func basicDataTypes() {

	// var greeting String
	// var greeting = "Hello Linkedin!"
	// greeting := "Hello Linkedin!"

	Greeting := "Hello Linkedin!" // short hand declaration
	fmt.Println(Greeting)

	fmt.Printf("[v0]: variable type = %T; value = %v;\n", v0, v0)
	fmt.Printf("[v1]: variable type = %T; value = %v;\n", v1, v1)
	fmt.Printf("[v2]: variable type = %T; value = %v;\n", v2, v2)
	fmt.Printf("[v3]: variable type = %T; value = %v;\n", v3, v3)
}

func pointers() {
	var ptr *string
	Greeting := "Hello Linkedin!"
	ptr = &Greeting
	fmt.Println("Greeting", Greeting)
	fmt.Println("Greeting Address", Greeting)
	fmt.Println("Greeting ptr", *ptr)

}

func functions() {
	g := add(2, 3)

	g1 := func(x int, y int) int {
		return x - y
	}

	fmt.Println(g)
	fmt.Println(g1(3, 2))

	//passing pointer to function
	negate(&g)
	fmt.Println("Value after negate functino ", g)

	x := 2
	y := 3
	area, circum := mutipleReturnValues(2, 3)
	fmt.Printf("rectangle (%v,%v): area=%v\n", x, y, area)
	fmt.Printf("rectangle (%v,%v): circumference=%v\n", x, y, circum)
}

func negate(x *int) {
	neg := -*x
	*x = neg
	fmt.Println("Value in negate functino ", *x)
}

func add(x int, y int) int {
	return x + y
}

func mutipleReturnValues(x int, y int) (area int, circumF int) {
	area = x * y
	circumF = 2 * (x + y)
	return
}

func ifElse() {
	parityCheck(2)

	// x := 2
	if x := 2; x%3 == 0 {
		fmt.Printf("%v is even.\n", x)
	} else {
		fmt.Printf("%v is odd.\n", x)
	}
}

func parityCheck(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	} else {
		fmt.Printf("%v is odd.\n", x)
	}
}

func areaUsingErrorHanlding(x int, y int) (*int, error) {
	if x%2 == 0 || y%2 == 0 {
		return nil, fmt.Errorf("Zero inputs: [%v,%v]", x, y)
	} else {
		area := x * y
		return &area, nil
	}
}

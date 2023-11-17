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

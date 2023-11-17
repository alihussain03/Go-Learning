package main

import (
	"fmt"
)

func main() {

	//slice
	fmt.Println("Slice Declaration")
	sliceDeclaration()

	fmt.Println("Slice functions")
	sliceFunctions()

	fmt.Println("Create Slice using make")
	createSliceUsingMake()

	// need to recheck
	fmt.Println("slicing slices")
	slicingSlices()

	fmt.Println("Converting arrary to slice and vice versa")
	converting()

	// Maps
	fmt.Println("Map Declaration")
	mapDeclaration()

	mapFunctions()

	// Struct
	fmt.Println("Struct Declaration")
	structDeclaration()

	// Struct comparison
	fmt.Println("Struct comparison")
	structcomparison()

	//	Exercise
	ex1()
	ex2()
	ex3()

}

func sliceDeclaration() {
	var slice1 = []int{10, 20, 30, 40}
	fmt.Println(slice1)

	var slice2 = [10]int{1, 2, 5: 55, 45}
	slice2[2] = 34
	fmt.Println(slice2)

	var nillSlice []int
	fmt.Println(nillSlice == nil)

	/*
		x := []int{1, 2, 3, 4, 5}
		y := []int{1, 2, 3, 4, 5}
		z := []int{1, 2, 3, 4, 5, 6}
		s := []string{"a", "b", "c"}
		fmt.Println(slices.Equal(x, y)) // prints true
		fmt.Println(slices.Equal(x, z)) // prints false
		fmt.Println(slices.Equal(x, s))
	*/
}

func sliceFunctions() {
	var nillSlice []int
	fmt.Printf("Slice length is %d \n", len(nillSlice))

	nillSlice = append(nillSlice, 3, 1, 2, 3, 4, 5)
	fmt.Printf("After adding, Slice length is %d \n", len(nillSlice))

	fmt.Printf("Slice capacity is %d \n", cap(nillSlice))

	// clear(nillSlice) // not working

	//Copying slices
	x1 := []int{1, 2, 3, 4}
	y1 := make([]int, 4)
	num := copy(y1, x1)
	fmt.Println(y1, num)

}

func createSliceUsingMake() {
	sliceUsingMake := make([]int, 5) // make([]int, 5,10) initial length and capacity
	sliceUsingMake = append(sliceUsingMake, 3)
	fmt.Printf("sliceUsingMake %d \n", sliceUsingMake)
}

func slicingSlices() {
	// slicing slices
	x := []string{"a", "b", "c", "d"}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
	fmt.Println("d:", d)
	fmt.Println("e:", e)

	// slices variable share the memory, after changing the value
	y = x[:2]
	z = x[1:]
	x[1] = "y"
	y[0] = "x"
	z[1] = "z"
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func converting() {
	var array = [4]int{1, 2, 3, 4}
	var sliceFromArray = array[:]
	fmt.Printf("sliceFromArray %d \n", sliceFromArray)

	x2 := [4]int{5, 6, 7, 8}
	y2 := x2[:2]
	z2 := x2[2:]

	fmt.Printf("sliceFromArray using sub array %d \n", y2)
	fmt.Printf("sliceFromArray using sub array %d \n", z2)

	xSlice := []int{1, 2, 3, 4}
	xArray := [4]int(xSlice)
	smallArray := [2]int(xSlice)
	xSlice[0] = 10
	fmt.Println(xSlice)
	fmt.Println(xArray)
	fmt.Println(smallArray)
}

func mapDeclaration() {
	var nilMap map[string]int
	fmt.Printf("Nil Map  %d \n", len(nilMap))

	zeroValueMap := map[string]int{}
	fmt.Printf("zero value Map  %d \n", len(zeroValueMap))

	teams := map[string][]string{"Orcas": []string{"Fred", "Ralph", "Bijou"}, "Lions": []string{"Sarah", "Peter", "Billie"}, "Kittens": []string{"Waldo", "Raul", "Ze"}}

	fmt.Printf("teams Map  %d \n", teams)

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Printf("Read map value  %d \n", m["hello"])
	fmt.Printf("Read not existing value from map %d \n", m["hello1"])
}

func mapFunctions() {

	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	fmt.Printf("Read map value  %d \n", m["hello"])
	fmt.Printf("Read not existing value from map %d \n", m["hello1"])

	v, ok := m["hello"]
	fmt.Println(v, ok)

	delete(m, "hello")
	fmt.Println("After deleting hello ", m)

	// fmt.Printf("After deleting hello from map %d \n", clear(m))
	// fmt.Println(maps.Equal(m, n))

	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	//fmt.Printf("Length of vals is: and length of intset is %d %d \n", vals, intSet)
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

}

func structDeclaration() {
	type person struct {
		name string
		age  int
		pet  string
	}

	var ali person // create zero object of struct person
	fmt.Printf(ali.name)

	bob := person{
		age:  30,
		name: "Beth",
	} // create  object of struct using literal
	fmt.Println(bob.name)

	bob.name = "Bob"
	fmt.Println(bob.name)

	julia := person{
		"Julia",
		40,
		"cat",
	}
	fmt.Println(julia)
}

func structcomparison() {
	type firstPerson struct {
		name string
		age  int
	}
	f := firstPerson{
		name: "Bob",
		age:  50,
	}
	var g struct {
		name string
		age  int
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
}

func ex1() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	fmt.Println(greetings)

	firstSlice := greetings[:2]
	fmt.Println(firstSlice)

	secondSlice := greetings[1:4]
	fmt.Println(secondSlice)

	thirdSlice := greetings[2:4]
	fmt.Println(thirdSlice)
}

func ex2() {
	message := "Hi 👩 and 👨"
	runes := []rune(message)
	fmt.Println(string(runes[3]))
}

func ex3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	firstEmployee := Employee{
		"ali",
		"hussain",
		3,
	}
	fmt.Println(firstEmployee)

	secondEmployee := Employee{
		firstName: "ali",
		lastName:  "hussain",
		id:        3,
	}

	var thirdEmployee = Employee{}
	thirdEmployee.firstName = "Ali"
	thirdEmployee.lastName = "Ali"
	thirdEmployee.id = 3
	fmt.Println(firstEmployee)
	fmt.Println(secondEmployee)
	fmt.Println(thirdEmployee)
}

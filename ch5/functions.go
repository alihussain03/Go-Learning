package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	fmt.Println("First functino call with Struct")
	funcWithStruct()

	fmt.Println("\nVariadic Func Call")
	variadicFuncCall()

	fmt.Println("\nMultiple return values Func Call")
	funcWithMultipleReturnValues()

	fmt.Println("\nfuntion are values. Function can be assigned to variabels")
	funcAreValeus()

	fmt.Println("\nClousers examples Inner function is called as clousers.")
	clousers()

	fmt.Println("\nGo functions call by value.")
	callByValue()

	fmt.Println("\nExercise1 pending	")

	fmt.Println("\nExercise2")
	exercise2()

	exercise3()

}

type person struct {
	age  int
	name string
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) {
	fmt.Println(opts.FirstName)
}

func funcWithStruct() {
	MyFunc(MyFuncOpts{
		FirstName: "Samuli",
		Age:       50,
	})
	MyFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}

func variadicFuncCall() { // variadic function can receive one or more arguments
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))
}

func funcWithMultipleReturnValues() {
	result, remainder, err := divAndRemainder(6, 9)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)

	x, y, z := divAndRemainderWithNameParams(5, 2)
	fmt.Println(x, y, z)
}

func funcAreValeus() {
	var myFuncVariable func(string) int
	myFuncVariable = f1
	res := myFuncVariable("Hello")
	fmt.Println(res)

	myFuncVariable = f2
	res = myFuncVariable("Hello")
	fmt.Println(res)

	funcAreValeus1()
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func divAndRemainderWithNameParams(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result = 20
	remainder = 30
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, v := range a {
		total += int(v)
	}
	return total
}

func add(i int, j int) (int, error) { return i + j, nil }

func sub(i int, j int) (int, error) { return i - j, nil }

func mul(i int, j int) (int, error) { return i * j, nil }

func div(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

func funcAreValeus1() {
	//function are values

	type opFuncType func(int, int) (int, error)
	//var opMap = map[string]func(int, int) int{
	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}

func clousers() {
	a := 20
	f := func() {
		fmt.Println(a)
		a := 30 // a can be shadow
		fmt.Println(a)
	}
	f()
	fmt.Println(a)

	//sort slice using clousers
	passClouerasFunctionArgument()

	//return Clouer From Function
	twoBase := returnClouerFromFunction(2)
	threeBase := returnClouerFromFunction(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

func passClouerasFunctionArgument() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobdaughter", 23},
		{"Fred", "Fredson", 18},
	}
	fmt.Println(people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
}

func returnClouerFromFunction(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func callByValue() {

	p := person{}
	i := 2
	s := "Hello"
	modifyFails(i, s, p)
	fmt.Println(i, s, p)

	m := map[int]string{
		1: "first",
		2: "second",
	}

	modMap(m)
	fmt.Println(m)

	s1 := []int{1, 2, 3}
	modSlice(s1)
	fmt.Println(s1)

}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func fileLen(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	data := make([]byte, 2048)
	total := 0
	for {
		count, err := f.Read(data)
		total += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return total, nil
}

func exercise2() {
	if len(os.Args) < 2 {
		return
	}
	count, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}

func exercise3() {
	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}

func prefixer(prefix string) func(body string) string {
	return func(body string) string { return prefix + " " + body }
}

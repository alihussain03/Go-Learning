package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("ifelse example")
	ifelse()

	//loops
	fmt.Println("Loops example")
	loops()

	//Labeling Statement
	fmt.Println("Labeling Statement in for loop")
	labelingForLoop()

	//Switch statement
	fmt.Println("Labeling Statement in for loop")
	switchStatement()

	ex()

}

func ex() {
	var intSlice = []int{}
	for i := 0; i < 100; i++ {
		randomNumber := rand.Intn(101)
		intSlice = append(intSlice, randomNumber)
	}

	for _, v := range intSlice {
		switch {
		case v%6 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}

	var total = 0
	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Print(total)
	}

	//fmt.Print(intSlice)
}

func ifelse() {
	x := 10
	if x > 5 {
		x, y := 5, 20
		fmt.Println(x, y)
	}
	fmt.Println(x)

	if n := rand.Intn(10); n == 0 { // special scope is very handy. It lets you create variables that are available only where they are needed
		fmt.Println("That's too low")
	} else if n > 5 {
		fmt.Println("That's too big:", n)
	} else {
		fmt.Println("That's a good number:", n)
	}
}

func loops() {
	for i := 0; i < 10; i++ {
		//fmt.Println(i)
	}

	i := 1
	for i < 100 {
		//fmt.Println(i)
		i = i * 2
	}

	loopOverCompositeType()

}

func loopOverCompositeType() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// only values, skip keys
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// Only keys
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	//print map
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}

	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	loopOverString()

}

// TODO need to check one time
func loopOverString() {
	//loop over string
	//  iterating over a string with a for-range loop. It iterates over the runes, not the bytes
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
}

func labelingForLoop() {
	samples1 := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples1 {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
}

func switchStatement() {

	simpleSwitch()

	fmt.Println("Switch with break")
	switchWithBreak()

	fmt.Println("Blank Switch")
	blankSwitch()

}

func simpleSwitch() {

	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		// size := len(word)
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			wordLen := len(word)
			fmt.Println(word, "is exactly the right length:", wordLen)
		case 6, 7, 8, 9:

		default:
			fmt.Println(word, "is a long word!")
		}
	}

}

func switchWithBreak() {

	//switch statement with break
	// If you have a switch statement inside a for loop, and you want to break out of the for loop, put a label on the for statement and put the name of the label on the break
loop:
	for i := 0; i < 10; i++ {
		switch i {
		case 0, 2, 4, 6:
			fmt.Println(i, "is even")
		case 3:
			fmt.Println(i, "is divisible by 3 but not 2")
		case 7:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

}

func blankSwitch() {
	wordsSwitch := []string{"hi", "salutations", "hello"}
	for _, word := range wordsSwitch {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

}

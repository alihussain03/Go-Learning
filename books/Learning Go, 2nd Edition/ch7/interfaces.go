package main

import (
	"fmt"
	"io"
	"os"
	"sort"
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

	fmt.Println("\nCode Your Methods for nil Instances")
	checkforNil()

	fmt.Println("\nMethods Are Functions")
	methodsAreFunctions()

	fmt.Println("\nType Declarations Aren’t Inheritance")
	Inheritance()

	fmt.Println("\nUse of iota")
	iotA()

	fmt.Println("\nUse Embedding for Composition")
	useEmbeddingForComposition()

	fmt.Println("\nFirst Interface")
	firstInterface()

	fmt.Println("\nExercise")
	ex()
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}
	if _, ok := l.Teams[team2]; !ok {
		return
	}
	if score1 == score2 {
		return
	}
	if score1 > score2 {
		l.Wins[team1]++
	} else {
		l.Wins[team2]++
	}
}

func (l League) Ranking() []string {
	names := make([]string, 0, len(l.Teams))
	for k := range l.Teams {
		names = append(names, k)
	}
	sort.Slice(names, func(i, j int) bool {
		return l.Wins[names[i]] > l.Wins[names[j]]
	})
	return names
}

func ex() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)

	RankPrinter(l, os.Stdout)
}

type Incrementer interface {
	Increment()
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	results := r.Ranking()
	for _, v := range results {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

func firstInterface() {
	var myStringer fmt.Stringer
	var myIncrementer Incrementer
	pointerCounter := &Counter{}
	valueCounter := Counter{}

	myStringer = pointerCounter    // ok
	myStringer = valueCounter      // ok
	myIncrementer = pointerCounter // ok
	// myIncrementer = valueCounter   // compile-time error!

	fmt.Println(myStringer, myIncrementer)
}

type Employeee struct {
	Name string
	ID   string
}

type Manager struct {
	Employeee
	Reports []Employeee
}

func (e Employeee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

func useEmbeddingForComposition() {

	m := Manager{
		Employeee: Employeee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employeee{},
	}
	fmt.Println(m.ID) // prints 12345
	fmt.Println(m.Description())

	//Embedding is not inheritence
	fmt.Println("\nEmbedding is not inheritence")
	o := Outer{
		Inner: Inner{
			A: 10,
		},
		S: "Hello",
	}
	fmt.Println(o.Double())
}

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	return i.IntPrinter(i.A * 2)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

const (
	Field1 = 0
	Field2 = 1 + iota
	Field3 = 20
	Field4
	Field5 = iota
)

func iotA() {

	fmt.Println(Field1, Field2, Field3, Field4, Field5)
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

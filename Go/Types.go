package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Score int
type HighScore Score



type Adder struct {
	start int
}

var P1 Person
var i int = 300
var s Score = 100



func (p2 Person) String() string {
	return fmt.Sprintf("%s %s, age %d", p2.FirstName, p2.LastName, p2.Age)
}
func (p2 *Person) StringPointer() string {
	return fmt.Sprintf("%s %s, age %d is PointerReceiver", p2.FirstName, p2.LastName, p2.Age)
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

const (
	F1 Score = 3 << iota
	F2
	F3
)

func main() {
	/*var p = Person{
	FirstName: "Fred",
		LastName:"Fredson",
			Age: 52,
	}*/

	fmt.Println(F3)
	s = Score(100)
	fmt.Println(s)
	adder := Adder{start: 10}
	fmt.Println(adder.AddTo(5))

	to := Adder.AddTo  //method expression
	fmt.Println(to(Adder{start: 12},7))
	fmt.Println(P1.StringPointer())
	fmt.Println(P1.String())
}

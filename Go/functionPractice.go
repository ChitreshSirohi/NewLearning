package main

import (
	"fmt"
	"sort"
	"strconv"
)

func add(i int, j int) int {
	return i + j
}

func sub(i int, j int) int {
	return i - j
}

func mul(i int, j int) int {
	return i * j
}

func div(i int, j int) int {
	return i / j
}

var opMap = map[string]func(i int,j int) int{
	"+":add,
	"-":sub,
	"*":mul,
	"/":div,
}

func main()  {
	goByValueDemo()
	//mathOperation()
	//Anonymous function
	clojure()
	s := []string{"A","B","C"}
	for i:= range s {

		func(s string){
			fmt.Println("this is string",s)
		}(s[i])

	}

	one := addOne(2)

	fmt.Println(one(3))
}
func addOne(j int) func(int) int{
	return func(i int) int {
		defer fmt.Println("Annon")
		return i + j
	}
}

func clojure(){
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	} )
	fmt.Println(people)
}

func mathOperation(){
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _,expression := range expressions{
		//fmt.Println("k=",k," expression",expression)
		if len(expression) !=3{
			fmt.Println("error Continue")
			continue
		}

		num1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		operator := expression[1]

		opFun, ok := opMap[operator]
		if !ok {
			fmt.Println(err,"operation method not found")
			continue
		}

		num2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}

		result := opFun(num1, num2)
		fmt.Println(result)

	}
}

func goByValueDemo()  {
	sampleMap := map[int]string{
		1:"hello",
		2:"Hi",
	}
	sampleSlice := []string{
		"hello","hi",
	}
	fmt.Println("SampleSlice:",sampleSlice)
	modifySlice(sampleSlice)
	fmt.Println("ModSampleSlice:",sampleSlice)
	fmt.Println("SampleMap",sampleMap)
	modifyMap(sampleMap)
	fmt.Println("mod samplemap",sampleMap)

}
func modifyMap(sampleMap map[int]string)  {
	sampleMap[3] = "Haloa"
	sampleMap[1] = "hii"
	fmt.Println(sampleMap)

}

func modifySlice(sampleSlice []string){
	sampleSlice[0]="Hey"
	strings := append(sampleSlice, "4")
	fmt.Println(strings)
}



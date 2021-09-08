package main

import "fmt"

type LogicProvider struct {}
type LogicProvider1 struct {}

func (lp LogicProvider) Process(data string) string{
	fmt.Println("ProviderCalled")
	return data
}

func (lp LogicProvider1) Process1(data string) string{
	fmt.Println("Provider1Called")
	return data
}

type Logic interface {
	Process(data string) string
}
type Logic1 interface {
	Process1(data string) string
}

type Client struct {
	L Logic
	L1 Logic1
}

type MyInt int

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.L.Process("Hello Logic")
	c1 := Client{
		L1: LogicProvider1{},
	}
	c1.L1.Process1("Hello Logic1")

	var i interface{}
	var mine MyInt = 20
	i = mine
	myInt := i.(MyInt)  //type Conversion
	fmt.Println(myInt)
	checkTypes(i)
}

func checkTypes(i interface{})  {

	switch j := i.(type){
	case int:
		fmt.Println("int")
	case MyInt:
		fmt.Println("MyInt",j)
	default:
		fmt.Println("no idea")
	}

}
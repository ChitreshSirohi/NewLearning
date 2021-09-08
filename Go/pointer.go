package main

import "fmt"

func main()  {
	x := "hello"
	pointerToX := &x
	fmt.Println("Memory Address of pointerX",pointerToX)
	fmt.Println("Value of pointerX",*pointerToX)

	var y int
	fmt.Println("y vl ",y)
	pointerToY := &y
	fmt.Println("y pointer ",pointerToY)
	fmt.Println("y pointer ",*pointerToY)

	var pointerZ = new(int)
	fmt.Println(pointerZ)
	fmt.Println(*pointerZ)
	var f int
	failedUpdate(&f)
	fmt.Println(f)
	passUpdate(&f)
	fmt.Println(f)
}

func failedUpdate(g *int) {
	fmt.Println("g",*g)
	x:=10
	g = &x
}

func passUpdate(g *int) {
	fmt.Println("g",*g)
	x:=10
	*g = x
}

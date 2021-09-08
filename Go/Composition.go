package main

import "fmt"

type Employee struct {
	Name         string
	ID           string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

var mgr Manager = Manager{
	Employee{Name: "C", ID: "1"},
	[]Employee{},
}
func main(){
	fmt.Println(mgr.Description())
}
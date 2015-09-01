package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Avinash Kunuje", 24}
	z := Person{"sathyanarayana", 48}
	fmt.Println(a, z)
}
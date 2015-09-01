package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"avinash": {-45.222, 45.222},
	"kunuje":    {12.333, -55.123},
}

func main() {
	fmt.Println(m)
}

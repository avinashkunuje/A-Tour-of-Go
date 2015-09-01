package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"avinash": Vertex{
		40.68433, -74.39967,
	},
	"kunuje": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}

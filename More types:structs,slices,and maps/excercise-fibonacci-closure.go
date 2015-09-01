package main

import "fmt"

func fibonacci() func() int {
	fst, snd := 1, 1
	return func() int {
		res := fst
		fst, snd = snd, fst+snd
		return res
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
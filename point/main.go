package main

import "fmt" 

type xy struct {
	name string
	value int
}

func main() {
	x := xy{}
	y := xy{}

	x.name = "x"
	x.value = 42

	y.name = "y"
	y.value = 21

	fmt.Printf("x = %d, y = %d\n",x.value, y.value)
}


// test go run . 
// output x = 42, y = 21
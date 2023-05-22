package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "Loeksh Pathrabe",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.name = "Mr strawberry"
	// OR
	// (*p).name = "Mr strawberry"
}

package main

import "fmt"

func main() {
	x := "lokesh"
	if x == "Lokesh" {
		fmt.Println(x)
	} else if x == "lokesh" {
		fmt.Println("Lokesh Pathrabe", x)
	} else {
		fmt.Println("No lokesh")
	}
}

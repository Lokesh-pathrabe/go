package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello")
}

func foo() {
	defer func() {
		fmt.Println("Hello anonymous foo")
	}()
	fmt.Println("Foo Ran")
}

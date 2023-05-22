package main

import "fmt"

func main() {
	i1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n1 := foo(i1...)
	fmt.Println(n1)

	i2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n2 := bar(i2)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

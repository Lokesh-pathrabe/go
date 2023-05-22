package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "James",
		age:   32,
	}

	u2 := user{
		first: "Strawberry",
		age:   27,
	}

	u3 := user{
		first: "Mango",
		age:   21,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}

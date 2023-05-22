package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	data := `[{"First": "James","Last":"Bond","Age": 32, "Sayings":["hi","hello"]
	},{"First": "Lokesh","Last":"Pathrabe","Age": 21, "Sayings":["hey","yo"]
}]`

	fmt.Println(data)

	var people []person

	err := json.Unmarshal([]byte(data), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, p := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", p.First, p.Last, p.Age)
		for _, val := range p.Sayings {
			fmt.Println("\t\t", val)
		}
	}
}

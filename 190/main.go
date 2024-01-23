package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Dr.","Last":"No","Age":87}]`
	bs := []byte(s)

	people := []Person{}

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(people)

}

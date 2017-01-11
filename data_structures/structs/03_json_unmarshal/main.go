package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Parents []person
	Name    string
}

func main() {
	var abel person
	jsonStr := `{"Parents":[{"Parents":null,"Name":"Adam"},{"Parents":null,"Name":"Eve"}],"Name":"Abel"}`
	err := json.Unmarshal([]byte(jsonStr), &abel)
	if err != nil {
		fmt.Println("unable to unmarshal: ", err.Error())
	}
	fmt.Printf("Name: %s\nParents: %v, %v\n", abel.Name, abel.Parents[0], abel.Parents[1])
}

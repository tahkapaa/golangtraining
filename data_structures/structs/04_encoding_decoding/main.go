package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	Parents []person
	Name    string
}

func main() {
	p := createPerson()
	encoder := json.NewEncoder(os.Stdout)
	if err := encoder.Encode(p); err != nil {
		fmt.Println(err.Error())
	}
}

func createPerson() person {
	var abel person
	jsonStr := `{"Parents":[{"Parents":null,"Name":"Adam"},{"Parents":null,"Name":"Eve"}],"Name":"Abel"}`
	err := json.Unmarshal([]byte(jsonStr), &abel)
	if err != nil {
		fmt.Println("unable to unmarshal: ", err.Error())
	}
	return abel
}

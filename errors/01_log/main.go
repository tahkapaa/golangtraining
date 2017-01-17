package main

import (
	"log"
	"os"
)

func init() {
	f, err := os.Create("log.txt")
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(f)
}

func main() {
	_, err := os.Open("name")
	if err != nil {
		log.Println(err)
	}
}

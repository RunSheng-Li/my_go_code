package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

func main() {

	hobbies := []string{"sex", "girl", "fuck"}

	p := Person{
		Name:    "sam",
		Age:     20,
		Hobbies: hobbies,
	}

	fmt.Printf("%+v\n", p)

	jsonByteData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	jsonStringData := string(jsonByteData)
	fmt.Println(jsonStringData)

}

package main

import (
	"io/ioutil"
	"log"
)

func main() {
	s := "fuck~~~!!!"
	err := ioutil.WriteFile("1.txt", []byte(s), 0644)
	if err != nil {
		log.Fatal(err)
	}

}

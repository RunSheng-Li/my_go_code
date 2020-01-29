package main

import (
	"fmt"
	"time"
)

func receiver(c chan string) {
	for msg := range c {
		fmt.Println(msg)
	}
}

func main() {

	messages := make(chan string, 2)
	messages <- "hello"
	messages <- "sam"
	close(messages)
	time.Sleep(time.Second * 1)
	receiver(messages)

}

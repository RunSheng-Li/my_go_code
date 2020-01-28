package main

import (
	"fmt"
	"time"
)

func fuck() {
	for i := 0; i < 5; i++ {
		fmt.Println("fucking!!!")
		time.Sleep(100 * time.Millisecond)
	}

}

func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("singing!")
		time.Sleep(100 * time.Millisecond)

	}

}

func main() {
	go sing()
	go fuck()

	for {
		;
	}

}

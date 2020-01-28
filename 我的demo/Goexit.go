package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccccccccccccc")
	runtime.Goexit()
	fmt.Println("ddddddddddddddddddd")

}

func main() {

	go func() {

		defer fmt.Println("aaaaaaaaaa")
		test()
		fmt.Println("bbbbbbbbbbbbbbb")

	}()

	for {
		;
	}

}

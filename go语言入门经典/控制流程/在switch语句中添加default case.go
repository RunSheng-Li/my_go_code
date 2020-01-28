package main

import "fmt"

func main() {

	s := "c"

	switch s {
	case "a":
		fmt.Println("这是a！！！")

	case "b":
		fmt.Println("这是b！！！")

	default:
		fmt.Println("又不是a又不b")

	}

}

package main

import "fmt"

func main() {

	const (
		one = iota + 1
		two
		three
		four
		five
	)
	fmt.Println(one)
	fmt.Println(three)
	fmt.Println(five)

	const (
		Open = 1 << iota
		Close
		Pending
	)
	fmt.Println(Open)
	fmt.Println(Close)
	fmt.Println(Pending)

}

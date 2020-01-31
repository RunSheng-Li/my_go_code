package main

import "fmt"

func main() {

	var s = make([]int, 0)
	s = append(s, 1, 2, 3)
	var s2 = []int{4, 5, 6}
	s = append(s, s2...)
	fmt.Println("s", s)
	fmt.Println("s2", s2)

}

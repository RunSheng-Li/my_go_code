package main

import "fmt"

func main() {

	s := []int{}
	for i := 0; i < 10; i++{
		s = append(s, i)
		fmt.Println(len(s), cap(s))
	}

}

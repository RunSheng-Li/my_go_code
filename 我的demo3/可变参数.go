package main

import "fmt"

func hello(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret

}

func main() {
	fmt.Println(hello(1, 2, 3, 4, 5))
}

package main

import "fmt"

func anotherFunction(f func() string) string {

	return f()
}

func main() {
	fn := func() string {
		return "函数被调用"
	}
	fmt.Println(anotherFunction(fn))

}

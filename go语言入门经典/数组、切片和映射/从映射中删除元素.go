package main

import "fmt"

func main() {

	var haha = make(map[string]int)
	haha["1"] = 1
	haha["2"] = 2
	haha["3"] = 3
	fmt.Println(haha)
	delete(haha, "1")
	fmt.Println(haha)

}

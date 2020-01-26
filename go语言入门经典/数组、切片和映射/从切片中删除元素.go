package main

import "fmt"

func main() {

	var haha = make([] string, 3)
	haha[0] = "一"
	haha[1] = "二"
	haha[2] = "三"
	fmt.Println(len(haha))
	fmt.Println(haha)
	haha = append(haha[:2], haha[2+1:]...)
	fmt.Println(len(haha))
	fmt.Println(haha)

}

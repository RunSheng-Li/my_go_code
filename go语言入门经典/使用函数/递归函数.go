package main

import "fmt"

func feedMe(portion int, eaten int) int {
	eaten = portion + eaten
	if eaten >= 5 {
		fmt.Printf("劳资吃饱了！我吃了%d\n", eaten)
		return eaten
	}
	fmt.Printf("劳资还饿！我吃了%d\n", eaten)
	return feedMe(portion, eaten)

}

func main() {
	//fmt.Println(feedMe(1, 0))
	feedMe(1, 0)

}

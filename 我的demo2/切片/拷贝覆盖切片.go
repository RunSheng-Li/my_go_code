package main

import "fmt"

func main(){
	var a1 = []int{1,2,3,4,5,6}
	var a2 = []int{10,20,30}
	var n = copy(a1,a2)

	fmt.Println(a1,n)


}
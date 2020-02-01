package main

import "fmt"

func main(){

	var s = make([]int,0)
	s = append(s,10,20,30,40,50)
	for index,value:= range s{

		fmt.Println(index,value)

	}


}

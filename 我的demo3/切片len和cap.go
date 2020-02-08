package main

import "fmt"

//func main() {
//
//	s := []int{}
//	for i := 0; i < 10; i++{
//		s = append(s, i)
//		fmt.Println(len(s), cap(s))
//	}
//
//}

func main(){

	year := []string{"one","two","three","four","five","six","seven","eight","nine","ten","eleven","twelve"}
	Q2 := year[3:6]
	fmt.Println(Q2,len(Q2),cap(Q2))
	summer := year[5:8]
	fmt.Println(summer,len(summer),cap(summer))
	summer[0]="Unknow"
	fmt.Println(Q2)
	fmt.Println(year)


}
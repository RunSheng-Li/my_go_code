package main

import "fmt"

func main(){

	var a = [10]int{0,1,2,3,4,5,6,7,8,9}
	// 在数组的基础上截取，得到切片
	var as1 = a[2:8]
	fmt.Printf("type=%T,value=%v\n",as1,as1)

	// 截取的方式[头下标:尾下标],含头不含尾，不写头默认头=0，不写尾默认截取到最后一位
	var sa2 = a[:8]
	var sa3 = a[2:]
	var sa4 = a[:]
	fmt.Println(sa2,sa3,sa4)




}
package main

import "fmt"

func main() {

	// 三种方式创建出类型、值、长度、容量都相同的切片
	var s1 []int = []int{0, 1, 2}
	var s2 = []int{0, 1, 2}
	s3 := []int{0, 1, 2}

	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s1, s1, len(s1), cap(s1))
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s2, s2, len(s2), cap(s2))
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s3, s3, len(s3), cap(s3))

	// 创建长度和容量都为3的切片
	s4 := make([]int, 3)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s4, s4, len(s4), cap(s4))

	// 创建长度3，容量4的切片
	var s5 = make([]int, 3, 4)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d\n", s5, s5, len(s5), cap(s5))

}

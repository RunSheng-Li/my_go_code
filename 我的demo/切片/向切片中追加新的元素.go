package main

import "fmt"

func main() {

	// 容量每突破一次就翻倍
	// 每翻倍一次，切片就获得一片新的堆地址，每一个元素的地址也都随之发生变化
	// 切片名称的栈地址是不会变的
	var s = make([]int, 0)
	s = append(s, 0)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, s, len(s), cap(s), &s, &s[0])

	s = append(s, 2)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, s, len(s), cap(s), &s, &s[0])

	s = append(s, 3)
	fmt.Printf("type=%T,value=%v,len=%d,cap=%d,saddr=%p,elemaddr=%p\n", s, s, len(s), cap(s), &s, &s[0])

}

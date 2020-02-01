package main

import "fmt"

func main(){

	// 创建一个string为键，值为任意类型的map
	var mmp = make(map[string]interface{})

	// 把键值对丢入map中
	mmp["name"] = "土狗"
	mmp["sex"] = "male"
	mmp["hobby"] = "female"
	mmp["rmb"] = 0.05

	// 删除rmb对应的键值对
	delete(mmp,"rmb")
	fmt.Println(mmp)



}
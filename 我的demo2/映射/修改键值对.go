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

	// 修改特定的键值对
	mmp["hobby"] = [...]string{"抽烟","喝酒","泡妞"}
	mmp["rmb"] = -122222
	mmp["最爱的书籍"] = "雅思"

	fmt.Println(mmp)



}
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
	//mmp["最爱的书"] = "神雕侠侣"

	var value1 = mmp["sex"]
	fmt.Println("value1=",value1)

	var value2 = mmp["sex2"]
	fmt.Println("value2=",value2)


	// 带有校验的查询
	var value3,ok = mmp["最爱的书"]
	fmt.Println(value3,ok)
	if ok == true{

		fmt.Println(value3)
	} else {
		fmt.Println("劳资没有最爱的书")
	}

	// 遍历键值
	for key,value :=range mmp{
		fmt.Println(key,value)
	}

	// 遍历所有键
	for key := range mmp{
		fmt.Println(key)
	}



}

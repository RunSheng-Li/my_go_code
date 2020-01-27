//通过rune字面量，可将解释型字符串字面量分成多行，还可在其中包含制表符和其他格式选项
package main

import "fmt"

func main() {

	s := "我\nn是一句\n\t话"
	fmt.Println(s)

	//原始字符串字面量用反引号括起
	s1 := `劳
资也
		是一句话`

	fmt.Println(s1)

}

package main

import(
	"fmt"
)
//布尔类型的基本使用
func main(){
	var b = false
	fmt.Println("b =",b)
	//注意事项
	//1、bool类型占用存储空间是1个字节
	fmt.Println("布尔类型占用字节空间数",unsafe.Sizeof(b))
	//2布尔只有两种取值：1.ture 2.false


} 
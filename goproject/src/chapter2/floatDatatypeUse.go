package main

import(
	"fmt"
)
//浮点类型的基本使用
func main(){
	//尾数部分可能丢失，造成进度损失
	var i float32 =-1235651.123
	var j float64 =-123456789.12345678910
	fmt.Println("i =",i,"j =",j)
}
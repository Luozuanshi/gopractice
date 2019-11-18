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

	//浮点类型默认是float64类型
	var f =1.1
	fmt.Println("%T",f)

	//十进制数的两种表现形式
	num1 := 5.21
	num2 := .21 //=>0.21
	fmt.Println("num1 =",num1,"num2 =",num2)
	
	//科学计数法形式
	num3 := 10.123e2
	num4 := 1012.123E2
	num5 := 10.123e-2

	fmt.Println("num3 =",num3,"num4 =",num4,"num5 =",num5)
	

}
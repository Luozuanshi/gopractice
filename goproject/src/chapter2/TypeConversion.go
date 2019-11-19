package main

import (
	"fmt"
)

func main(){
	var i int =100
	var f float32 = float32(i) //不同类型转换，使用显示转换
	var i32 int32 = int32(i) //高精度->低精度
	var i64 int64 = int64(i) //低精度->高精度
	fmt.Printf("i =%d ,f =%f,int32 =%d,int64 =%d",i,f,i32,i64)

	//被转换的是变量存储的数据（值），变量本身的数据类型并没有变化
	fmt.Printf("\nI is type %T,\n",i)

	//在转换中比如将int64转换成int8{-128~127}，编译时不会报错
	//转换的结果按溢出处理
	var num1 = 999999
	var num2 = int8(num1) //自动判定数据类型，在初始化的时候 得到一个 int8数据类型，最终num2为int8数据类型
	fmt.Println("num2 =",num2)
	num2 =127
	fmt.Printf("num2 type is %T and value's is %d",num2,num2) //num2 type is int8 and value's is 127

}
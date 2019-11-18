package main

import "fmt"
//golang中整型类型的使用
func main(){
	//测试一下整数类型的范围 -128~127
	//其他的 int16 int32 int64 以此类推
	var i int8 = 127 //constant 128 overflows int8
	fmt.Println("i =",i)
	
	//无符号整形类型的范围 0~255
	//其他的 uint8 uint16 uint32 uint64以此类推
	var j uint8 = 1
	fmt.Println("j =",j) //constant -1 overflows uint8



	//其他整数类型的使用 int uint rune byte

	var k int = -666
	fmt.Println("k =",k) 
	var o uint = 666
	fmt.Println("o =",o) 
	var p byte = 255
	fmt.Println("p =",p) 
}
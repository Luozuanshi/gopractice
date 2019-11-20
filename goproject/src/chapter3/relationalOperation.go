package main

import (
	"fmt"
	_ "time"
)
//关系运算符的使用
func main () {
	
	//演示关系运算符的使用
	var a int =9
	var b int =8

	fmt.Println("a == b",a==b)
	fmt.Println("a != b",a!=b)
	fmt.Println("a > b",a>b)
	fmt.Println("a >= b",a>=b)
	fmt.Println("a < b",a<b)
	fmt.Println("a <= b",a<=b)
	flag := a==b
	fmt.Printf("a value's %v",flag)

}

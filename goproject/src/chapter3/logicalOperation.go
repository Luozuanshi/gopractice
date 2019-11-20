package main

import (
	"fmt"
	_ "time"
)
//关系运算符之逻辑运算的使用

func test() bool {
	fmt.Println("TEST...")
	return true
}
func main () {
	var v int = 10
	//演示短路与
	// v > 9 为true,test() 为true 因此先输出TEST...后输出ok...
	if v > 9 && test() {
		fmt.Println("ok...") 
	}
	// v > 11 为false,不判断后面的 test() 因此不输出
	if v > 11 && test() {
		fmt.Println("ok...") 
	}
	//演示短路或 
	//v > 9 为true 直接不继续判断 test() 因此直接输出 pk...
	if v > 9 || test() {
		fmt.Println("pk...") 
	}
	//v > 11 为false 继续判断 test() 因此先输出TEST...后输出pk...
	if v > 11 || test() {
		fmt.Println("pk...") 
	}

	


	//演示逻辑 && 操作
	var age int =40

	if age > 30 && age <50 {
		fmt.Println("ok1")
		} 
		
	if age >30 && age < 40 {
		fmt.Println("ok2")	
	}		

	//演示逻辑 || 操作
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}

	//演示逻辑 ! 操作
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}
}

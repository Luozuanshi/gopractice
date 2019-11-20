package main

import (
	"fmt"
	_ "time"
)
//关系运算符之逻辑运算的使用
func main () {
	
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

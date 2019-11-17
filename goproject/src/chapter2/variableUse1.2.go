//在Go语言中+的使用
package main

import "fmt"
	
func main(){
	
	var i int = 1
	var o int = 2
	j := i + o
	fmt.Println("j=",j)// + 号 求和运算
	
	str1 := "hello"
	str2 :="world"
	textresult := str1 +str2
	fmt.Println("textresult=",textresult)// + 号 字符连接操作
}


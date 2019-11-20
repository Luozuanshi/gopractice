package main

import "fmt" //ftm 包中提供格式化
//算术运算的注意事项
func main () {
	
	//golang中没有++i --i 这种操作
	var i int =10
	var j int
	// ++i // syntax error: unexpected ++, expecting
	// --i // syntax error: unexpected --, expecting

	//golang中自增自减只能独立的使用，不能这样使用
	// j = i++ //syntax error: unexpected ++ at end of statement
	// j = i-- //syntax error: unexpected -- at end of statement

	// if i++ > 10{ //syntax error: i++ used as value
	// 	fmt.Println("ok")
	// }

	fmt.Println("i = ,j = ",i,j)
}


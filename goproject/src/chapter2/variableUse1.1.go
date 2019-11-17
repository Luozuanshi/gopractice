//变量的使用细节
package main

import "fmt"
	
func main(){
	
	//该区域的数据值可以在同一类型范围内不断变化
	var i int =10
	i = 20
	i = 30
	fmt.Println("i =",i)

	i = 1.2 //只能在同一类型范围的变化，int 类型 不能赋值float类型 constant 1.2 truncated to integer

	//变量在 同一个作用域（在一个函数或者代码块）内不能重名
	//var i int =59
	// i := 100


}

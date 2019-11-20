package main

import "fmt" //ftm 包中提供格式化
//算术运算
func main () {
	
	// / 除法的使用
	//如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println("10/4=",10/4)

	//即使使用float32类型接收，得到的结果依然是整数部分
	var f float32
	f = 10 / 4
	fmt.Println("f = ",f)

	// 解决方案:运算的时候,数值类型必须有一个包含小数点
	f = 10.0 / 4
	fmt.Println("f = ",f)

	// % 取模的使用
	fmt.Println("---------------------")
	fmt.Println("10%3=",10%3)
	fmt.Println("-10%3=",-10%3)
	fmt.Println("10%-3=",10%-3)
	fmt.Println("-10%-3=",-10%-3)

	//++ -- 的使用
	var i int =10
	i++
	fmt.Println("i = ",i)
	i--
	fmt.Println("i = ",i)
}


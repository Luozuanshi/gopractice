package main

import(
	"fmt"
)
//其他运算符的使用
func main(){
	// * 和&的使用
	a := 100
	fmt.Println("a变量的内存地址",&a)
	var ptr *int =&a
	fmt.Println("通过指针i读取i变量值所引用地址的值为",*ptr)

	//求两个数的最大值
	var n1 = 10
	var n2 = 30
	var n3 = 20
	max := n1
	if n1 > n2 {
		max = n1
	}else{
		max = n2
	}
	//求三个数的最大值
	fmt.Println("n1,n2两个数最大值:",max)

	if max < n3 {
		max =n3
	}
	fmt.Println("n1,n2,n3三个数最大值:",max)

	
}
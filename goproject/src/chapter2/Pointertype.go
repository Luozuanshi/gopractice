package main

import(
	"fmt"
)
//指针类型的使用
func main(){
	var i int =100
	fmt.Printf("\ni变量的内存地址：%v",&i)
	fmt.Printf("\ni变量的值：%v",i)
	
	//*int 代表一个int指针类型 ，
	// &i，&表示取i的内存地址 
	// *ptr 表示取 ptr变量的所引用的内存地址的值
	var ptr *int = &i
	fmt.Printf("\nptr指针变量的内存地址：%v",&ptr)
	fmt.Printf("\nptr指针变量的值：%v",ptr)
	fmt.Printf("\nptr指针变量的值引用：%v",*ptr)

	
}
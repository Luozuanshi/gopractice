package main

import(
	"fmt"
)
//字符类型的基本使用
func main(){
	var c1 = 'a'
	var c2 byte = '0'
	//当我们直接输出了byte值，就直接输出了对应的ASCII值
	fmt.Println("c1 =",c1,"c2 =",c2)

	//如果希望输出对应的数值，需要格式化输出
	fmt.Println("%c%c",c1,c2)

	var c3 int = '罗'
	fmt.Print("对应字符%c 对应Unicode码值%d",c3,c3)

	//可以直接给一个变量赋一个数字，然后按照格式化输出%c，会输出对应的Unicode字符
	var c4 int = 1024
	fmt.Println("\n%c",c4)

	//字符类型是可以运算的，相当于一个整数，因为它都对应有Unicode码
	var n1 = 10 + 'a'
	fmt.Println("n1 = ",n1) 

} 
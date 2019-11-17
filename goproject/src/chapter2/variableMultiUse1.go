//一次性声明多个变量
package main

import "fmt"

//全局变量

//全局变量的一次性声明多个变量
var(
	n5 = 100
	n6 = 200
	n7 = "lisi" 
)
func main(){
	//方式一：一次性声明多个变量
	var n1,n2,n3 int
	fmt.Println("局部变量：\nn1 =",n1,"n2 =",n2,"n3 =",n3)
	
	// //方式二：一次性声明多个不同类型的变量
	// var name,age,color = "zhangsan",18,"red"
	// fmt.Println("name =",name,"age =",age,"color =",color)
	
	//方式三:使用类型推导，并且不使用关键字var声明并赋值
	name,age,color := "zhangsan",18,"red"
	fmt.Println("局部变量：\nname =",name,"age =",age,"color =",color)
	
	fmt.Println("全局变量：\nn5 =",n5,"n6 =",n6,"n7 =",n7)
}
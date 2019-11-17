package main

import "fmt"
//变量的使用 - 快速入门
func main(){

    var i int
    i = 10
    fmt.Println("i=",i)

    //golang变量使用方式
    //第一种：指定变量类型声明后若不赋值，使用默认值
    var a int
    fmt.Println("a=",a)

    //第二种方式 根据值自行判断变量类型
    var num =10.66
    fmt.Println("num=",num)

    //第三种方式：省略var 注意:=左侧的变量不应该是已经声明过的，否则会导致编译错误
    //下面的方式等价 var name string ; name = "pangluo"
    name:="pangluo"
    fmt.Println("name=",name)
}
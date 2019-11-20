package main

import (
	"fmt"
	_ "time"
)
//算术运算的课堂作业
func main () {
	
	//假如还有97天放假问：还有几个星期零几天
	var recessday int = 97
	fmt.Printf("recessday has %dweek%dday\n",recessday/7,recessday%7)

	//定义一个变量保存华氏温度，华氏温度转换摄氏温度的公式为：5/9*(华氏温度-100)，请求出华氏温度对应的摄氏温度
	var fahrenheit float64 =134.2
	fmt.Printf("%f对应的摄氏温度为%f",fahrenheit,5.0/9*(fahrenheit-100))

}

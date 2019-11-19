package main

import(
	"fmt"
)
//string的基本使用
func main(){
	//string的基本使用
	var address string = "我爱中国"
	fmt.Println(address)
	
	//字符串一旦赋值了就不能修改了，在Go中字符串是不可变的
	// var str = "hello"
	// str[0] = 'c' //cannot assign to str[0]
	
	//字符串的两种变现形式
	//1)双引号，会识别转义字符
	str1 := "pangluo\tluo"
	fmt.Println(str1)
	
	//使用反引号
	str2 :=`package main

	import(
		"fmt"
	)
	//string的基本使用
	func main(){
		//string的基本使用
		var address string = "我爱中国"
		fmt.Println(address)
		
		//字符串一旦赋值了就不能修改了，在Go中字符串是不可变的
		var str = "hello"
		str[0] = 'c' //cannot assign to str[0]
		
		//字符串的两种变现形式
		//1)双引号，会识别转义字符
		str1 := "pangluo/tluo"
		fmt.Println(str1)
		
		//使用反引号
	} `
	fmt.Println(str2)

	//字符串拼接
	var str3 = "hello " + "world "
	str3 += ",china"
	fmt.Println(str3)

	//当一行字符串太长，需要使用多行字符串，需要将+连接符号保留在上一行
	str5 := "hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+
	"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+"hey,guys"+
	" boy"
	fmt.Println(str5)

	//基本数据类型的默认值
	var a int
	var b float32
	var c float64
	var d bool
	var e string
	fmt.Printf("\na =%d,b =%f,c =%f,d =%v,e =%v",a,b,c,d,e)
} 
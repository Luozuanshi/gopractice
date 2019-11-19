package main
import(
	"fmt"
	_ "unsafe"
	 "strconv"
)
//基本数据类型和string类型的转换
func main(){

	//b,_ =strconv.ParseBool(str)
	//说明
	// 1.strconv.ParseBool(str)函数会返回两个值（value bool，err error）
	// 2.只想获取到value bool，不想获取err 所以使用_忽略
	var str string = "true"
	var b bool
	b,_ = strconv.ParseBool(str)
	fmt.Printf("b type %T b =%v",b,b)

	var str2 string = "123456789101112"
	var n1 int64
	var n2 int
	n1, _ = strconv.ParseInt(str2,10,32) //第三个参数32： 先把字符串转换成范围为32位的int64，再把类型为int64的值返回给n1 结果2147483647 
	n2 = int(n1)
	fmt.Printf("\nn1 type %T,n1=%d,n2 type %T,n2 =%d",n1,n1,n2,n2)

	var str3 string ="123456.123456"
	var f1 float64
	f1,_ = strconv.ParseFloat(str3,64)
	fmt.Printf("\nf1 type %T，f1 =%f",f1,f1)

}
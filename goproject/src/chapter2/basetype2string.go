package main
import(
	"fmt"
	_ "unsafe"
	 "strconv"
)
//基本数据类型和string类型的转换
func main(){
	//第一种方式 使用fmt包下的Sprintf  转换
	var num1 int = 100
	var float1 float64 =0.111136
	var boolx bool = true
	
	var str string
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type is %T,value's %q",str,str)

	str = fmt.Sprintf("%f",float1)
	fmt.Printf("\n str type is %T,value's %q",str,str)

	str = fmt.Sprintf("%v",boolx)
	fmt.Printf("\n str type is %T,value's %q",str,str)
	
	//第二种方式 使用strconv包下的 formatxxx
	
	var num2 int = 66
	var float2 float64 =66.222
	var boolx2 bool = false
	
	str2 := strconv.FormatInt(int64(num2),10)
	fmt.Printf("\n str2 type is %T,value's %q",str2,str2)
	//说明： 'f' 需要格式化的格式 ；10：表示保留的小数位数10位；表示这个小数是float64 
	str2 = strconv.FormatFloat(float2,'f',10,64)
	fmt.Printf("\n str2 type is %T,value's %q",str2,str2)
	str2 = strconv.FormatBool(boolx2)
	fmt.Printf("\n str2 type is %T,value's %q",str2,str2)

}
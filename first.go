//程序所属的包
package main
//导入依赖包
import "fmt"
//定义常量
const NAME  string= "zhouboxi"
//定义全局变量
var mainName string= "main name"
//一般类型声明
type boboInt int
//结构的声明
type Learn struct {

}
//接口声明
type Ilearn interface {

}
//函数第一
func learnMore(){
	fmt.Print("hello learnMore")
}
//main函数
func main() {
	fmt.Println("hello main")
	fmt.Println(NAME)
	fmt.Println(mainName)

}

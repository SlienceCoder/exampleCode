package main

/*4. 每个源文件都可以包含一个init函数，这个init函数自动被go运行框架调用。开发一个程序
  演示这个功能？*/

import (
	a "exampleCode/example4/test4"
	// b "exampleCode/example4/test_4"
	"fmt"
)

var c = "初始化"

func main() {
	fmt.Println(a.James, c)
}

func init() {
	c = "nidaye de "
	fmt.Println(c)
}

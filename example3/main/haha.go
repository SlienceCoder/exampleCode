package main

// 3. 包别名的应用，开发一个程序，使用包别名来访问包中的函数？
import (
	a "exampleCode/example3" // 如果是_则表示忽略
	"fmt"
)

func main() {

	fmt.Println(a.Add(2, 3))
}

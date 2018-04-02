package main

/*2. 一个程序包含两个包add和main，其中add包中有两个变量：Name和age。请问main
包中如何访问Name和age？*/
import (
	a "exampleCode/example2"
	"fmt"
)

func main() {
	fmt.Println(a.Name)
	fmt.Println(a.Age)
}

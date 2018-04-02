package main

import (
	a "exampleCode/example7/test7"
	"fmt"
)

// 8. 写一个程序用来打印值类型和引用类型变量到终端，并观察输出结果。

func main() {
	var c chan int = make(chan int, 1)
	c <- 6
	temp1 := <-c
	fmt.Println(temp1)

	go a.Add(c)
	temp := <-c
	fmt.Println(temp)

	var a1, b = 22, 33
	fmt.Println(a1, b)
	a.Sub(a1, b)
	fmt.Println(a1, b)

	a.ChangeNum(&a1, &b)
	fmt.Println(a1, b)

}

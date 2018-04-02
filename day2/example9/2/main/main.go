package main

import (
	"fmt"
)

/*
2. 打印出100-999中所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字
立方和等于该数本身。例如：153 是一个“水仙花数”，因为 153=1 的三次
方＋5 的三次方＋3 的三次方
*/
func main() {
	var a, b, c int
	for i := 100; i < 1000; i++ {
		a = i % 10      // 个位
		b = i / 10 % 10 // 十位
		c = i / 100     // 百位
		// fmt.Println(a, b, c)
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}
}

package main

import (
	"fmt"
)

// 1. 写一个程序，对于给定一个数字n，求出所有两两相加等于n的组合。
func main() {
	test1(5)
}

func test1(n int) {

	for i := 0; i <= n; i++ {
		fmt.Println(i, "+", n-i, "=", n)
	}

}

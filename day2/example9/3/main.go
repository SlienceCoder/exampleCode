package main

import (
	"fmt"
)

// 3. 对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n!

func main() {
	pingfang(3)
}

func pingfang(n int) {
	var sum int
	for i := 1; i <= n; i++ {
		var subSum int = 1
		for j := 1; j <= i; j++ {
			subSum *= j
			// fmt.Println(subSum)
		}
		sum += subSum
	}
	fmt.Println(sum)
}

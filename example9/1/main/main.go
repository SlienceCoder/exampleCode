package main

import (
	"fmt"
	"math"
)

// 1. 判断 101-200 之间有多少个素数，并输出所有素数。
func main() {
	var count int
	for i := 101; i < 201; i++ {
		for j := 2; j < math.Sqrt(i); j++ {
			if i%j == 0 {
				break
			}
			if k > math.Sqrt(i) {
				fmt.Println(i)
			}

		}
	}
	fmt.Println(count)
}

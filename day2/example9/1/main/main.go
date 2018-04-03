package main

import (
	"fmt"
	// "math"
)

// 1. 判断 101-200 之间有多少个素数，并输出所有素数。

func judgeSushu(n int) bool {
	var flag bool = true
	for i := 2; i < n; i++ {
		if n%i == 0 {
			flag = false
		}

	}

	if flag {
		fmt.Println("我是素数--", n)
		return true
	} else {
		return false
	}
}

func main() {

	var n, m int = 101, 200

	// Scanf("%d%d", &n, &m)

	for i := n; i < m; i++ {

		judgeSushu(i)

	}

	// fmt.Println(judgeSushu(10))

}

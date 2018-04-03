package main

import (
	"fmt"
)

func main() {
	// wanshua(1, 1000)
	var str string
	fmt.Scanf("%s", &str)
	huiwen(str)

}

/*输入一个字符串，判断其是否为回文。回文字符串是指从左到右读和从右到
左读完全相同的字符串*/
func huiwen(str string) bool {
	var tempStr string
	var strByte = []byte(str)
	for i := 0; i < len(str); i++ {
		tempStr += string(strByte[len(str)-i-1])
	}
	// fmt.Println(tempStr, str)

	if tempStr == str {
		fmt.Println("是回文数")
		return true
	} else {
		fmt.Println("不是回文数")
		return false
	}

}

/*一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
编程找出1000以内的所有完数*/
func wanshua(a, b int) {
	for i := a; i <= b; i++ {
		wanshu(i)
	}
}
func wanshu(n int) {
	var sum int = 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	if sum == n {
		fmt.Println(n)
	}
}

// 编写程序，在终端输出九九乘法表。
func niee() {
	for i := 1; i < 10; i++ {
		var str string
		for j := 1; j <= i; j++ {
			// fmt.Println(j, "*", i, "=", i*j)
			s := fmt.Sprintf("%d * %d = %d  ", j, i, i*j)
			str += s

		}
		fmt.Println(str)
	}
}

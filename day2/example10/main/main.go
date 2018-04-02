package main

// 字符串的相关操作以及格式化输出
import (
	"fmt"
)

func main() {
	var str1, str2 = "abc", "def"

	str3 := str1 + "---" + str2

	fmt.Println(str3)

	// slice(str3[-3:])
	fmt.Println(str3[0:5])
	// fmt.Println(str3.reverse())
	// 字符串的反转
	fmt.Println(reverse(str3))
	fmt.Println(string([]byte(str3)))
}

func reverse(str string) string {

	var n int = len(str)

	var tempStr string
	for i := 0; i < n; i++ {
		tempStr = tempStr + fmt.Sprintf("%c", str[n-1-i])
	}
	return tempStr

}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// strings和strconv的使用

func main() {

	var str = "aheHEBNh dshh sghhh_sdsd   k90"

	// 1.strings.HasPrefix(s string, prefix string) bool：判断字符串s是否以prefix开头
	a := strings.HasPrefix(str, "ah")
	fmt.Println("HasPrefix--", a)

	// 2. strings.HasSuffix(s string, suffix string) bool：判断字符串s是否以suffix结尾。
	a1 := strings.HasSuffix(str, "k")
	fmt.Println("HasSuffix--", a1)

	// 3. strings.Index(s string, str string) int：判断str在s中首次出现的位置，如果没有,返回-1
	a3 := strings.Index(str, "e")
	fmt.Println("Index--", a3)

	// 4. strings.LastIndex(s string, str string) int：判断str在s中最后出现的位置，如果没有,返回-1
	a4 := strings.LastIndex(str, "d")
	fmt.Println("LastIndex--", a4)

	// 5. strings.Replace(str string, old string, new string, n int)：字符串替换
	a5 := strings.Replace(str, "ds", "---", 2)
	fmt.Println("Replace--", a5)

	// 6. strings.Count(str string, substr string)int：字符串计数
	a6 := strings.Count(str, "ds")
	fmt.Println("Count--", a6)

	// 7. strings.Repeat(str string, count int)string：重复count次str
	a7 := strings.Repeat(str, 2)
	fmt.Println("Repeat--", a7)

	// 8. strings.ToLower(str string)string：转为小写
	a8 := strings.ToLower(str)
	fmt.Println("ToLower--", a8)

	// 9. strings.ToUpper(str string)string：转为大写
	a9 := strings.ToUpper(str)
	fmt.Println("ToUpper--", a9)

	// 10. strings.TrimSpace(str string)：去掉字符串首尾空白字符
	a10 := strings.TrimSpace(str)
	fmt.Println("TrimSpace--", a10)

	// 11.strings.Trim(str string, cut string)：去掉字符串首尾cut字符
	a11 := strings.Trim(str, "0")
	fmt.Println("Trim--", a11)

	// 12.strings.TrimLeft(str string, cut string)：去掉字符串首cut字符
	a12 := strings.TrimLeft(str, "a")
	fmt.Println("TrimLeft--", a12)

	// 13.strings.TrimLeft(str string, cut string)：去掉字符串首cut字符
	a13 := strings.TrimRight(str, "90")
	fmt.Println("TrimRight--", a13)

	// 14. strings.Field(str string)：返回str空格分隔的所有子串的slice
	a14 := strings.Fields(str)
	fmt.Println("Fields--", a14)

	// 15.strings.Split(str string, split string)：返回str split分隔的所有子串的slice
	a15 := strings.Split(str, "ds")
	fmt.Println("Split--", a15, len(a15))

	// 16.strings.Join(s1 []string, sep string)：用sep把s1中的所有元素链接起来
	a16 := strings.Join(["1","2","3"] , "4")
	fmt.Println("Join--",a16)

	// 17.strings.Itoa(i int)：把一个整数i转成字符串
	a17 := strconv.Itoa(344)
	fmt.Println("Itoa--", a17)

	// 18.strings.Atoi(str string)(int, error)：把一个字符串转成整数
	a18, err := strconv.Atoi("78")
	fmt.Println("Atoi--", a18, err)

	

}

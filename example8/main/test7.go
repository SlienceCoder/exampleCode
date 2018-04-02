package main

// 9. 写一个程序，交换两个整数的值。比如： a=3; b=4; 交换之后：a=4;b=3
import (
	"fmt"
)

func main() {
	var a, b = 33, 55
	fmt.Println(a, b)
	exchamge(&a, &b)
	fmt.Println(a, b)
}

func exchamge(a, b *int) {
	*a, *b = *b, *a
}

package main

import (
	"fmt"
	"time"
)

func main() {

	begin := time.Now().Nanosecond()

	fmt.Println(time.Now().Day())

	fmt.Println(time.Now().Minute())

	fmt.Println(time.Now().Month())

	fmt.Println(time.Now().Year())

	fmt.Println(time.Now().Hour())

	fmt.Println(time.Now().Second())

	now := time.Now()

	fmt.Printf("%04d/%02d/%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(now.Format("02/1/2006 15:02:09"))
	fmt.Println(now.Format("2006/03/02 15:02"))
	fmt.Println(now.Format("2006/04/02"))

	end := time.Now().Nanosecond()
	fmt.Println((end - begin) / 1000)

	// arr := []int{1, 2, 3, 4, 5}
	// fmt.Println(add(arr))
	f()
}

func f() {

	for i := 0; i < 5; i++ {
		defer fmt.Println("结束了--", i)
	}

}

func add(arg ...int) int {

	var sum int

	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum

}

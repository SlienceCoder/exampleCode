package main

import (
	"fmt"
	"os"
)

func main() {
	var goos = os.Getenv("GOOS")
	fmt.Println(goos)
	var path = os.Getenv("PATH")
	fmt.Println(path)
}

package test7

func Add(pipe chan int) {
	pipe <- 7
}

func Sub(a, b int) {
	a, b = b, a
}

func ChangeNum(a, b *int) {

	*a, *b = *b, *a

}

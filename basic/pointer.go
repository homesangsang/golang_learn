package main

import "fmt"

/**
交换2个变量的值
*/
func swap(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Print(a, b)
}

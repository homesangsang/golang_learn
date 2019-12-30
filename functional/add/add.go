package main

import "fmt"

func addr() func(int) int {
	num := 0
	return func(i int) int {
		num += i
		return num
	}
}

func main() {
	a := addr()
	for i := 0; i < 10; i++ {
		fmt.Println("1 + 2 + ... + ", i, a(i))
	}

	b := a(10)
	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", a, a)
}

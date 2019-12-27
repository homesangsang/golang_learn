package main

import (
	"fmt"
	"math"
)

/**
返回单个值
*/
func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("unsupported operation: " + op)
	}
}

/**
返回多个值
*/
func div(a, b int) (int, int) {
	return a / b, a % b
}

/**
返回值别名
*/
func div1(a, b int) (q, r int) {
	q = q / b
	r = a % b
	return q, r
}

/**
函数式
*/
func apply(op func(a, b int) int, a, b int) int {
	return op(a, b)
}

/**
可变参数
*/
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(2, 3, "+"))
	fmt.Println(div(9, 4))
	a, b := div(9, 4)
	fmt.Println(a, b)
	// 接受单个值的方式
	c, _ := div(9, 4)
	fmt.Println(c)

	// 函数式
	fmt.Println(
		apply(func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 3, 4),
	)

	// 可变参数
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))
}

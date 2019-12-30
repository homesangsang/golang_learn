package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func demo() {
	c := 3 + 4i
	r := cmplx.Abs(c) // 绝对值
	fmt.Println(r)
}

func euler() {
	//result := cmplx.Pow(math.E, 1i * math.Pi) + 1
	result := cmplx.Exp(1i*math.Pi) + 1 // cmplx.Exp 是e的指数
	fmt.Println(result)
}
func main() {
	euler()
}

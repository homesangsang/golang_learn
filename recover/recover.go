package main

import (
	"errors"
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {
	// 匿名函数
	defer func() {
		r := recover()
		if error, ok := r.(error); ok {
			fmt.Println("Error occurred: ", error)
		} else {
			panic(error)
		}
	}()
	panic(errors.New("this is an error"))
}

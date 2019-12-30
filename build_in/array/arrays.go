package main

import "fmt"

/**
 * 使用指针传递数组，可改变数组的值
 */
func PrintArr5UsingPointer(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	// 可以改变原数组的值
	arr[0] = 1000
}

/**
 * 数组是值类型，所以必须指明数量
 * 如果不指明，则arr [] int表示切面，和arr [5]int是完全不同的对象
 */
func PrintArr5(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

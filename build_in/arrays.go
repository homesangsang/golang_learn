package main

import "fmt"

/**
 * 数组是值类型，所以必须指明数量
 * 如果不指明，则arr [] int表示切面，和arr [5]int是完全不同的对象
 */
func printArr5(arr [5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

/**
 * 使用指针传递数组，可改变数组的值
 */
func printArr5UsingPointer(arr *[5]int) {
	for i, v := range arr {
		fmt.Println(i, v)
	}
	// 可以改变原数组的值
	arr[0] = 1000
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 3, 4, 5, 6}
	fmt.Println(arr1, arr2, arr3)

	printArr5(arr1)
	//printArr5(arr2) 报错，因为 [3]int和[5]int是2种数据类型
	printArr5(arr3)

	printArr5UsingPointer(&arr1)
	fmt.Println(arr1[0])
}

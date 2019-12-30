package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3}
	fmt.Println("array ", a)

	for x := range a {
		fmt.Println(a[x])
	}

	for key, value := range a {
		fmt.Println("key: ", key, ", value: ", value)
	}

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	var arr1 [5]int
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 3, 4, 5, 6}
	fmt.Println(arr1, arr2, arr3)

	PrintArr5(arr1)
	//printArr5(arr2) 报错，因为 [3]int和[5]int是2种数据类型
	PrintArr5(arr3)

	PrintArr5UsingPointer(&arr1)
	fmt.Println(arr1[0])
}

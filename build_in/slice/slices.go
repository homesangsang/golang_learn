package main

import "fmt"

/**
 * 函数参数为切片，切片会改变原参数的值
 */
func slicesDemo(arr []int) {
	fmt.Println("arr【2:6] = ", arr[2:6]) // [2, 3, 4, 5]
	fmt.Println("arr【:6] = ", arr[:6])   // [1, 2, 3, 4, 5]
	fmt.Println("arr【2:] = ", arr[2:])   // [2, 3, 4, 5, 6, 7]
	fmt.Println("arr【:] = ", arr[:])     // [0, 1, 2, 3, 4, 5, 6, 7]
	// 会改变原数组的值
	arr[0] = 1000
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	slicesDemo(arr[:])
	fmt.Println(arr)

	// slice改变原数组中的数据
	arr1 := [...]int{1, 2, 3, 4, 5}
	s := arr1[2:4]
	s[0] = 1000
	fmt.Println(arr1)

	// reslice: 在slice的基础上继续获取slice
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6}
	sarr21 := arr2[2:6] // 2, 3, 4, 5
	sarr22 := sarr21[3:5]
	fmt.Println(sarr21, sarr22)           // 5, 6
	fmt.Println(cap(sarr21), len(sarr21)) // 5, 4

	// 测试 slice扩容后对原数组的影响
	testSliceAndArray()
}

/**
 * 详细测试了基于数组创建sclice后，对slice的操作如何影响数组
 */
func testSliceAndArray() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	b := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a == b)
	s1 := a[2:6] // len = 4, cap = 6
	s2 := s1[3:6]
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	// 当slice的len <= cap时，slice 不会分配新的数组空间，此时对slice的append操作，可以对应修改到原数组
	s1 = append(s1, 10) // len = 5, cap = 6；数组不会扩容
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len=%d, cap=%d\n", s2, len(s2), cap(s2))
	fmt.Println(a)
	s1 = append(s1, 11) // len = 6, cap = 6; 数组不会扩容
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Println(a)
	s1 = append(s1, 12) // len = 7, cap = 12; 数组会扩容，12不会添加到原数组a中
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Println(a)
	s1[0] = 13 // 此时s1已经指向了新的数组，对s1的修改不会再修改a的值
	fmt.Printf("s1: %v, len=%d, cap=%d\n", s1, len(s1), cap(s1))
	fmt.Println(a)
	s1[0] = 13

}

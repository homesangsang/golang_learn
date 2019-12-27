package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func createSlice() {
	var a []int // a 此时是nil, 但是a可以直接使用
	printSlice(a)
	b := []int{} // b 此时是[]
	printSlice(b)
	c := []int{1, 2, 3, 4}
	printSlice(c)
	d := c[2:3]
	printSlice(d)
	e := make([]int, 16) // 创建一个len=16, cap=16的slice
	printSlice(e)
	f := make([]int, 10, 32) // 创建一个len=10, cap=32的slice
	printSlice(f)
}

/**
 * 复制切片内容
 */
func copyingSLice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 10)
	copy(s2, s1) // s2 = [1, 2, 3, 4, 5, 0, 0, 0, 0, 0]
	printSlice(s2)
}

/**
 * 删除切片内容
 * 对于删除操作，由于len一直小于cap，cap长度不变，所以对原切片数据的修改，不会导致内存空间结构的变化
 */
func deletingSlice() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("从开头删除")
	s1 = s1[1:]
	printSlice(s1) // s1 = [1, 2, 3, 4, 5, 6, ,7 ,8, 9]
	fmt.Println("从中间位置删除")
	// 删除索引为3的元素, ...表示获取索引4后面的所有值
	s1 = append(s1[:3], s1[4:]...) // s1 = [1, 2, 3, 5, 6, 7, 8, 9]
	printSlice(s1)
	// copy函数返回被复制元素的个数
	s1 = s1[:3+copy(s1[3:], s1[4:])] // s1 = [1, 2, 3, 6, 7, 8, 9]
	printSlice(s1)
	fmt.Println("从尾部删除")
	s1 = s1[:len(s1)-1] // 删除尾部1个元素 s1 = [1, 2, 3, 6, 7, 8]
	printSlice(s1)
	s1 = s1[:len(s1)-3] // 删除尾部3个元素 s1 = [1, 2, 3]
	printSlice(s1)
}

func nilSliceDemo() {
	var a []int
	// i = 0时，不会因为a == nil 而报错
	// zero value for slice is nil
	for i := 0; i < 100; i++ {
		a = append(a, i*2+1)
	}
	fmt.Println(a)
}

func main() {
	//createSlice()
	//copyingSLice()
	deletingSlice()
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "golang学习"
	// 直接打印数组底层的存储结构，得到二进制数组
	for i, elem := range []byte(str) {
		fmt.Printf("(%d %X) ", i, elem)
	}
	fmt.Println()

	// 直接获取rune，观察索引i
	for i, ch := range str {
		fmt.Printf("(%d, %c) ", i, ch)
	}
	fmt.Println()

	// 通过utf8.DecodeRune，将字符串按照utf8的格式解析，打印出单个字符
	bytes := []byte(str)
	for len(bytes) > 0 {
		ru, size := utf8.DecodeRune(bytes)
		fmt.Printf("(%c %d) ", ru, size)
		bytes = bytes[size:]
	}
	fmt.Println()

	// 将字符串强制转化为字符切片, 通过结果可以发现与转成byte数组有了结构的变化
	fmt.Println(str)
	for _, ch := range []rune(str) {
		fmt.Printf("(%c %X) ", ch, ch)
	}
	fmt.Println()

}

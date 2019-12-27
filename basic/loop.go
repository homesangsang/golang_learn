package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// loop 循环

/**
将整数转换成二进制的表达式
*/
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic("err")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func main() {
	//fmt.Println(10, convertToBin(10))
	//fmt.Printf("%b", 10)
	printFile("abc.txt")
}

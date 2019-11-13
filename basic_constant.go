package main

import "fmt"

func enums()  {
	const (
		cpp = iota
		_
		java
		golang
		python
	)
	// iota 参与运算; iota 可以作为自增值的种子
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, pb)
}

func constant() {
	const filename = "xxx.txt"
	const a, b = 3, 4
	fmt.Printf("%T, %T \n", a, b)
	fmt.Printf("%T", filename)
}
func main() {
	//constant()
	enums()
}

package main

import "fmt"

var (
	a, b, c = 3, true, "fff"
)

func variableZeroValue()  {
	var a int
	var b string
	// \q可以打印出空字符串
	fmt.Printf("%d, %q\n", a, b) // 0， ""
}

func variableInitialValue()  {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b)
	fmt.Println(s)
}

func variableTypeDeduction()  {
	var a, b, c = 1, true, "abc"
	fmt.Println(a, b, c)
}

func variableShorter() {
	a, b, c := 2, false, "cde"
	fmt.Println(a, b, c)
}

func main() {
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeduction()
	variableShorter()
	fmt.Println(a, b, c)
}

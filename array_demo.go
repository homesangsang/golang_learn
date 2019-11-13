package main

import "fmt"

func main() {
	var a = [...]int {1, 2, 3}
	fmt.Println("array ", a)

	for x := range a {
		fmt.Println(a[x])
	}

	for key, value := range a {
		fmt.Println("key: ", key, ", value: " , value)
	}

	for i := 0; i < len(a); i++  {
		fmt.Println(a[i])
	}

}

package fun

import (
	"bufio"
	"fmt"
	"golang_learn/functional/fb"
	"os"
)

func TryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func TryDefer1() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	panic("hello world")
	defer fmt.Println(4)
}

func WriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	// 后执行的defer
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 先执行的defer
	defer writer.Flush()

	f := fb.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func WriteFileErr(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
	}
	// 后执行的defer
	defer file.Close()

	writer := bufio.NewWriter(file)
	// 先执行的defer
	defer writer.Flush()

	f := fb.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

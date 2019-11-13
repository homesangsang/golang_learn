package main

import (
	"fmt"
	"io/ioutil"
)

// branch 分支, 条件控制

func branch1(filename string) {
	content, err := ioutil.ReadFile(filename) // 读取文件的内容
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func branch2(filename string) {
	if content, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", content)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("异常")
	case score < 60:
		g = "不及格"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	//branch1(filename)
	branch2(filename)
}

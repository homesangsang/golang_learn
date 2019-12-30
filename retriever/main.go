package main

import (
	"fmt"
	"golang_learn/retriever/mock"
	customReal "golang_learn/retriever/real"
)

type Retriever interface {
	Get(url string) string
	Get1(url string, param string) string
}

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func download1(r Retriever) string {
	return r.Get1("https://www.baidu.com/s", "wd=hello")
}

// 获取接口内部的类型和值
func inspect(r Retriever) {
	fmt.Printf("%T %v\n", r, r)
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case customReal.Retriever:
		fmt.Println("TimeOut: ", v.TimeOut)
	}
}

func main() {
	var r Retriever
	r = mock.Retriever{"www.baidu.com"}
	inspect(r)
	r = customReal.Retriever{}
	inspect(r)

	//fmt.Println(download1(r))
}

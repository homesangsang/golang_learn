package main

import "golang_learn/tree"

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}              // 因为left是指针类型，所以必须取地址
	root.Right = &tree.Node{5, nil, nil}  // 初始化方式
	root.Right.Left = new(tree.Node)      // 初始化方式
	root.Right.Right = tree.CreateNode(2) // 工厂函数方式

	root.Right.Left.SetValue(4)
	root.Right.Left.Print()

	// 结构体的函数
	root.Traverse()
}

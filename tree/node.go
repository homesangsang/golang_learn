package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

/**
工厂函数
*/
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

/**
 * 结构体的函数
 * 接收者（node treeNode）
 * 调用时，使用 node.print()
 */
func (node treeNode) print() {
	fmt.Println(node.value)
}

/**
 * 必须加*号，如果不加，则是值传递，无法更改引用的值
 */
func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node treeNode) getValue() int {
	return node.value
}

/**
 * 普通函数， 调用时使用 print(node)
 * 与上一个函数没有本质区别
 */
func print(node treeNode) {
	fmt.Println(node.value)
}

/**
 * 中序遍历
 */
func (node *treeNode) traverse() {
	if nil == node {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}             // 因为left是指针类型，所以必须取地址
	root.right = &treeNode{5, nil, nil} // 初始化方式
	root.right.left = new(treeNode)     // 初始化方式
	root.right.right = createNode(2)    // 工厂函数方式

	root.right.left.setValue(4)
	root.right.left.print()

	// 结构体的函数
	root.print()
	root.traverse()
}

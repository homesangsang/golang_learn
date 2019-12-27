package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

/**
工厂函数
*/
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

/**
 * 结构体的函数
 * 接收者（node treeNode）
 * 调用时，使用 node.print()
 */
func (node Node) Print() {
	fmt.Println(node.Value)
}

/**
 * 必须加*号，如果不加，则是值传递，无法更改引用的值
 */
func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node Node) GetValue() int {
	return node.Value
}

/**
 * 普通函数， 调用时使用 print(node)
 * 与上一个函数没有本质区别
 */
func print(node Node) {
	fmt.Println(node.Value)
}

/**
 * 中序遍历
 */
func (node *Node) Traverse() {
	if nil == node {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

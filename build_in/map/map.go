package main

import (
	"fmt"
)

/**
 * map的创建方式
 */
func mapDefine() {
	m := map[string]string{
		"name": "Join",
		"age":  "38",
	}
	fmt.Println(m)
	m1 := make(map[int]int) // m1 == empty map -> false
	fmt.Println(m1, m1 == nil)
	var m2 = map[int]int{}
	fmt.Println(m2, m2 == nil) // m2 == empty map -> false
	var m3 map[int]int
	fmt.Println(m3, m3 == nil) // m3 == nil -> true, 不影响计算
}

/**
 * map遍历
 */
func mapTraversing(m map[string]string) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

/**
 * 获取值，同时获取key是否存在
 */
func mapGetValue(m map[string]string) {
	// exist变量可以获取到当前的key是否存在
	name, exist := m["name"]
	fmt.Println(name, exist)
	number, exist := m["number"]
	fmt.Println(number, exist) // 如果key不存在，会默认返回一个空的变量
}

/**
 * 删除值
 */
func mapDeletingValue(m map[string]string) {
	delete(m, "name")
	fmt.Println(m)
}

func main() {
	m := map[string]string{
		"name": "Join",
		"age":  "38",
	}
	//mapDefine(&m) // map[age:38 name:Join]
	//mapTraversing(&m)
	//mapGetValue(&m)
	mapDeletingValue(m)
}

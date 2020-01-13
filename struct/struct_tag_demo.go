package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StructDemo struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func structDemo() {
	demo := StructDemo{Id: "202001011010", FirstName: "Smith", LastName: "John"}
	j, _ := json.Marshal(demo)
	fmt.Println(string(j))
}

func anotherDemo() {
	type User struct {
		UserId   int    `json:"userId" bson:"user_id"`
		UserName string `json:"userName" bson:"user_name"`
	}
	// 输出json格式
	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
	// 输出内容：{"user_id":1,"user_name":"tony"}

	// 获取tag中的内容
	t := reflect.TypeOf(u)
	field := t.Elem().Field(0)
	fmt.Println(field.Tag.Get("json"))
	// 输出：user_id
	fmt.Println(field.Tag.Get("bson"))
	// 输出：user_id
}

func main() {
	structDemo()
	fmt.Println()
	anotherDemo()
}

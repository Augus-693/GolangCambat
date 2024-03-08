package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/7 13:30
 * @GoVersion go1.22.0
 * @Version 1.0
 */

type Student struct {
	Id      int    `json:"_"`                 //在json编码时 这个字段不参与
	Name    string `json:"student_name"`      //在json编码时 这个字段会被重命名
	Age     int    `json:"Age,string"`        //在json编码时 将Age转换成String类型
	Address string `json:"Address,omitempty"` //在json编码时 如果这个字段是空的 那么忽略 不参与编码
	gender  string //由于gender在结构体中是小写的 json编码时忽略
}

func main() {
	augus := Student{
		Id:      001,
		Name:    "Augus",
		Age:     23,
		Address: "",
		gender:  "男",
	}

	//序列化（编码）
	encodeInfo, err := json.Marshal(&augus)
	if err != nil {
		fmt.Println("json.Marshal failed:", err)
		return
	}
	fmt.Println("encodeInfo:", string(encodeInfo))

	//反序列化（解码）
	var augus2 Student
	if err := json.Unmarshal([]byte(encodeInfo), &augus2); err != nil {
		fmt.Println("json.Unmarshal failed:", err)
		return
	}
	fmt.Println("id: ", augus2.Id)
	fmt.Println("name: ", augus2.Name)
	fmt.Println("age: ", augus2.Age)
	fmt.Println("address: ", augus2.Address)
	fmt.Println("gender: ", augus2.gender) //由于gender在结构体中是小写的 json编码时忽略

}

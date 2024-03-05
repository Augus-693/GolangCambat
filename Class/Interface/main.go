package main

import (
	"fmt"
	"reflect"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 11:01
 * @GoVersion go1.22.0
 * @Version 1.0
 */

// interface不仅仅用于处理多态 它可以接收任意数据类型
func main() {
	var i, j, k interface{}
	names := []string{"Augus", "Cupid"}
	i = names
	fmt.Println("the data of i is:", i, ",the type of i is:", reflect.TypeOf(i))

	age := 23
	j = age
	fmt.Println("the data of j is:", j, ",the type of j is:", reflect.TypeOf(j))

	str := "hello"
	k = str
	fmt.Println("the data of k is:", k, ",the type of k is:", reflect.TypeOf(k))

	//创建一个具有三个接口的切片
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "Hello World"
	array[2] = true

	for _, value := range array {
		//获取当前接口的真正数据类型
		switch v := value.(type) {
		case int:
			fmt.Printf("当前数据类型为int，内容为：%d\n", v)
		case string:
			fmt.Printf("当前数据类型为string，内容为：%s\n", v)
		case bool:
			fmt.Printf("当前数据类型为bool，内容为：%v\n", v) //%v可以自动推导数据类型
		default:
			fmt.Println("其他数据类型")
		}
	}

}

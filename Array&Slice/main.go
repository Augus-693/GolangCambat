package main

import (
	"fmt"
)

/**
 * @Project GolangCambat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2023/8/12 16:31
 * @GoVersion go1.20.5
 * @Version 1.0
 */

func array() {
	fmt.Println("定长数组")
	//1- 定义 定长数组
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//2- 遍历，方式1
	for i := 0; i < len(nums); i++ {
		fmt.Println("i: ", i, ", j: ", nums[i])
	}

	//2- 遍历，方式2：for range
	//key是数组下标，val是数组下标对应的值
	for key, val := range nums {
		//val全程是一个临时变量，不断地被重复赋值，修改它不会修改原始数组
		fmt.Println("key: ", key, ", val: ", val)
	}

	//在go语言中，忽略一个值使用_
	for _, value := range nums {
		fmt.Println("value: ", value)
	}
}

func slice() {
	fmt.Println("切片：不定长数组")
	//切片：slice，它的底层也是数组，可以动态改变长度
	addresses := []string{"北京", "上海", "广州", "深圳"}
	fmt.Printf("追加元素前——The length of addresses:%d, the capacity of addresses:%d\n", len(addresses), cap(addresses))
	//cap是切片容量，容量不够时，切片容量X2倍增加
	for i, v := range addresses {
		fmt.Println("i:", i, " v:", v)
	}

	//追加数据
	//append：切片数据增加的函数
	addresses = append(addresses, "重庆")
	fmt.Println("addresses追加元素赋值给自己:", addresses)
	fmt.Printf("追加元素后——The length of addresses:%d, the capacity of addresses:%d\n", len(addresses), cap(addresses))

	//切片截取
	nums := []int{0, 1, 2, 3, 4, 5}
	nums1 := nums[0:3] //左闭右开
	fmt.Println("nums1:", nums1)

	//使用make进行创建切片
	//nums := make([]int, 6, 8)
	////语法：func make([]T, len, cap) []T
	//fmt.Printf("nums = %v, \nlength = %d, \ncapacity = %d\n",
	//	nums, len(nums), cap(nums))
	//
	//for i := 0; i < len(nums); i++ {
	//	nums[i] = i
	//
	//}
	//fmt.Println("nums: ", nums)

}
func main() {
	array() //定长数组
	slice() //切片：不定长数组
}

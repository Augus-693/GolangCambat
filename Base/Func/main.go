package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/2 11:16
 * @GoVersion go1.22.0
 * @Version 1.0
 */

// 1.函数返回值在参数列表之后
// 2。函数返回值可以有多个，需要用圆括号，多个参数之间使用逗号分割
func test1(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

func test2(a, b int, c string) (res int, str string, bl bool) {
	var i, j, k int
	i = 1
	j = 2
	k = 3
	fmt.Println(i, j, k)

	//直接使用返回值的变量名字进行计算
	res = a + b
	str = c
	bl = true

	//当返回值有名字时，可以直接简写return
	return
}

// 如果返回值只有一个参数 并且没有名字 你们不需要加圆括号
func test3() int {
	return 10
}

func memoryEscape() *string {
	city := "Beijing"
	ptr := &city
	return ptr
}
func main() {
	v1, s1, _ := test1(10, 20, "hello")
	fmt.Println("v1:", v1, "s1:", s1)

	v2, s2, _ := test2(10, 20, "hello")
	fmt.Println("v2:", v2, "s2:", s2)

	v3 := test3()
	fmt.Println("v3:", v3)

	//演示内存逃逸
	p1 := memoryEscape()
	fmt.Println("*p1 = ", *p1)
}

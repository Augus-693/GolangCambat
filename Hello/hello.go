package main

//每个go程序都必须有一个 package main
//每个go程序，都是 .go 结尾的
//go 程序中没有 .h，没有 .o，只有 .go
//一个package，包名，相当于命名空间
//std:cout

import "fmt"

//这是导入一个标准包 fmt，format，一般用于格式化输出、

/**
 * @Project GolangCambat
 * @File    hello.go.go
 * @Author  Augus Lee
 * @Date    2023/7/8 7:58
 * @GoVersion go1.20.5
 * @Version 1.0
 */

func main() {
	//主函数，所有的函数必须使用 func 开头
	//一个函数的返回值，不会放在 func 前，而是放在参数后面16
	//函数左花括号必须与函数名同行，不能写到下一行
	fmt.Println("Hello, World!")
}

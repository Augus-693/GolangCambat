package main

import (
	"fmt"
	"os"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/4 19:00
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func readFile(filename string) {
	f1, err := os.Open("filename")
	//defer f1.Close()
	defer func() {
		_ = f1.Close()
		fmt.Println("file is closed")
	}() //创建一个匿名函数 同时调用
	if err != nil {
		fmt.Println("Error opening file, ", err)
		return
	}

	buf := make([]byte, 1024)
	n, _ := f1.Read(buf)
	fmt.Println("读取文件实际长度：", n)
	fmt.Println("读取文件内容:", string(buf))
}

func main() {
	/*
		1.延迟：关键字 可以用于修饰语句、函数 确保语句可以在当前栈退出时执行
		2.一般用于做资源清理工作
		3.解锁、关闭文件
		4.在同一个函数中多次调用defer 执行时类似于栈的机制 先入后出
	*/

	filename := "test.txt"
	readFile(filename)
}

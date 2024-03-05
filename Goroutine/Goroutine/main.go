package main

import (
	"fmt"
	"time"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 13:46
 * @GoVersion go1.22.0
 * @Version 1.0
 */

// 子Go程
func display(num int) {
	count := 1
	for {
		fmt.Println("------Starting child goroutine...", count, "num is :", num)
		count++
		time.Sleep(time.Second * 1)
	}
}

func main() {
	//启动子Go程
	//go display() //使用函数

	/*
		go func() { //使用匿名函数
			count := 1
			for {
				fmt.Println("------Starting child goroutine...", count)
				count++
				time.Sleep(time.Second * 1)
			}
		}()
	*/

	for i := 0; i < 3; i++ {
		go display(i)
	}

	//主Go程
	count := 1
	for {
		fmt.Println("Starting main goroutine...", count)
		count++
		time.Sleep(time.Second * 1)
	}
}

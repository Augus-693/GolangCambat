package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 14:59
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	numChan := make(chan int, 10)
	go func() {
		for i := 0; i < 50; i++ {
			numChan <- i
			fmt.Println("写入数据：", i)
		}
		fmt.Println("写入完毕 管道即将关闭")
		close(numChan)
	}()

	//遍历管道时 只返回一个值
	for v := range numChan {
		fmt.Println("读取数据：", v)
	}
}

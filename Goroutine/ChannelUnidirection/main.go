package main

import (
	"fmt"
	"time"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 16:03
 * @GoVersion go1.22.0
 * @Version 1.0
 */

// producer 生产者 提供一个只写的通道
func producer(out chan<- int) {
	for i := 0; i < 50; i++ {
		out <- i
		fmt.Println("====>向管道中写入数据：", i)
	}

}

// console 消费者 提供一个只读的通道
func console(in <-chan int) {
	for v := range in {
		fmt.Println("从管道中读取数据：", v)
	}
}

func main() {
	//生产者消费者模型
	//在主函数中创建一个双向管道
	numChan := make(chan int, 10)

	//将numChan传给producer 负责生产
	go producer(numChan)

	//将numChan传给consumer 负责消费
	go console(numChan)

	time.Sleep(time.Second + 3)
	fmt.Println("the program is over")

}

package main

import (
	"fmt"
	"time"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 14:12
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	numChan := make(chan int, 10)

	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("子Go程1 读取数据  ====》 data：", data)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			numChan <- i
			fmt.Println("子Go程2 写入数据 ", i)
		}
	}()

	for i := 20; i < 50; i++ {
		numChan <- i
		fmt.Println("===》主Go程 写入数据 ", i)
	}

	time.Sleep(time.Second * 3)

}

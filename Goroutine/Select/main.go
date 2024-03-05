package main

import (
	"fmt"
	"time"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 16:16
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		for {
			fmt.Println("监听中......")
			select {
			case data1 := <-chan1:
				fmt.Println("从chan1抓取数据成功 data1：", data1)
			case data2 := <-chan2:
				fmt.Println("=======》从chan1抓取数据成功 data2：", data2)
			}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chan1 <- i
			time.Sleep(time.Second * 1 / 2)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			chan2 <- i
			time.Sleep(time.Second * 1)
		}
	}()

	for {
		fmt.Println("Over")
		time.Sleep(time.Second * 3)
	}
}

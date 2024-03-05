package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 14:04
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	go func() {
		func() {
			fmt.Println("this is the function in goroutine")
			//return           //返回当前函数
			//os.Exit(-1)      //退出进程
			runtime.Goexit() //退出当前goroutine
		}()
		fmt.Println("goroutine is exited")
	}()

	fmt.Println("this is the main function")
	time.Sleep(time.Second * 3)
	fmt.Println("the program is over")
}

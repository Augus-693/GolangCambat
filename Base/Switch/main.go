package main

import (
	"fmt"
	"os"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/2 15:48
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	//Go：os.Args ==>直接可以获取命令输入 是一个字符串切片
	cmds := os.Args
	for key, cmd := range cmds {
		fmt.Println("key:", key, "cmd:", cmd, "cmds length:", len(cmd))
	}

	if len(cmds) < 2 {
		fmt.Println("请正确输入参数：")
		return
	}

	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		//Go的switch中 默认加上了break，不需要手动处理
		//如果想要向下穿透 需要加上关键字 fallthrough
		fallthrough
	case "world":
		fmt.Println("world")
	default:
		fmt.Println("default called")

	}
}

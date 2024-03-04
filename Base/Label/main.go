package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/2 16:03
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	//标签  LABEL
	//goto LABEL
	//break LABEL
	//continue LABEL
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL1 //下次进入循环时 i不会保存之前的状态 重新从0开始计算
				//continue LABEL1 //continue会调到指定的位置 但是会记录之前的状态
				break LABEL1 //直接跳出指定位置的循环
			}
			fmt.Println("i", "j", i, j)
		}
	}
	fmt.Println("over")
}

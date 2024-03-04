package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/2 16:22
 * @GoVersion go1.22.0
 * @Version 1.0
 */

const (
	/*
		1.iota是常量组计数器
		2.iota从0开始 每换行递增1
		3.iota如果不赋值 默认与上一行表达式相同
		4.iota如果一行出现两个 值相同
		5.每个常量组的iota是独立的 如果遇到const iota清零
	*/

	MONDAY = iota + 1
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func main() {

	fmt.Println(MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY)
}

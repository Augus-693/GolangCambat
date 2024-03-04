package sub

import "fmt"

/**
 * @Project GolangCombat
 * @File    sub.go
 * @Author  Augus Lee
 * @Date    2024/3/2 13:49
 * @GoVersion go1.22.0
 * @Version 1.0
 */

// init函数没有参数 没有返回值
// 一个包中包含多个init函数时 调用顺序不确定
// init函数不允许用户显式调用
// 如果引用一个包只想使用init函数（MySQL的init对驱动进行初始化）
// 使用 import   _ "xxx/xxx/xxx"
func init() {
	fmt.Println("this is the first init() function in package sub")
}

func init() {
	fmt.Println("this is the second init() function in package sub")
}

func Sub(a, b int) int {
	return a - b
}

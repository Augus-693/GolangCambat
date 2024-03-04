package add

import "fmt"

/**
 * @Project GolangCombat
 * @File    add.go
 * @Author  Augus Lee
 * @Date    2024/3/2 13:48
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func init() {
	fmt.Println("this is the first init() function in package add")
}

func init() {
	fmt.Println("this is the second init() function in package add")
}

func Add(a, b int) int {
	return a + b
}

package main

import (
	"GolangCombat/Base/Init/add"
	"GolangCombat/Base/Init/sub"
	"fmt"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/4 18:43
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	resAdd := add.Add(15, 10)
	resSub := sub.Sub(15, 10)
	fmt.Println("resAdd:", resAdd, "resSub:", resSub)

}

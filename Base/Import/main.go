package main

import (
	"GolangCombat/Base/Import/add"
	"GolangCombat/Base/Import/sub"
	"fmt"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/2 13:47
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	resAdd := add.Add(15, 10)
	resSub := sub.Sub(15, 10)
	fmt.Println("resAdd:", resAdd, "resSub:", resSub)

}

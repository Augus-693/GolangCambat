package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/2 16:34
 * @GoVersion go1.22.0
 * @Version 1.0
 */

type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	augus := Student{
		name:   "Augus",
		age:    22,
		gender: "male",
		score:  100,
	}
	fmt.Println("augus:", augus.name, augus.age, augus.gender, augus.score)
	s1 := &augus
	fmt.Println("s1使用指针s1.打印:", s1.name, s1.age, s1.gender, s1.score)
	fmt.Println("s1使用指针(*s1).打印:", (*s1).name, (*s1).age, (*s1).gender, (*s1).score)
}

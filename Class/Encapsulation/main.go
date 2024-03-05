package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 9:56
 * @GoVersion go1.22.0
 * @Version 1.0
 */

type Person struct {
	//成员属性
	name   string
	age    int
	gender string
	score  float64
}

// 在类外面绑定方法
func (this *Person) Eat() {
	fmt.Println("Person is eating")
	//类的方法 可以使用自己的成员
	fmt.Println(this.name + " is eating")
}

func main() {
	augus := Person{
		name:   "Augus",
		age:    23,
		gender: "male",
		score:  99,
	}
	fmt.Println("augus:", augus)
	augus.Eat()
}

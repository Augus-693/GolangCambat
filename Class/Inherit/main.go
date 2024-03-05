package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 10:16
 * @GoVersion go1.22.0
 * @Version 1.0
 */

type Human struct {
	//成员属性
	name   string
	age    int
	gender string
}

// 定义一个学生类 嵌套 Human
type Student struct {
	hum    Human
	school string
	score  float64
}

// 定义一个老师类 去继承 Human
type Teacher struct {
	Human
	subject string
}

func (tihs *Human) Eat() {
	fmt.Println(tihs.name + " is eating")
}

func main() {
	s1 := Student{
		hum: Human{
			name:   "Augus",
			age:    22,
			gender: "male",
		},
		school: "Nyist",
		score:  99,
	}

	fmt.Println("name of s1 is", s1.hum.name)
	fmt.Println("school of s1 is", s1.school)

	t1 := Teacher{}
	t1.name = "Cupid"
	t1.age = 35
	t1.gender = "female"
	t1.subject = "English"
	fmt.Println("the information of t1 is : ", t1)
	t1.Eat()
}

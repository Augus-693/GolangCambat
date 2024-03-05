package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 13:17
 * @GoVersion go1.22.0
 * @Version 1.0
 */

// 实现Go多态 需要事先定义接口

// 定义一个接口
type IAttack interface {
	//接口函数可以有多个 但只能是函数原型 不能有实现
	Attack()
}

// 低等级
type HumLowLevel struct {
	name  string
	level int
}

// 高等级
type HumHighLevel struct {
	name  string
	level int
}

func (a *HumLowLevel) Attack() {
	fmt.Println("name is", a.name, ",level is", a.level)
}

func (a *HumHighLevel) Attack() {
	fmt.Println("name is", a.name, ",level is", a.level)
}

func DoAttack(a IAttack) {
	a.Attack()
}

func main() {
	var player IAttack //定义一个包含IAttack的接口变量

	lowLevel := HumLowLevel{
		name:  "David",
		level: 1,
	}
	highLevel := HumHighLevel{
		name:  "Mike",
		level: 10,
	}
	lowLevel.Attack()
	highLevel.Attack()

	player = &lowLevel
	player.Attack()
	player = &highLevel
	player.Attack()

	fmt.Println("多态 传入不同参数 调用不同接口")
	DoAttack(&lowLevel)
	DoAttack(&highLevel)
}

package main

import "fmt"

/**
 * @Project GolangCambat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    7/19/2023 4:42 PM
 * @GoVersion go1.20.5
 * @Version 1.0
 */

func variableDefinition() {
	fmt.Println("演示变量定义")
	//变量定义 var

	//先定义再赋值
	var name string //方式1——先定义再赋值
	name = "augus"

	//定义同时赋值
	var age int = 23

	//定义直接赋值，使用类型推导
	var sex = "male"

	//定义直接赋值，不使用 var
	address := "河南"

	//平行赋值
	email1, email2 := "auguslee529@gmail.com", "auguslee693@gmail.com"

	fmt.Printf("name: %s, age: %d, sex: %s, address: %s, email1: %s, email2: %s\n", name, age, sex, address, email1, email2)
}

func autoIncrement() {
	fmt.Println("演示自增变量")
	i, j := 20, 30
	i++ //必须单独一行
	fmt.Println("i: ", i)
	j--
	fmt.Println("i: ", i)
}

func pointer() {
	fmt.Println("演示指针")
	//go语言在使用指针时，会使用内部的垃圾回收机制（gc:garbage collector），开发人员不需要手动释放内存
	//c语言不允许返回栈上的指针，go语言可以返回栈上的指针，程序会在编译的时候确定变量的分配位置
	//编译的时候，如果发现有必要的话，就将变量分配到堆上
	name := "Cupit"
	ptr := &name
	fmt.Println("name:", name)
	fmt.Println("name ptr:", ptr)

	//使用new关键字定义
	name2Ptr := new(string)
	*name2Ptr = "Jack"
	fmt.Println("name2Ptr:", *name2Ptr)
	fmt.Println("name2Ptr ptr:", name2Ptr)

	//可以返回栈上的指针，编译器在编译程序时，会判断这段代码，将city变量分配在堆上
	res := testPtr()
	fmt.Println("res city :", *res, ", address : ", res)
}

// 定义一个函数，返回一个string类型的指针，go语言返回写在参数列表后面
func testPtr() *string {
	city := "beijing"
	ptr := &city
	return ptr
}

func stringType() {
	fmt.Println("演示String")
	//1- 定义
	name := "Augus"

	//需要换行，原生输出字符串时，使用反引号``
	usage := `./a.out <option>
		-h helo
		-a xxx`

	fmt.Println("name : ", name)
	fmt.Println("usage : ", usage)

	//2- 长度，访问
	l1 := len(name)
	fmt.Println("length of name : ", l1)

	//字符串遍历
	for i := 0; i < len(name); i++ {
		fmt.Printf("i: %d, v: %c\n", i, name[i])
	}

	//字符串拼接
	i, j := "hello", "world"
	fmt.Println("i + j: ", i+j)
}

func main() {
	variableDefinition() //变量定义
	autoIncrement()      //自增变量
	pointer()            //指针
	stringType()         //字符串
}

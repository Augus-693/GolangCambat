package main

import "fmt"

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2023/12/19 21:16
 * @GoVersion go1.21.3
 * @Version 1.0
 */

func main() {
	//定义一个字典：id——>name idNames
	var idNames map[int]string //定义一个map

	//分配空间，使用make，可以不指定长度，建议直接指定长度，性能更好
	idNames = make(map[int]string, 10)

	//定义时直接分配空间
	//idNames := make(map[int]string, 10)  //建议使用这种方式

	idNames[0] = "Augus"
	idNames[1] = "Cupid"

	//遍历Map
	for key, value := range idNames {
		fmt.Println("key:", key, ",value:", value)
	}

	//如何确定一个kay是否存在于Map中
	//在 Map中不存在访问越界的问题，它认为所有的key都是有效的，所以访问一个不存在的key不会崩溃，返回这个类型的零值
	//零值： bool=false,int=0,float=0,complex=0,string=""
	name9 := idNames[9]
	fmt.Println("name9:", name9) //输出 name9:

	//无法通过获取value值来判断一个key是否存在，因此我们需要一个能够校验key是否存在的方式
	value, ok := idNames[1] //如果id=1是存在的，那么value就是key=1对应值，ok返回true，反之返回零值，false
	if ok {
		fmt.Println("key=1这个key是存在的，value为：", value)
	}

	value, ok = idNames[10]
	if ok {
		fmt.Println("id=10这个key是存在的，value为：", value)
	} else {
		fmt.Println("id=10这个key不存在，value为：", value)
	}

	//删除Map中元素
	//使用自由函数delete来删除指定的key
	fmt.Println("idNames删除前:", idNames)
	delete(idNames, 1)
	delete(idNames, 100) //删除不存在的key，不会有任何影响
	fmt.Println("idNames删除后:", idNames)

	//并发任务处理时 需要对Map进行上锁 //TODO
}

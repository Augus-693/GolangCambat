package main

import (
	"fmt"
	"io"
	"net/http"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/7 10:28
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {

	http.HandleFunc("/user", func(write http.ResponseWriter, request *http.Request) {
		//request：包含客户端发来的数据
		fmt.Println("用户请求详情：", request)

		//write：通过write将数据返回给客户端
		_, _ = io.WriteString(write, "这是/user请求返回的值")
	})

	http.HandleFunc("/name", func(write http.ResponseWriter, request *http.Request) {
		//request：包含客户端发来的数据
		fmt.Println("用户请求详情：", request)

		//write：通过write将数据返回给客户端
		_, _ = io.WriteString(write, "这是/name请求返回的值")
	})

	http.HandleFunc("/id", func(write http.ResponseWriter, request *http.Request) {
		//request：包含客户端发来的数据
		fmt.Println("用户请求详情：", request)

		//write：通过write将数据返回给客户端
		_, _ = io.WriteString(write, "这是/id请求返回的值")
	})

	fmt.Println("http server start......")

	if err := http.ListenAndServe(":8654", nil); err != nil {
		fmt.Println("http start err", err)
		return
	}

	/*
		err := http.ListenAndServe(":8654", nil)
		if err != nil {
			fmt.Println("http start err", err)
			return
		}
	*/

}

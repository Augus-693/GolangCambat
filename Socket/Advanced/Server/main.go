package main

import (
	"fmt"
	"net"
	"strings"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 18:27
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func handleFunc(conn net.Conn) {
	for {
		//创建一个容器 用于接收读取到的数据
		buf := make([]byte, 1024)

		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("net.Read error: ", err)
			return
		}

		fmt.Println("Client ===》 Server，长度：", cnt, "数据：", string(buf[:cnt]))

		//服务器对客户端请求进行响应
		//将数据转为大写
		upperData := strings.ToUpper(string(buf[:cnt]))
		cnt, err = conn.Write([]byte(upperData))
		if err != nil {
			fmt.Println("net.write error: ", err)
			return
		}
		fmt.Println("Client 《=== Server，长度：", cnt, "数据：", upperData)
	}

	//关闭服务器
	conn.Close()
}

func main() {

	//创建监听
	ip := "127.0.0.1"
	port := 8654
	address := fmt.Sprintf("%s:%d", ip, port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net.Listen error: ", err)
		return
	}

	for {
		fmt.Println("监听中......")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net.Accept error:", err)
			return
		}
		fmt.Println("连接建立成功！")

		go handleFunc(conn)
	}
}

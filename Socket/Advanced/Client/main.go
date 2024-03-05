package main

import (
	"fmt"
	"net"
	"time"
)

/**
 * @Project GolangCombat
 * @File    main.go
 * @Author  Augus Lee
 * @Date    2024/3/5 18:44
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	//与服务器建立连接
	conn, err := net.Dial("tcp", ":8654")
	if err != nil {
		fmt.Println("net.Dial error: ", err)
		return
	}
	fmt.Println("client与server连接建立成功")
	sendData := []byte("helloworld")

	for {
		//向服务器发送数据
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("net.Write error: ", err)
			return
		}
		fmt.Println("Client ===》 Server，长度：", cnt, "数据：", string(sendData))

		//接受服务器返回的数据
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("net.Read error: ", err)
			return
		}

		fmt.Println("Client 《=== Server，长度：", cnt, "数据：", string(buf[0:cnt]))
		time.Sleep(time.Second)
	}
	//关闭客户端
	conn.Close()
}

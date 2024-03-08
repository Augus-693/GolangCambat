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
 * @Date    2024/3/6 13:53
 * @GoVersion go1.22.0
 * @Version 1.0
 */

func main() {
	client := http.Client{}
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("http.Get failed: ", err)
		return
	}

	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("Content-Type: ", ct)
	fmt.Println("Date: ", date)
	fmt.Println("Server: ", server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status
	fmt.Println("url: ", url)
	fmt.Println("code: ", code)
	fmt.Println("status: ", status)

	body := resp.Body
	readBodyStr, err := io.ReadAll(body)
	if err != nil {
		fmt.Println("read body err: ", err)
		return
	}
	fmt.Println("body string: ", string(readBodyStr))

}

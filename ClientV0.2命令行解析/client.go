package main

import (
	"flag"
	"fmt"
	"net"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
}

func NewClient(serverIP string, serverPort int) *Client {
	//创建客户端对象
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
	}

	//链接server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn

	//返回对象
	return client
}

var serverIP string
var serverPort int

//./client -ip 127.0.0.1

func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器IP地址(默认是127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口(默认是8888)")
}

func main() {
	//命令行解析
	flag.Parse()

	client := NewClient("127.0.0.1", 8888)

	if client == nil {
		fmt.Println(">>>>>链接服务器失败...")
		return
	}

	fmt.Println(">>>>>链接服务器成功")

	//启动客户端业务
	select {}
}

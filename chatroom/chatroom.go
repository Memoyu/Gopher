package main

import (
	"fmt"
	"net"
	"strings"
)

// 用户信息结构
type User struct {
	id   string      // 用户Id
	name string      // 用户名称
	msg  chan string // 消息管道
}

// 定义全局用户map，用于保存所有连接用户
var users = make(map[string]User)

//定义消息全局管道，用于接收用户发来的消息
var messageChan = make(chan string, 10)

func main() {
	// 创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}
	fmt.Println("服务器启动成功，监听中")

	// 启用消息处理go程，用于监听消息并发送给每个用户
	go handlerBroadcast()

	// 遍历接收连接
	for {
		fmt.Println("=========主go程监听中===========")
		// 监听
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}
		fmt.Println("建立连接成功")

		// 启动处理业务的go程
		go handlerUserConn(conn)
	}
}

// 处理用户连接后以及读取用户发送的消息
func handlerUserConn(conn net.Conn) {
	fmt.Println("启动用户业务")
	// 获取客户端Ip and port
	userAddr := conn.RemoteAddr().String()
	user := User{
		id:   userAddr,          // 使用客户端信息作为唯一Id
		name: userAddr,          // 同时作为初始化name
		msg:  make(chan string), // 初始化聊天信息管道,加10是为了使用带缓冲的chan,否则在出管道时会阻塞
	}
	// 将用户存入全局users map
	users[user.id] = user

	//定义退出管道，用于处理用户退出
	quitChan := make(chan bool)

	// 启动用户退出go程
	go handlerWatchQuit(&user, conn, quitChan)

	// 启动用户消息回写go程
	go handlerWriteBackMessgaeClient(&user, conn)

	// 向messageChan中写入连接消息,达到广播所有用户该用户已上线
	messageChan <- fmt.Sprintf("[%s]:[%s] ===> online", user.id, user.name)

	// 遍历接收消息
	for {
		// 读取客户端数据
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)

		// 消息长度为0说明用户已经主动退出
		if cnt == 0 {
			fmt.Println("用户主动退出")
			quitChan <- true
		}

		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		message := string(buf[:cnt-1])

		// 单纯只输入回车
		if len(message) == 0 {
			continue
		}

		// 校验命令输入
		first := message[:1]
		if first == ">" {
			cmd := message[1:]
			cmdToLower := strings.ToLower(cmd)
			if len(cmd) == 3 && cmdToLower == "who" {
				var userStrs []string
				for _, user := range users {
					userStr := fmt.Sprintf("UserId:%s, UserName:%s", user.id, user.name)
					userStrs = append(userStrs, userStr)
				}

				msg := strings.Join(userStrs, "\n")
				user.msg <- msg
			} else if strings.Contains(cmdToLower, "rename") && cmdToLower[:6] == "rename" {
				newName := strings.Split(cmd, "|")[1]
				user.name = newName
				users[user.id] = user
				user.msg <- fmt.Sprintf("Rename Successfully NewName is %s", user.name)
			}
		} else {
			messageChan <- fmt.Sprintf("[%s]:%s", user.name, message)
			fmt.Println("服务器接收到客户端发来的数据为:", message, ", cnt:", cnt)
		}
	}
}

// 向所有用户广播消息，每当用户发送消息时
func handlerBroadcast() {
	fmt.Println("启动消息广播业务")
	defer fmt.Println("broadcast 退出")

	for {
		// 读取消息管道中的消息
		message := <-messageChan

		// 将消息发送给每个用户
		for _, user := range users {
			// 如果msg 是非缓存的，则会在这里阻塞
			user.msg <- message
		}
	}
}

// 将消息写回客户端
func handlerWriteBackMessgaeClient(user *User, conn net.Conn) {
	fmt.Println("启动用户消息回写业务")
	for message := range user.msg {
		fmt.Println("进行消息读取-msg:", message)
		_, _ = conn.Write([]byte(message))
	}
}

func handlerWatchQuit(user *User, conn net.Conn, quitChan <-chan bool) {

	for {
		select {
		case <-quitChan:
			fmt.Println("删除当前用户:", user.name)
			delete(users, user.id)
			messageChan <- fmt.Sprintf("[%s]:[%s] ===> exit", user.id, user.name)
			conn.Close()
			return
		}

	}
}

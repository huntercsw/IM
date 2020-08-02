package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"runtime/debug"
	"time"
)

type Client struct {
	Addr          string          // 客户端地址
	Socket        *websocket.Conn // 用户连接
	Send          chan []byte     // 待发送的数据
	AppId         uint32          // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	FirstTime     uint64          // 首次连接事件
	HeartbeatTime uint64          // 用户上次心跳时间
	LoginTime     uint64          // 登录时间 登录以后才有
}

type MessageContent struct {
	AppId   string      `json:"app_id"`
	UserId  string      `json:"user_id"`
	Message interface{} `json:"message"`
}

type Message struct {
	Seq  string         `json:"seq"`
	Cmd  string         `json:"cmd"`
	Data MessageContent `json:"data"`
}

func NewClient(addr string, socket *websocket.Conn, firstTime uint64) *Client {
	return &Client{
		Addr:          addr,
		Socket:        socket,
		Send:          make(chan []byte, 500),
		AppId:         0,
		UserId:        "",
		FirstTime:     firstTime,
		HeartbeatTime: 0,
		LoginTime:     0,
	}
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		log.Println("读取客户端数据 关闭send", c)
		close(c.Send)
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			log.Println("读取客户端数据 错误", c.Addr, err)
			return
		}

		readMessageHandler(message)

		time.Sleep(time.Millisecond)
	}
}

func readMessageHandler(message []byte) (err error) {
	msg := new(Message)
	if err = json.Unmarshal(message, msg); err != nil {
		return err
	}

	switch msg.Cmd {
	case "login":
		log.Println("login msg:", msg.Data)
	case "heartbeat":
		log.Println("heartbeat msg:", msg.Data)
	default:
		log.Println("default msg:", msg.Data)
	}

	return nil
}

func (c *Client) write() {
	defer func() {
		if r := recover(); r != nil {
			Logger.Error("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		c.Socket.Close()
		log.Println("Client发送数据 defer", c)
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				log.Println("Client发送数据 关闭连接", c.Addr, "ok", ok)
				return
			}
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (c *Client) SendMsg(msg []byte) {

	if c == nil {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Println("SendMsg stop:", r, string(debug.Stack()))
		}
	}()

	c.Send <- msg
}

func (c *Client) close() {
	close(c.Send)
}

package socket

import (
	"github.com/gorilla/websocket"
)

// Result 数据的返回结构
type Result struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

// Channel 装饰用户连接
type Channel struct {
	conn *websocket.Conn // WebSocket 连接
	send chan Result     // 传出的 packets 队列
}

func NewChannel(conn *websocket.Conn) *Channel {

	c := &Channel{
		conn: conn,
		send: make(chan Result),
	}

	go c.reader()
	//go c.writer()

	return c
}

func (c Channel) reader() {
	//buf := bufio.NewReader(c.conn)
}

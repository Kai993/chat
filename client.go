package main

import "github.com/gorilla/websocket"

// client : チャットを行っている1人のユーザー
type client struct {
	// socket : クライアントのためのWebSocket
	socket *websocket.Conn

	// send : メッセージが送られるチャネル
	send chan []byte

	// room : クライアントが参加しているチャットルーム
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

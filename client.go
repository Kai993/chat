package main

import (
	"time"

	"github.com/gorilla/websocket"
)

// client : チャットを行っている1人のユーザー
type client struct {
	// socket : クライアントのためのWebSocket
	socket *websocket.Conn

	// send : メッセージが送られるチャネル
	send chan *message

	// room : クライアントが参加しているチャットルーム
	room *room

	// userData : ユーザーに関する情報を保持する
	userData map[string]interface{}
}

// read : メッセージ受信
func (c *client) read() {
	for {
		var msg *message

		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			msg.AvatarURL, _ = c.room.avatar.GetAvatarURL(c)
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// write : メッセージ送信
func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}

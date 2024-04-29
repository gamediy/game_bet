package model

import (
	"context"
	"github.com/goccy/go-json"
	"golang.org/x/net/websocket"
	"log"
)

const (
	HEART_CHECK = "HEART_CHECK"
	GAME_ISSUE  = "GAME_ISSUE"
)

type Message struct {
	Code    string      `json:"code"`
	Content interface{} `json:"content"`
	Message string      `json:"message"`
}

func MessageClientHandle(cancel context.CancelFunc, buf []byte, user *User) {
	message := Message{}
	err := json.Unmarshal(buf, &message)
	if err != nil {
		log.Println(err)

		user.WsConn.WriteMessage(
			websocket.TextFrame,
			[]byte("err:"),
		)
		return
	}
	switch message.Code {

	}
	marshal, _ := json.Marshal(&message)
	err = user.WsConn.WriteMessage(websocket.TextFrame, marshal)
	if err != nil {
		cancel()
	}
	return
}

func MessageServerHandle(mesage *Message) error {
	switch mesage.Code {

	}
	return nil
}

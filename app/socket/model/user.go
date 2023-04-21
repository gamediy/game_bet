package model

import (
	"bet/core/auth"
	"context"
	"github.com/gorilla/websocket"
	"log"
)

type User struct {
	UserInfo *auth.UserInfo
	WsConn   *websocket.Conn
	SendChan chan []byte
}

func NewUser(conn *websocket.Conn, userinfo *auth.UserInfo) {
	var user = &User{
		UserInfo: userinfo,
		WsConn:   conn,
		SendChan: make(chan []byte),
	}
	ctx, cancelFunc := context.WithCancel(context.Background())

	user.online()
	go user.recvMessage(ctx, cancelFunc)
	go user.sendMessage(ctx, cancelFunc)

}
func (user *User) recvMessage(ctx context.Context, cancel context.CancelFunc) {
	defer user.offline(cancel)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			_, p, err := user.WsConn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			user.SendChan <- p
			log.Printf("Recv [%s] msg:%s", user.UserInfo.Account, p)
		}

	}
}

func (user *User) sendMessage(ctx context.Context, cancel context.CancelFunc) {
	defer user.offline(cancel)
	for {
		select {
		case <-ctx.Done():
			return
		case buf := <-user.SendChan:
			go MessageClientHandle(cancel, buf, user)
			log.Printf("Send [%s] msg:%s", user.UserInfo.Account, buf)

		}

	}
}

func (user *User) online() {
	server.AddOnlineUserMap(user)
	log.Printf("[%v] 个在线用户", len(server.OnlineUser))
	log.Printf("[%s] 上线了", user.UserInfo.Account)
}

func (user *User) offline(cancel context.CancelFunc) {
	cancel()
	user.WsConn.Close()
	server.DeleteOnlineUserMap(user)
	log.Printf("[%s] 下线了", user.UserInfo.Account)

}

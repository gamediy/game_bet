package model

import (
	"sync"
)

type Server struct {
	OnlineUser  map[string]*User
	UserMapLock sync.RWMutex
}

var server = Server{
	OnlineUser: make(map[string]*User),
}

func (server *Server) AddOnlineUserMap(user *User) {
	server.UserMapLock.Lock()
	server.OnlineUser[user.UserInfo.Account] = user
	server.UserMapLock.Unlock()
}

func (server *Server) DeleteOnlineUserMap(user *User) {
	server.UserMapLock.Lock()
	delete(server.OnlineUser, user.UserInfo.Account)
	server.UserMapLock.Unlock()

}

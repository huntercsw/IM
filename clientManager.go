package main

import (
	"fmt"
	"sync"
)

type ClientManager struct {
	ClientDict  map[string]string // [user_id的开头]hostInfo, 根据user_id获取连接所在的主机
	ClientsLock sync.RWMutex
	Users       map[string]*Client // [user_id]*Client, 本机应该维护的所有用户map， 用户不在线的nil
	UserLock    sync.RWMutex
	Register    chan *Client
	Login       chan *Client
	Unregister  chan *Client
	Broadcast   chan []byte
}

func NewClientManager() *ClientManager {
	cm := &ClientManager{
		ClientDict:  make(map[string]string),
		ClientsLock: sync.RWMutex{},
		Users:       make(map[string]*Client),
		UserLock:    sync.RWMutex{},
		Register:    nil,
		Login:       nil,
		Unregister:  nil,
		Broadcast:   nil,
	}

	if err := cm.clientDictionaryInit(); err != nil {
		panic(fmt.Sprintf("client dictionary init error: %f", err))
	}

	return cm
}

func (cm *ClientManager) clientDictionaryInit() (err error) {
	//TODO get all host metadata from etcd

	return nil
}

func (cm *ClientManager) localClientsInit() {
	//TODO get all users which are allocated to local host

	userList := make([]string, 100000)

	//TODO user list init

	for _, userId := range userList {
		cm.Users[userId] = nil
	}
}

func (cm *ClientManager) UserSignUp(userId string) {

}

func (cm *ClientManager) UserLogin(userId string) {

}

func (cm *ClientManager) UserLogout(userId string) {

}
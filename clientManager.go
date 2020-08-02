package main

import "sync"

type ClientManager struct {
	ClientMap   map[string]*Client // [user_id]*Client
	ClientsLock sync.RWMutex
	Users       map[string]*Client
	UserLock    sync.RWMutex
	Register    chan *Client
	Login       chan *Client
	Unregister  chan *Client
	Broadcast   chan []byte
}

func (cm *ClientManager) NewClientManager() *ClientManager {
	return &ClientManager{
		ClientMap:   make(map[string]*Client),
		ClientsLock: sync.RWMutex{},
		Users:       nil,
		UserLock:    sync.RWMutex{},
		Register:    nil,
		Login:       nil,
		Unregister:  nil,
		Broadcast:   nil,
	}
}
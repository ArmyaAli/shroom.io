package websocket

import (
	"game_server/pkg/data"
	"net/http"
	"sync"
)

type Message struct {
	Channel   string      `json:"channel"`
	SessionId string      `json:"sessionid"`
	Timestamp float64     `json:"timestamp"`
	Payload   data.Player `json:"payload"`
}

type MessageQueue struct {
	store []Message
	size  int
}

type Session struct {
	Id     string
	Email  string
	Cookie http.Cookie
}

type _SessionMap map[string]Session

type LockingSessionMap struct {
	store _SessionMap
	mu    sync.Mutex
}

func InitLockingSessionMap() *LockingSessionMap {
	return &LockingSessionMap{
		store: make(map[string]Session),
	}
}

func (lsm *LockingSessionMap) Get(key string) Session {
	lsm.mu.Lock()
	defer lsm.mu.Unlock()
	value, exists := lsm.store[key]

	if !exists {
		return Session{}
	}

	return value
}

func (lsm *LockingSessionMap) Set(key string, sesh Session) {
	lsm.mu.Lock()
	defer lsm.mu.Unlock()
	lsm.store[key] = sesh
}

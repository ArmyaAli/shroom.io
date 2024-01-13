package data

import (
	"sync"
)

type Vector2 struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Player struct {
	Id   string  `json:"id"`
	Nick string  `json:"nick"`
	Pos  Vector2 `json:"pos"`
	Vel  Vector2 `json:"vel"`
}

type LockingPlayerBuffer struct {
	mu    sync.Mutex
	store [1024]Player
	curr  int
}

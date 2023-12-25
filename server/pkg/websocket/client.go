package websocket

// Contains information about the client 

import (
	"github.com/gorilla/websocket"
	"time"
)

type Player struct {
  Channel string`json:"channel"`
	Id     int32 `json:"id"`
	Name   string
	X      int32
	Y      int32
	active bool
}

type Client struct {
	Id     int32 `json:"id"`
	Conn   *websocket.Conn
	TLA    time.Time // Last Active timestamp
	Active bool
}

type Session struct {
	Id            int32
	Session_state string
	Active_for    time.Time
}

type Message struct {
  Channel string
	Player  Player
}

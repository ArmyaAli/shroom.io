package websocket

import (
  "game_server/pkg/data"
)

type Message struct {
  Channel   string `json:"channel"`
  SessionId string `json:"sessionid"`
	Timestamp float64 `json:"timestamp"`
	Payload   data.Player `json:"payload"`
}

type MessageQueue struct {
  store []Message
  size int
}

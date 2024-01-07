package websocket

import (
	"fmt"
  "game_server/pkg/data"
)

type channelOperation func(*Message) 

var DispatchMap map[string]channelOperation

func Init_channels() {
   DispatchMap = make(map[string]channelOperation, 64)
   DispatchMap["player"] = player_channel
   DispatchMap["position"] = position_channel
}

func player_channel(m *Message) {
  // Register our player
  fmt.Println("player_channel")
  fmt.Println(m.Payload.Id)
  
  data.PLAYER_MAP.Set(m.SessionId, m.Payload)
  m.SessionId = m.Payload.Id
}
            
func position_channel(m *Message) {
  fmt.Println("position_channel")
  fmt.Println(m.Payload.Id, m.Payload.Pos.X, m.Payload.Pos.Y)

  data.PLAYER_MAP.Set(m.SessionId, m.Payload)
}

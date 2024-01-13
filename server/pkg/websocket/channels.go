package websocket

import (
	"fmt"
	"game_server/pkg/data"
	"github.com/gorilla/websocket"
	"sync"
)

type channelOperation func(*Message, *websocket.Conn)

type LockingDispatchMap struct {
	store map[string]channelOperation
	mu    sync.Mutex
}

var DispatchMap LockingDispatchMap

func (ldm *LockingDispatchMap) Set(key string, routine channelOperation) {
	ldm.mu.Lock()
	defer ldm.mu.Unlock()
	ldm.store[key] = routine
}

func (ldm *LockingDispatchMap) Get(key string) channelOperation {
	ldm.mu.Lock()
	defer ldm.mu.Unlock()
	val, exists := ldm.store[key]

	if !exists {
		return nil
	}

	return val
}

func InitChannels() {
	DispatchMap.store = make(map[string]channelOperation, 64)
	DispatchMap.store["player"] = player_channel
	DispatchMap.store["position"] = position_channel
}

func player_channel(m *Message, conn *websocket.Conn) {
	// Register our player
	//fmt.Println("player_channel")
	//fmt.Println(m.Payload.Id)

	data.PlayerBuffer.Add(m.Payload)

	for _, p := range data.PlayerBuffer.GetBuffer() {
		fmt.Println(p.Id, p.Nick)
		// If empty break
		if p == (data.Player{}) {
			break
		}
	}
}

func position_channel(m *Message, conn *websocket.Conn) {
	//fmt.Println("position_channel")
	//fmt.Println(m.Payload.Id, m.Payload.Pos.X, m.Payload.Pos.Y)
	// Check if player already in buffer
	id := m.Payload.Id
	found := false

	for _, p := range data.PlayerBuffer.GetBuffer() {
		if id == p.Id {
			found = true
			break
		}
	}

	// Only add to player buffer on the same Id
	if found {
		data.PlayerBuffer.Set(id, m.Payload)
	}

}

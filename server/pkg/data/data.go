package data

import (
	"fmt"
)

var PlayerBuffer = InitLockingPlayerBuffer()

func InitLockingPlayerBuffer() *LockingPlayerBuffer {
	return &LockingPlayerBuffer{
		store: [1024]Player{},
		curr:  0,
	}
}

func (lpb *LockingPlayerBuffer) Add(p Player) {
	lpb.mu.Lock()
	defer lpb.mu.Unlock()
	lpb.store[lpb.curr] = p
	lpb.curr++
}

func (lpb *LockingPlayerBuffer) Set(id string, p Player) {
	lpb.mu.Lock()
	defer lpb.mu.Unlock()

	i := 0
	for _, _p := range lpb.store {
		if id == _p.Id {
			//fmt.Println("Setting")
			lpb.store[i] = p
			break
		}
		i++
	}
}

func (lpb *LockingPlayerBuffer) Evict(id string) {
	lpb.mu.Lock()
	defer lpb.mu.Unlock()

	i := 0
	for _, p := range lpb.store {
		if id == p.Id {
			lpb.store[i] = Player{}
			break
		}
		i++
	}
}

func (lpb *LockingPlayerBuffer) Length() int {
	lpb.mu.Lock()
	defer lpb.mu.Unlock()

	i := 0
	for _, p := range lpb.store {
		if p == (Player{}) {
			break
		}
		i++
	}

	fmt.Println(i + 1)
	return i + 1
}

func (lpb *LockingPlayerBuffer) GetBuffer() [1024]Player {
	lpb.mu.Lock()
	defer lpb.mu.Unlock()
	return lpb.store
}

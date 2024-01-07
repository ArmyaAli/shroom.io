package data

var PLAYER_MAP = Init_PLAYER_MAP()

func Init_PLAYER_MAP() *CONCURRENT_PLAYER_MAP {
	return &CONCURRENT_PLAYER_MAP{
		PLAYER_MAP: make(map[string]Player, 1024),
	}
}

func(m *CONCURRENT_PLAYER_MAP) Set(key string, val Player) {
  m.mu.Lock()
  defer m.mu.Unlock()
  m.PLAYER_MAP[key] = val
}

func(m *CONCURRENT_PLAYER_MAP) Get(key string) (*Player, bool) {
  m.mu.Lock()  
  defer m.mu.Unlock()  
  val, ok := m.PLAYER_MAP[key]
  return &val, ok
}


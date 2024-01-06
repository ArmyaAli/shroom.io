package data

type Vector2 struct {
  X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Player struct {
  Id   string `json:"id"`
	Nick string `json:"nick"`
	Pos  Vector2 `json:"pos"`
	Vel  Vector2 `json:"vel"`
}

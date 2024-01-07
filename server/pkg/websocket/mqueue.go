package websocket 

func (q *MessageQueue) Enqueue(val Message) {
  q.store = append(q.store, val)
  q.size--
}

func (q *MessageQueue) Dequeue() (Message, bool) {
  if len(q.store) == 0 {
    return Message{}, false
  }

  value := q.store[0]
  q.store = q.store[1:]
  q.size--

  return value, true
}

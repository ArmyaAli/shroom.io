package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v5"
	"log"
	"net/http"
  "time"
  "encoding/json"
  "game_server/pkg/data"
)

var upgrader = websocket.Upgrader{} // use default options
var ws_conn websocket.Conn

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// TODO(Ali): Create a list of acceptable origins and check the request Host value against the list
	upgrader.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

// Websocket connection
func Init_ws(ctx echo.Context) error {
	c, err := Upgrade(ctx.Response().Unwrap(), ctx.Request())

	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

  // Start our message listener

  // Start our Reader

  Init_channels()

  LISTNER_reader(c) 
  // Start ticker
  TIMER_playerUpdate(c)
  return nil
}

func LISTNER_reader(conn *websocket.Conn) {
  go func() {
    for {
      err := Reader(conn)

      if err != nil {
        break
      }
    }
  }()
}

func Reader(conn *websocket.Conn) error {
    var message Message

		_, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return err
		}

    // Read the data into player
    err = json.Unmarshal(p, &message)

    if err != nil {
        fmt.Println("Error:", err)
        return err
    }
    
    DispatchMap[message.Channel](&message)

    return nil
}

func writeMessage(message string, conn *websocket.Conn, kill chan bool) {
    if conn == nil {
      fmt.Print("Connection Lost")
      return
    }

    err := conn.WriteMessage(websocket.TextMessage, []byte(message))

		if err != nil {
      fmt.Println("Error", err)
      // If connection closed evict the client from the map somehow
      kill <-true
			return
		}
}


func TIMER_playerUpdate(conn *websocket.Conn) {
	ticker := time.NewTicker(15 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
        ticker.Stop();
        fmt.Println("Exiting routine")
				return
			case _ = <-ticker.C:
        json, err := json.Marshal(data.PLAYER_MAP)
        if err != nil {
          fmt.Println("Error converting to json")
        }

        writeMessage(string(json), conn, done)
			}
		}
	}()
}

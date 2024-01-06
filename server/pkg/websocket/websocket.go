package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v5"
	"log"
	"net/http"
  "time"
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

  // Start ticker
  TIMER_playerUpdate(c)
  
  return nil
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}


		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func writeMessage(message string, conn *websocket.Conn, kill chan bool) {

    if conn == nil {
      fmt.Print("Connection Lost")
      return
    }

		fmt.Println("Sending")

    p := data.Player{
      Id: "1111", 
      Nick: "lordmushroom", 
      Pos: data.Vector2{X: 33, Y: 100}, 
      Vel: data.Vector2{X: 10, Y: 10},
    }

    err := conn.WriteJSON(p)        

		if err != nil {
      if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
        fmt.Println("Unexpected Close Error:", err)
      }
      fmt.Println("Error", err)
      kill <-true
      // End the timer
			return
		}
}


func TIMER_playerUpdate(conn *websocket.Conn) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
        ticker.Stop();
        fmt.Print("Exiting routine")
				return
			case t := <-ticker.C:
        // Sample array to send
        fmt.Print(t)
        writeMessage("Hello", conn, done)
			}
		}
	}()
}

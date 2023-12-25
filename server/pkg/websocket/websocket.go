package websocket

import (
	"fmt"
  "encoding/json"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v5"
	"io"
	"log"
	"net/http"
	//"time"
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

	defer c.Close()

  for {
    mt, message, err := c.ReadMessage()
    var receivedData Message

    json.Unmarshal(message, &receivedData)
        
    if err != nil {
      log.Println("read:", err)
      break
    }

    log.Printf("recv: %s", receivedData.Channel)
    log.Printf("recv: %s", receivedData.Player.Name)

    err = c.WriteMessage(mt, message)

    if err != nil {
      log.Println("write:", err)
      break
    }
  }

  return nil
}

func Reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func Writer(conn *websocket.Conn) {
	for {
		fmt.Println("Sending")

		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}

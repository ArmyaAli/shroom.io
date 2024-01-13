package websocket

import (
	"encoding/json"
	"fmt"
	"game_server/pkg/data"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v5"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{} // use default options
var ws_conn websocket.Conn

func RegisterClient(ctx echo.Context, rc chan string) error {
	fmt.Println("Setting cookie")

	id, email := ctx.QueryParam("id"), ctx.QueryParam("email")
	w := ctx.Response().Unwrap()
	r := ctx.Request()

  ctx.Response().Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:9000")
  ctx.Response().Header().Set("Access-Control-Allow-Credentials", "true")

	cookie, err := r.Cookie("registration_cookie")

	fmt.Println(cookie, err)

	if err == nil {
    fmt.Println(cookie.Value)
		rc <- cookie.Value
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "Client already registered",
		})
	}

	NewSession(id, email)
	sesh := SessionMap.Get(id)

	fmt.Println(sesh.Cookie)


	http.SetCookie(w, &sesh.Cookie)
	// Returning a JSON response
	responseData := map[string]interface{}{
		"message": "Registration successful",
	}

	rc <- sesh.Id

	return ctx.JSON(http.StatusOK, responseData)
}

// Websocket connection
func InitWebsocket(ctx echo.Context, rc chan string) error {
	// Before we upgrade our connection let's do some housekeeping and administrative work
	registered := <-rc

  fmt.Println("Registered Id", registered)

	ctx.Response().Unwrap()
	c, err := Upgrade(ctx.Response().Unwrap(), ctx.Request())

	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

	InitChannels()
	Listener(c)

	// Start ticker
	UpdatePlayerPosition(c, registered)

	return nil
}

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

func Listener(conn *websocket.Conn) {
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

	//fmt.Println(message.Channel)
	// First message set up the session

	routine := DispatchMap.Get(message.Channel)

	if routine == nil {
		return nil
	}

	routine(&message, conn)

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
		kill <- true
		return
	}
}

func UpdatePlayerPosition(conn *websocket.Conn, id string) {
	ticker := time.NewTicker(30 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
        data.PlayerBuffer.Evict(id)
				return
			case _ = <-ticker.C:
				json, err := json.Marshal(data.PlayerBuffer.GetBuffer())

				if err != nil {
					fmt.Println("Error converting to json")
				}

				writeMessage(string(json), conn, done)
			}
		}
	}()
}

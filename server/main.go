package main;

import (
  //STD
	"log"
	"net/http"
	"os"
  //"io"
  "fmt"
  //"strings"
  "encoding/json"
	"github.com/gorilla/websocket"
  // Required for pocketbase
  "github.com/labstack/echo/v5"
  // Pocketbase
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Player struct {
  Id int32 `json:"id"`
  Name string;
  X int32;
  Y int32;
  active bool;
}


var upgrader = websocket.Upgrader{} // use default options
var Players [1024]Player;
var id_count int32 = 0;

func main() {
	log.Print("Server starting up")

  upgrader.CheckOrigin = checkOrigin
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

    // Register our websocket endpoint
		e.Router.GET("/websocket", initWebsocket)

		return nil
	})

	err := app.Start()

	if err != nil {
  	log.Fatal(err)
	}
}

// TODO(Ali): Create a list of acceptable origins and check the request Host value against the list
func checkOrigin(req *http.Request) bool {
  return true
}

// TODO(Ali): I need to implement broadcast of a limited amount of coordinates to a player when they login
// First i'll try to make it work and then add on fnality ;)
// At this point I just send back and forth a message between my client and server and we establish a correct

// Websocket connection 
func initWebsocket(ctx echo.Context) error {
  c, err := upgrader.Upgrade(ctx.Response().Unwrap(), ctx.Request(), nil)
  
  
	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

	defer c.Close()

	for {
	  var receivedData Player

		_, message, err := c.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			break
		}

    // Unmarshal JSON
		if err := json.Unmarshal(message, &receivedData); err != nil {
			log.Println("Error unmarshalling JSON:", err)
		}

    // Every successful connection let's create a player
    Players[id_count] = Player {
      Id: id_count,
      Name: receivedData.Name, 
      X: receivedData.X, 
      Y: receivedData.Y,
      active: receivedData.active,
    }
    
    // Send all the players to the client initially
    json_out, _ := json.Marshal(&Players)
    c.WriteMessage(1, json_out)

    // Access the JSON data
    fmt.Printf("Received content: %d\n", receivedData.Id)
    fmt.Printf("Received content: %s\n", receivedData.Name)
    fmt.Printf("Received content: %d\n", receivedData.X)
    fmt.Printf("Received content: %d\n", receivedData.Y)
	}

  id_count++;


  return nil
}

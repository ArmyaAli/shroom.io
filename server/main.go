package main;

import (
  //STD
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
  // Required for pocketbase
  "github.com/labstack/echo/v5"
  // Pocketbase
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)


var upgrader = websocket.Upgrader{} // use default options

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

  print("Hello World")
  
	if err != nil {
		log.Print("upgrade:", err)
		return nil
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}

	}

  return nil
}

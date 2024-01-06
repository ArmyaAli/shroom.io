package main

import (
	//STD
	"fmt"
	"log"
	"os"
	// Pocketbase
	"game_server/pkg/websocket"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	log.Print("Server starting up")

	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		// Register our websocket endpoint
		e.Router.GET("/websocket", websocket.Init_ws)

		return nil
	})
  
  if err := app.Start(); err != nil {
    fmt.Print("Hello")
  }

}

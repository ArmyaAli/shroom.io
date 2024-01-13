package main

import (
	//STD
	"fmt"
	"log"
	"os"
	// Pocketbase
	"game_server/pkg/websocket"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	log.Print("Server starting up")

	app := pocketbase.New()

	registerChannel := make(chan string, 2048)
	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
		// Register our websocket endpoint
		e.Router.GET("/register", func(c echo.Context) error { 
      return websocket.RegisterClient(c, registerChannel) 
    })
		e.Router.GET("/websocket", func(c echo.Context) error { 
      return websocket.InitWebsocket(c, registerChannel) 
    })

		return nil
	})

	if err := app.Start(); err != nil {
		fmt.Print("Hello")
	}

}

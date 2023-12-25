package main

import (
	//STD
	"fmt"
	"log"
	"os"
	"time"
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
  

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	//ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

  err := app.Start()
  
  if err != nil {
    fmt.Print("Hello")
  }

}

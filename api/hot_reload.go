package api

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func WithHotReload(app *fiber.App) {
	id := []byte(uuid.New().String())

	currentId := -1
	connections := map[int]*websocket.Conn{}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP)

	go func() {
		for {
			s := <-sig
			fmt.Printf("Received %s\n", s)

			id = []byte(uuid.New().String())

			for i, c := range connections {
				if err := c.WriteMessage(websocket.TextMessage, id); err != nil {
					fmt.Println("Removing socket", i)
					delete(connections, i)
				}
			}
		}
	}()

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/hotreload", websocket.New(func(c *websocket.Conn) {
		for {
			currentId++
			connections[currentId] = c

			if _, _, err := c.ReadMessage(); err != nil {
				break
			}

			if err := c.WriteMessage(websocket.TextMessage, id); err != nil {
				break
			}
		}
	}))

}

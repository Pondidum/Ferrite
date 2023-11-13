package api

import (
	"fmt"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func WithHotReload(app *fiber.App) {
	id := uuid.New()
	app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.

		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {

		fmt.Println("On Websocket Message")
		for {

			_, msg, err := c.ReadMessage()
			if err != nil {
				fmt.Println("read:", err)
				break
			}
			fmt.Printf("recv: %s\n", msg)

			if err := c.WriteMessage(websocket.TextMessage, []byte(id.String())); err != nil {
				fmt.Println("write:", err)
				break
			}
		}

	}))

}

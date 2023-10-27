package api

import (
	"ferrite/goes/sqlite"
	"ferrite/layout"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func NewApiV2(store *sqlite.SqliteStore) (*fiber.App, error) {

	app := fiber.New(fiber.Config{})

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("%s %s \n", c.Method(), c.Path())
		return c.Next()
	})

	app.Use(cors.New())

	// app.Get("/layouts", func(c *fiber.Ctx) error {

	// 	var names []string
	// 	if err := store.Query(AllLayoutNames(&names)); err != nil {
	// 		return err
	// 	}

	// })

	app.Get("/api/layouts/:name", func(c *fiber.Ctx) error {

		name := c.Params("name")
		viewJson := ""

		if err := store.Query(layout.GetLayoutRaw(name, &viewJson)); err != nil {
			return err
		}

		c.Set("content-type", fiber.MIMEApplicationJSON)
		return c.Send([]byte(viewJson))
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.Dir("./webui/dist"),
	}))

	return app, nil

}

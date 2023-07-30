package api

import (
	"ferrite/zmk"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func NewApi() (*fiber.App, error) {

	app := fiber.New(fiber.Config{
		// Views: engine,

	})

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("%s %s \n", c.Method(), c.Path())
		return c.Next()
	})

	app.Use(cors.New())

	kb, err := zmk.ReadKeyboardInfo("./config/keyboard.json")
	if err != nil {
		return nil, err
	}

	keys, err := zmk.ReadKeys()
	if err != nil {
		return nil, err
	}

	zmkKeyMap := zmk.BuildKeyMap(keys)

	app.Get("/api/zmk/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]any{
			"layout": kb.Layout,
			"keys":   zmkKeyMap,
		})
	})

	app.Get("/api/device", func(c *fiber.Ctx) error {
		keymapFile, err := os.Open("./config/cradio.keymap")
		if err != nil {
			return err
		}

		zmkTree, err := zmk.Parse(keymapFile)
		if err != nil {
			return err
		}

		keymap := KeymapFromZmk(zmkKeyMap, zmkTree)
		return c.JSON(keymap)
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.Dir("./webui/dist"),
	}))

	return app, nil
}

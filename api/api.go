package api

import (
	"ferrite/keyboard"
	"ferrite/zmk"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/mitchellh/mapstructure"
)

func NewApi() (*fiber.App, error) {

	app := fiber.New(fiber.Config{
		// Views: engine,
	})

	app.Use(cors.New())

	kb, err := keyboard.ReadKeyboardInfo("./config/keyboard.json")
	if err != nil {
		return nil, err
	}

	keys, err := zmk.ReadKeys()
	if err != nil {
		return nil, err
	}

	app.Get("/api/zmk/", func(c *fiber.Ctx) error {
		fmt.Println("GET /api/zmk/")

		return c.JSON(map[string]any{
			"layout": kb.Layout,
			"keys":   zmk.BuildKeyMap(keys),
		})
	})

	app.Get("/api/device", func(c *fiber.Ctx) error {
		fmt.Println("GET /api/device")

		keymap, err := os.Open("./config/cradio.keymap")
		if err != nil {
			return err
		}

		f, err := zmk.Parse(keymap)
		if err != nil {
			return err
		}

		temp := map[string]any{}
		if err := mapstructure.Decode(f, &temp); err != nil {
			return err
		}

		response := File{}
		if err := mapstructure.Decode(temp, &response); err != nil {
			return err
		}

		return c.JSON(f)
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.Dir("./webui/dist"),
	}))

	return app, nil
}

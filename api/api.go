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

	app.Get("/api/zmk/", func(c *fiber.Ctx) error {
		return c.JSON(map[string]any{
			"layout": kb.Layout,
			"keys":   keys,
		})
	})

	app.Get("/api/device", func(c *fiber.Ctx) error {
		keymapFile, err := os.Open("./config/cradio.keymap")
		if err != nil {
			return err
		}
		defer keymapFile.Close()

		zmkTree, err := zmk.Parse(keymapFile)
		if err != nil {
			return err
		}

		keymap := KeymapFromZmk(zmkTree)
		return c.JSON(keymap)
	})

	app.Post("/api/device/binding", func(c *fiber.Ctx) error {

		var dto BindingUpdate
		if err := c.BodyParser(&dto); err != nil {
			return err
		}

		fmt.Printf("layer %v\n", dto.Layer)
		fmt.Printf("key: %v\n", dto.Key)

		fmt.Printf("binding: %+v\n", dto.Binding)

		keymapFile, err := os.Open("./config/cradio.keymap")
		if err != nil {
			return err
		}
		defer keymapFile.Close()

		zmkTree, err := zmk.Parse(keymapFile)
		if err != nil {
			return err
		}

		keymapFile.Close()

		zmkTree.Device.Keymap.Layers[dto.Layer].Bindings[dto.Key] = BehaviorFromBinding(&dto.Binding)

		outFile, err := os.Create("./config/cradio.keymap")
		if err != nil {
			return err
		}
		defer outFile.Close()

		kb, err := zmk.ReadKeyboardInfo("./config/keyboard.json")
		if err != nil {
			return err
		}
		zmk.Write(outFile, kb, zmkTree)

		return c.SendStatus(200)
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.Dir("./webui/dist"),
	}))

	return app, nil
}

func BehaviorFromBinding(binding *Binding) *zmk.Behavior {

	params := make([]*zmk.Binding, len(binding.Params))
	for i, p := range binding.Params {
		params[i] = &zmk.Binding{
			Number:  p.Number,
			KeyCode: p.KeyCode,
		}
	}
	return &zmk.Behavior{
		Action: binding.Action,
		Params: params,
	}
}

type BindingUpdate struct {
	Layer   int
	Key     int
	Binding Binding
}

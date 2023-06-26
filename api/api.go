package api

import (
	"ferrite/zmk"
	"fmt"
	"net/http"
	"os"
	"reflect"

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

	kb, err := zmk.ReadKeyboardInfo("./config/keyboard.json")
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
		d, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
			DecodeHook: KeyCodeToKeysHookFunc(),
			Result:     &response,
		})
		if err != nil {
			return err
		}

		if err := d.Decode(temp); err != nil {
			return err
		}

		return c.JSON(response.Device)
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root: http.Dir("./webui/dist"),
	}))

	return app, nil
}

func KeyCodeToKeysHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if t != reflect.TypeOf([]string{}) {
			return data, nil
		}

		if f.Kind() == reflect.Pointer && f.Elem().Kind() == reflect.String {
			str := data.(*string)
			keys := parseKeys(*str)
			return keys, nil
		}

		if f.Kind() == reflect.String {
			str := data.(string)
			keys := parseKeys(str)
			return keys, nil
		}

		return data, nil
	}
}

func parseKeys(input string) []string {

	keys := []string{}

	current := []rune{}
	for _, char := range input {

		if char == '(' {
			keys = append(keys, string(current))
			// keys.push(current + "(code)"); // modifiers are defined as "LS(code)"
			current = []rune{}
		} else if char == ')' {
			break
		} else {
			current = append(current, char)
		}
	}

	if len(current) > 0 {
		keys = append(keys, string(current))
	}

	return keys
}

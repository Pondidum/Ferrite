package api

import (
	"ferrite/devices"
	"ferrite/goes/sqlite"
	"ferrite/layout"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"

	"github.com/gofiber/template/html/v2"
)

func NewApiV2(store *sqlite.SqliteStore) (*fiber.App, error) {

	engine := html.New("./app", ".html")
	engine.Reload(true)

	engine.AddFunc("dict", func(v ...interface{}) map[string]interface{} {
		dict := map[string]interface{}{}
		lenv := len(v)
		for i := 0; i < lenv; i += 2 {
			key := fmt.Sprint(v[i])
			if i+1 >= lenv {
				dict[key] = ""
				continue
			}
			dict[key] = v[i+1]
		}
		return dict
	})

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	WithHotReload(app)

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("%s %s \n", c.Method(), c.Path())
		return c.Next()
	})

	app.Use(cors.New())

	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.Dir("./app/static"),
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		names, err := layout.GetLayoutNames(store)
		if err != nil {
			return err
		}

		values := fiber.Map{
			"LayoutNames": names,
		}

		return c.Render("index", values, "layouts/main")
	})

	app.Get("/layouts/:name", func(c *fiber.Ctx) error {
		return c.Redirect(c.Path() + "/0")
	})

	app.Get("/layouts/:name/:layer/:key?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		layer, _ := c.ParamsInt("layer", 0)
		key, _ := c.ParamsInt("key", -1)

		view, err := sqlite.ViewByProperty[layout.LayoutView](store, "$.name", name)
		if err != nil {
			return err
		}

		layoutNames, err := layout.GetLayoutNames(store)
		if err != nil {
			return err
		}

		values := fiber.Map{
			"LayoutNames":       layoutNames,
			"LayoutName":        name,
			"Layers":            view.Keymap.Layers,
			"CurrentLayer":      view.Keymap.Layers[layer],
			"CurrentLayerIndex": layer,
		}

		if view.Device != "" {
			kb, err := devices.ReadDevice(view.Device)
			if err != nil {
				return err
			}

			values["Device"] = kb
		}

		if key > -1 {
			values["IsEdit"] = true
			values["KeyIndex"] = key
			values["Key"] = view.Keymap.Layers[layer].Bindings[key]
		}

		return c.Render("layout", values, "layouts/main")
	})

	// app.Get("/api/layouts/:name", func(c *fiber.Ctx) error {

	// 	name := c.Params("name")
	// 	viewJson, err := layout.GetLayoutRaw(store, name)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	c.Set("content-type", fiber.MIMEApplicationJSON)
	// 	return c.Send([]byte(viewJson))
	// })

	return app, nil

}

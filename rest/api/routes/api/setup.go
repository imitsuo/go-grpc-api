package routes

import (
	"grcp-api/api/configs"
	"net/http"

	"github.com/gofiber/fiber"
)

func Hello(c *fiber.Ctx) error {
	var response = []interface{}{"ok", true}

	return c.Status(http.StatusOK).JSON(response)
}

var settings = configs.GetSettings()

func RoutesV1(app *fiber.App) {
	// v1 := app.Group(settings["API_V1"])

}

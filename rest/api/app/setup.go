package app

import (
	// "jnj-api/api/middlewares"
	// routes "jnj-api/api/routes/api"
	"go-grpc-api/api/schemas/responses"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func MakeApp() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "There was an error with your request."

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				message = e.Message
			}
			log.Printf("error at process a request %s", err.Error())

			err = ctx.Status(code).JSON(responses.InfoResponse{Message: message})
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(responses.InfoResponse{Message: message})
			}

			return nil
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	routes.RoutesV1(app)

	return app
}

package server

import (
	"demo-server/handlers"
	"demo-server/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func Server() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE",
	}))
	app.Get("/heath", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "200",
			"message": "Demo service is working fine",
		})
	})
	var Log *log.Logger
	service, err := services.NewProductRequest()
	if err != nil {
		log.Println("error while starting the product services: ", err)
		return
	}
	handlers := handlers.NewHandler(Log).ProductTypeHandler(service)
	Routes(app, handlers)
	err = app.Listen(":8001")
	if err != nil {
		log.Println("error while starting the server: ", err)
	}
}

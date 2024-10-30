package server

import (
	"demo-server/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	// Starting service
	app.Post("/productType", h.NewProductTypeHandler)
}

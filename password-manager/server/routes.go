package server

import (
	"password-manager/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	// Master Service Api's
	app.Post("/editKey", h.EditKeyRequest)
	app.Get("/getInfo", h.GetInfoRequest)
	app.Patch("/algorithm", h.UpdateAlgorithmRequest)
	app.Post("/addUser", h.CreateUserRequest)
	app.Get("/listUsers", h.ListUserRequest)
	//Users Services Api's
	app.Post("/addwebiste", h.CreatePasswordRequest)
	app.Post("/getPasswrod", h.GetPasswordRequest)
	app.Get("/listWebiste", h.GetWebsiteRequest)
}

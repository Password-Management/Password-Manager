package server

import (
	"password-manager/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App, h *handlers.Handler) {
	// Login Service
	login := app.Group("/login")
	login.Post("/user", h.LoginUser)
	login.Post("/master", h.LoginMaster)
	login.Put("/logout", h.Logout)
	//Admin API's
	admin := app.Group("/admin", AdminMiddleware)
	admin.Post("/create", h.CreateMaster)
	// Master Service Api's
	master := app.Group("/master", RequireMasterID)
	master.Post("/editKey", h.EditKeyRequest)
	master.Get("/getInfo", h.GetInfoRequest)
	master.Patch("/algorithm", h.UpdateAlgorithmRequest)
	master.Post("/addUser", h.CreateUserRequest)
	master.Get("/listUsers", h.ListUserRequest)
	//Users Services Api's
	user := app.Group("/user", RequireUserIDAndMasterID)
	user.Post("/addwebiste", h.CreatePasswordRequest)
	user.Post("/getPasswrod", h.GetPasswordRequest)
	user.Get("/listWebiste", h.GetWebsiteRequest)
	user.Get("/getInfo", h.GetUserInfo)
	user.Delete("/password", h.DeleteWebsiteEntry)
}

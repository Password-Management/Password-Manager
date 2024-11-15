package server

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func RequireMasterID(c *fiber.Ctx) error {
	log.Println("Got inside the middleware")
	masterID := c.Get("master-id")
	if masterID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "master-id header is required",
		})
	}
	return c.Next()
}

func RequireUserIDAndMasterID(c *fiber.Ctx) error {
	userID := c.Get("user-id")
	masterID := c.Get("master-id")
	if userID == "" || masterID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user-id and master-id headers are required",
		})
	}
	return c.Next()
}

func AdminMiddleware(c *fiber.Ctx) error {
	if c.Get("Admin-id") == "" {
		return c.Status(http.StatusUnauthorized).SendString("Unauthorized as the admin Id is required for this operation")
	}
	return c.Next()
}

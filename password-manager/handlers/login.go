package handlers

import (
	"errors"
	"log"
	"password-manager/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) LoginMaster(c *fiber.Ctx) error {
	var requestBody *models.MasterLoginRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.LoginService.LoginMaster(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) LoginUser(c *fiber.Ctx) error {
	var requestBody *models.UserLoginRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body" + err.Error())
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.LoginService.LoginUser(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) Logout(c *fiber.Ctx) error {
	userId := c.Query("id")
	if userId == "" {
		log.Println("Query parameter of userid is empty")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("Query parameter of userid is empty"))
	}
	resp, err := h.LoginService.Logout(uuid.MustParse(userId))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}
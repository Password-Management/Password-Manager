package handlers

import (
	"errors"
	"log"
	"password-manager/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreatePasswordRequest(c *fiber.Ctx) error {
	var requestBody *models.CreatePasswordRequest
	userid := c.Get("user-id")
	masterid := c.Get("master-id")
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.UserService.CreateWebsiteEntry(requestBody, userid, masterid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) GetPasswordRequest(c *fiber.Ctx) error {
	var requestBody *models.GetPasswordRequest
	userid := c.Get("user-id")
	masterid := c.Get("master-id")
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.UserService.GetPassword(requestBody, userid, masterid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) GetWebsiteRequest(c *fiber.Ctx) error {
	var requestBody *models.ListWebsiteRequest
	userid := c.Get("user-id")
	masterid := c.Get("master-id")
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.UserService.ListWebsites(userid, masterid)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) GetUserInfo(c *fiber.Ctx) error {
	userId := c.Get("user-id")
	masterId := c.Get("master-id")
	resp, err := h.UserService.GetUserInfo(userId, masterId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) DeleteWebsiteEntry(c *fiber.Ctx) error {
	userId := c.Get("user-id")
	masterId := c.Get("master-id")
	websiteName := c.Query("webisteName")
	resp, err := h.UserService.DeletePassword(websiteName, masterId, userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}
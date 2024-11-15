package handlers

import (
	"errors"
	"log"
	"password-manager/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) EditKeyRequest(c *fiber.Ctx) error {
	var requestBody *models.EditKeyRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.MasterService.EditKey(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) GetInfoRequest(c *fiber.Ctx) error {
	specialKey := c.Query("specialKey")
	if specialKey == "" {
		log.Println("specialKey query parameter is missing")
		return c.Status(fiber.StatusBadRequest).JSON(errors.New("specialKey is required"))
	}
	resp, err := h.MasterService.GetInfo(specialKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) UpdateAlgorithmRequest(c *fiber.Ctx) error {
	var requestBody *models.UpdateAlgorithmRequest
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.MasterService.UpdateAlgorithm(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) CreateUserRequest(c *fiber.Ctx) error {
	var requestBody *models.CreateUserRequest
	masterId := c.Get("master-id")
	log.Println("the requestBody: ", requestBody)
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	resp, err := h.MasterService.CreateUser(requestBody, masterId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) ListUserRequest(c *fiber.Ctx) error {
	specialKey := c.Query("specialKey")
	if specialKey == "" {
		log.Println("specialKey query parameter is missing")
		return c.Status(fiber.StatusBadRequest).JSON(errors.New("specialKey is required"))
	}

	// Call the ListUser function with specialKey only
	resp, err := h.MasterService.ListUser(specialKey)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}

	// Return the success response
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

package handlers

import (
	"errors"
	"log"
	"os"
	"password-manager/helpers"
	"password-manager/models"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) CreateMaster(c *fiber.Ctx) error {
	err := helpers.Getenv()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors.New("error from the admin handler while loading the env file:" + err.Error()))
	}
	adminToken := os.Getenv("ADMIN_UUID")
	adminId := c.Get("Admin-id")
	log.Println("the adminToken = ", adminToken)
	log.Println("the header =", adminId)
	if adminId != adminToken {
		return c.Status(fiber.StatusBadRequest).JSON(models.Failed(&fiber.Error{
			Code:    400,
			Message: "Admin Id doesnt match the correct token",
		}))
	}
	log.Println("Before the service")
	err = h.AdminService.Create()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success("The creation of user is success."))
}

func (h *Handler) CreateOTP(c *fiber.Ctx) error {
	userID := c.Get("user-id")
	if userID == "" {
		return &fiber.Error{
			Code:    404,
			Message: "userId not present in the header",
		}
	}
	resp, err := h.AdminService.CreateOTP(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) VerifyOTP(c *fiber.Ctx) error {
	userID := c.Get("user-id")
	otp := c.Query("otp")
	if userID == "" {
		return &fiber.Error{
			Code:    404,
			Message: "userId not present in the header",
		}
	}
	if otp == "" {
		return &fiber.Error{
			Code:    404,
			Message: "otp not present in the query parameter",
		}
	}
	resp, err := h.AdminService.VerifyOTP(userID, otp)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

func (h *Handler) GetPlanInformation(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return &fiber.Error{
			Code:    404,
			Message: "customerId is not present in the query parameter",
		}
	}
	resp, err := h.AdminService.GetPlanInformation(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

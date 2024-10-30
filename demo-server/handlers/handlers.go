package handlers

import (
	"demo-server/models"
	"demo-server/services"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Logger             *log.Logger
	ProductTypeService services.Product
}

func NewHandler(logger *log.Logger) *Handler {
	return &Handler{Logger: logger}
}

func (h *Handler) ProductTypeHandler(pt services.Product) *Handler {
	h.ProductTypeService = pt
	return h
}

func (h *Handler) NewProductTypeHandler(c *fiber.Ctx) error {
	var requestBody *models.ProductDetailRequestBody
	err := c.BodyParser(&requestBody)
	if err != nil {
		log.Println("Error in parsing the request Body")
		return c.Status(fiber.StatusBadGateway).JSON(errors.New("error while parsing the request Body"))
	}
	if requestBody.Email == "" || requestBody.Name == "" || requestBody.ProductType == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    404,
			Message: "Email, Name or Product Type cannot be empty",
		}))
	}
	resp, err := h.ProductTypeService.GetProductDetails(requestBody)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Failed(&fiber.Error{
			Code:    500,
			Message: err.Error(),
		}))
	}
	return c.Status(fiber.StatusOK).JSON(models.Success(resp))
}

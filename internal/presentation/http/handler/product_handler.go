package handler

import (
	"login_module/internal/application/service"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	product := h.productService.GetProduct(id)
	return c.SendString(product)
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	// 제품 생성 로직
	return c.SendStatus(fiber.StatusCreated)
}

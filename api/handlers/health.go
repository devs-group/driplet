package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HealthHandler struct{}

func NewHealthHandler() (*HealthHandler, error) {
	return &HealthHandler{}, nil
}

func (h *HealthHandler) GET_health(c *fiber.Ctx) error {
	return c.SendString("OK")
}

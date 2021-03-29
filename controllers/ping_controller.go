package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Ping(c *fiber.Ctx) error {
	c.Status(http.StatusOK).SendString("ping!!! x")
	return nil
}

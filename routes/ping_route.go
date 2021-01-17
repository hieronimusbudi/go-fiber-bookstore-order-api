package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/controllers"
)

func PingRoutes(app *fiber.App) {
	app.Get("/api/orders/ping", controllers.Ping)
}

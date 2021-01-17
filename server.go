package main

import (
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/datasources/mongodb"
	events "github.com/hieronimusbudi/go-fiber-bookstore-order-api/events/consumer"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.OrderRoutes(app)
	routes.PingRoutes(app)

	go events.ConsumeItemCreatedEvent()

	mongodb.InitMongo()
	app.Listen(":9010")
}

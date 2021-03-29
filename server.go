package main

import (
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/datasources/mongodb"
	events "github.com/hieronimusbudi/go-fiber-bookstore-order-api/events/consumer"
	"github.com/hieronimusbudi/go-fiber-bookstore-order-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.PingRoutes(app)
	routes.OrderRoutes(app)

	go events.ConsumeItemCreatedEvent()
	go events.ConsumeItemUpdatedEvent()
	go events.ConsumeItemDeletedEvent()

	mongodb.InitMongo()
	app.Listen(":9010")
}

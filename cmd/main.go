package main

import (
	"github.com/NIC-Golang/API_gateway_GO_Alimka/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	app.Listen(":3000")
}

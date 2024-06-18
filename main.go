package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yasaswiniyerraballi/studentdatabase-management/routes"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
	// Start the server
	app.Listen(":8080") // Change port if needed
}

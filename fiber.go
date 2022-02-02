package main

import (
	"connector-api-cdc/database"
	"connector-api-cdc/handler"
	"connector-api-cdc/repository"
	"connector-api-cdc/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

func Run(port int) {

	gormdb := database.InitDb()

	app := fiber.New()
	api := app.Group("/api")

	WireConnectors(gormdb, api)

	// Prepare an endpoint for 'Not Found'.
	app.All("*", func(c *fiber.Ctx) error {
		errorMessage := fmt.Sprintf("Route '%s' does not exist in this API!", c.OriginalURL())

		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"status":  "fail",
			"message": errorMessage,
		})
	})

	// Listen to port 3000.
	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}

func WireConnectors(gormdb *gorm.DB, api fiber.Router) {
	// Create repositories.
	connectorRepository := repository.NewConnectorRepository(gormdb)

	// Create all of our services.
	connectorService := service.NewConnectorService(connectorRepository)

	// Prepare our endpoints for the API.
	handler.NewConnectorHandler(api.Group("/v1/connectors"), connectorService)
}

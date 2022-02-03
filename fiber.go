package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelorencini/connector-api-cdc/database"
	"github.com/rafaelorencini/connector-api-cdc/handler"
	"github.com/rafaelorencini/connector-api-cdc/kafka_connect"
	"github.com/rafaelorencini/connector-api-cdc/repository"
	"github.com/rafaelorencini/connector-api-cdc/service"
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

	read_yaml := service.NewReadYamlService()
	get_secrets := service.NewGetSecretsServiceService()

	buil_request := service.NewBuildRequestService(read_yaml, get_secrets)
	kafka_connect_service := kafka_connect.NewKafkaConnect()

	// Create repositories.
	connectorRepository := repository.NewConnectorRepository(gormdb)

	// Create all of our services.
	connectorService := service.NewConnectorService(connectorRepository, buil_request, kafka_connect_service)

	// Prepare our endpoints for the API.
	handler.NewConnectorHandler(api.Group("/v1/connectors"), connectorService)
}

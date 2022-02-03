package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelorencini/connector-api-cdc/database"
	"github.com/rafaelorencini/connector-api-cdc/domain"
	"github.com/rafaelorencini/connector-api-cdc/handler"
	"github.com/rafaelorencini/connector-api-cdc/kafka_connect"
	"github.com/rafaelorencini/connector-api-cdc/repository"
	"github.com/rafaelorencini/connector-api-cdc/service"
	request_builders "github.com/rafaelorencini/connector-api-cdc/service/request-builders"
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

	mapFunctions := map[string]domain.RequestBuilder{
		"mongo_sink":   request_builders.NewMongoSinkBuilder(),
		"mongo_source": request_builders.NewMongoSourceBuilder(),
		"mysql_sink":   request_builders.NewMysqlSinkBuilder(),
		"mysql_source": request_builders.NewMysqlSourceBuilder(),
	}

	readYaml := service.NewReadYamlService()
	getSecrets := service.NewGetSecretsServiceService()

	builRequest := service.NewBuildRequestService(readYaml, getSecrets, mapFunctions)
	kafkaConnectService := kafka_connect.NewKafkaConnect()

	// Create repositories.
	connectorRepository := repository.NewConnectorRepository(gormdb)

	// Create all of our services.
	connectorService := service.NewConnectorService(connectorRepository, builRequest, kafkaConnectService)

	// Prepare our endpoints for the API.
	handler.NewConnectorHandler(api.Group("/v1/connectors"), connectorService)
}

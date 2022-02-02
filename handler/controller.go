package handler

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/rafaelorencini/connector-api-cdc/domain"
	"gorm.io/gorm"
	"net/http"
)

func NewConnectorHandler(app fiber.Router, service domain.ConnectorServices) {
	app.Post("/", CreateConnector(service))
	app.Get("/", GetConnectors(service))
	app.Get("/:ID", GetConnector(service))
	app.Put("/:ID", UpdateConnector(service))
	app.Delete("/:ID", DeleteConnector(service))
}

func DeleteConnector(service domain.ConnectorServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var connector domain.Connector
		id := c.Params("ID")
		err := service.DeleteConnector(&connector, id)
		if err != nil {
			return err
		}
		return c.Status(200).JSON(&fiber.Map{
			"status": "delete connector with success",
			"error":  nil,
		})
	}

}

func UpdateConnector(service domain.ConnectorServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var connector domain.Connector
		id := c.Params("ID")
		err := service.GetConnector(&connector, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusNotFound).SendString(err.Error())
				return err
			}
			c.Status(http.StatusInternalServerError).SendString(err.Error())
			return err
		}
		c.BodyParser(&connector)
		err = service.UpdateConnector(&connector)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString(err.Error())
			return err
		}
		return c.Status(200).JSON(connector)
	}
}

func GetConnector(service domain.ConnectorServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("ID")
		var connector domain.Connector
		err := service.GetConnector(&connector, id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Status(http.StatusNotFound).SendString(err.Error())
				return err
			}
			c.Status(http.StatusInternalServerError).SendString(err.Error())
			return err
		}
		return c.Status(200).JSON(connector)
	}
}

func GetConnectors(service domain.ConnectorServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var connector []domain.Connector
		err := service.GetConnectors(&connector)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString(err.Error())
			return err
		}
		return c.Status(200).JSON(connector)
	}
}

func CreateConnector(service domain.ConnectorServices) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var connector domain.Connector
		c.BodyParser(&connector)
		err := service.CreateConnector(&connector)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString(err.Error())
			return err
		}
		return c.Status(201).JSON(connector)
	}
}

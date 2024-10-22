package server

import (
	"errors"

	"github.com/assaidy/bookstore/internals/database"
	"github.com/assaidy/bookstore/internals/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberServer struct {
	*fiber.App
    db *database.DBService
}

func NewFiberServer() *FiberServer {
	fs := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "bookstore",
			AppName:      "bookstore",
			ErrorHandler: errorHandler,
		}),
		db: database.NewDBService(),
	}
	fs.Use(logger.New())
	return fs
}

func errorHandler(c *fiber.Ctx, err error) error {
	var apiE utils.ApiError
	if errors.As(err, &apiE) {
		c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
		return c.Status(apiE.Code).JSON(apiE)
	}
	code := fiber.StatusInternalServerError
	var fiberE *fiber.Error
	if errors.As(err, &fiberE) {
		code = fiberE.Code
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return c.Status(code).SendString(err.Error())
}

package middleware

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func LoggerMiddleware(app *fiber.App) {
    app.Use(logger.New())
}

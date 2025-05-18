package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/violasptntels/WorkWise_Backend/config"
    "github.com/violasptntels/WorkWise_Backend/middleware"
    "github.com/violasptntels/WorkWise_Backend/routes"
)

func main() {
    app := fiber.New()

    config.ConnectDB()
    middleware.CorsMiddleware(app)
    middleware.LoggerMiddleware(app)

    routes.KaryawanRoutes(app)
    routes.TugasRoutes(app)

    app.Listen(":3000")
}
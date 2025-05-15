package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/violasptntels/WorkWise-Backend/config"
    "github.com/violasptntels/WorkWise-Backend/middleware"
    "github.com/violasptntels/WorkWise-Backend/routes"
)

func main() {
    app := fiber.New()

    config.ConnectDB()
    middleware.CorsMiddleware(app)
    middleware.LoggerMiddleware(app)

    routes.KaryawanRoutes(app)
    routes.TugasRoutes(app)
    routes.StatusRoutes(app)
    routes.DeadlineRoutes(app)

    app.Listen(":3000")
}
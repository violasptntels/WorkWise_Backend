package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/violasptntels/WorkWise_Backend/controllers"
)

func KaryawanRoutes(app *fiber.App) {
    route := app.Group("/karyawan")
    route.Get("/", controllers.GetAllKaryawan)
    route.Post("/", controllers.CreateKaryawan)
    route.Get("/:id", controllers.GetKaryawanByID)
    route.Put("/:id", controllers.UpdateKaryawan)
    route.Delete("/:id", controllers.DeleteKaryawan)
}

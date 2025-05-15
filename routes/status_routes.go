package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise_Backend/controllers"
)

func StatusRoutes(app *fiber.App) {
	status := app.Group("/status")
	status.Get("/", controllers.GetAllStatus)
	status.Get("/:id", controllers.GetStatusByID)
	status.Post("/", controllers.CreateStatus)
	status.Put("/:id", controllers.UpdateStatus)
	status.Delete("/:id", controllers.DeleteStatus)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise_Backend/controllers"
)

func DeadlineRoutes(app *fiber.App) {
	deadline := app.Group("/deadline")
	deadline.Get("/", controllers.GetAllDeadline)
	deadline.Get("/:id", controllers.GetDeadlineByID)
	deadline.Post("/", controllers.CreateDeadline)
	deadline.Put("/:id", controllers.UpdateDeadline)
	deadline.Delete("/:id", controllers.DeleteDeadline)
}

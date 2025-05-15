package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/violasptntels/WorkWise_Backend/controllers"
)

func TugasRoutes(app *fiber.App) {
	tugas := app.Group("/tugas")
	tugas.Get("/", controllers.GetAllTugas)
	tugas.Get("/:id", controllers.GetTugasByID)
	tugas.Post("/", controllers.CreateTugas)
	tugas.Put("/:id", controllers.UpdateTugas)
	tugas.Delete("/:id", controllers.DeleteTugas)
}

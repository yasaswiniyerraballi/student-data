package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yasaswiniyerraballi/studentdatabase-management/handlers"
)

func Setup(app *fiber.App) {
	// Student CRUD operations
	app.Get("/students", handlers.GetStudents)
	app.Post("/students", handlers.CreateStudent)
	app.Get("/students/:id", handlers.GetStudentByID)
	app.Put("/students/:id", handlers.UpdateStudent)
	app.Delete("/students/:id", handlers.DeleteStudent)
}

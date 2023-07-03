package routes

import (
	"github.com/aceino/belajar_restapi/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(route *fiber.App) {

	route.Get("/book", controllers.GetAllBook)
	route.Get("/book/:id", controllers.GetBookById)
	route.Post("/book", controllers.InsertBook)
	route.Put("/book/:id", controllers.UpdateBook)
	route.Delete("/book/:id", controllers.DeleteBook)
}

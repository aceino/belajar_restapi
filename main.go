package main

import (
	"github.com/aceino/belajar_restapi/config"
	"github.com/aceino/belajar_restapi/routes"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors")

func main() {

	app := fiber.New()

	// router
	routes.Routes(app)

	// middleware
	app.Use(cors.New())

	// database
	config.ConnectionDB()

	app.Listen(":3000")
}

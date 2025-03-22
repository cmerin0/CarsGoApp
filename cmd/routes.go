package main

import (
	"github.com/cmerin0/CarsGoApp/handlers"
	"github.com/gofiber/fiber/v2"
)

// Function that gathers the routes in a single file
func setupRoutes(app *fiber.App) {

	// Main route of test
	app.Get("/", handlers.Home)

	// Routes of Make
	app.Get("/makes", handlers.GetMakes)
	app.Get("/makes/:id", handlers.GetMakeById)
	app.Post("/makes", handlers.CreateMake)
	app.Put("/makes/:id", handlers.UpdateMake)
	app.Delete("/makes/:id", handlers.DeleteMake)

	// Routes of Car
	app.Get("/cars", handlers.GetCars)
	app.Post("/cars", handlers.CreateCar)
}

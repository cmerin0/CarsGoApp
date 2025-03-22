package handlers

import (
	"github.com/cmerin0/CarsGoApp/db"
	"github.com/cmerin0/CarsGoApp/models"
	"github.com/gofiber/fiber/v2"
)

func GetCars(c *fiber.Ctx) error {
	cars := []models.Car{} // Creates an array of Models cars

	db.DB.Db.Find(&cars) // Find all the cars and store them in the cars variable

	return c.Status(200).JSON(cars) // Return a status 200 code response and the cars
}

func CreateCar(c *fiber.Ctx) error {
	car := new(models.Car) // Create a new car

	// Check if body sent is properly sent
	if err := c.BodyParser(car); err != nil { // Parse the body of the request and store it in the car variable
		return c.Status(400).JSON(err.Error()) // Return a status 400 code response and the error
	}

	db.DB.Db.Create(&car) // Create the car in the database

	return c.Status(201).JSON(car) // Return a status 201 code response and the car
}

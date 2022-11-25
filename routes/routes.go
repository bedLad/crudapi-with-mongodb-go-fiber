package routes

import (
	"log"

	"github.com/bedLad/go-fiber-mongo-hrms/database"
	"github.com/bedLad/go-fiber-mongo-hrms/models"
	"github.com/gofiber/fiber/v2"
)

func Init() {
	app := fiber.New()
	app.Get("/employee", getEmployees)
	app.Get("/employee/:id", getEmployeeById)
	app.Post("/employee", createEmployee)
	app.Put("/employee/:id", updateEmployee)
	app.Delete("/employee/:id", deleteEmployee)

	log.Fatal(app.Listen(":3000"))
}

func getEmployees(c *fiber.Ctx) error {
	employees := database.GetCollections()
	return c.JSON(employees)
}

func createEmployee(c *fiber.Ctx) error {
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	database.CreateCollection(employee)
	return c.JSON(employee)
}

func getEmployeeById(c *fiber.Ctx) error {
	id := c.Params("id")
	employee := database.GetCollectionByID(id)
	return c.Status(201).JSON(employee)
}

func deleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	database.DeleteCollection(id)
	return nil
}

func updateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee
	if err := c.BodyParser(&employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	database.UpdateCollection(id, employee)
	return nil
}

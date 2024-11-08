package routes

import (
	"onez19/middlewares"
	"onez19/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/register", services.Register)
	app.Post("/login", services.Login)

	app.Use(middlewares.AuthRequired)
	app.Get("/users", services.GetUsers)
	app.Get("/contracts/:username", services.GetAllContractsByUsername)
	app.Get("/contracts/:contract_number/:contract_year", services.GetContractDetails)
	app.Get("/bills/:tenant_username", services.GetBillsByTenantUsername)
	app.Get("/bills/:bill_id/items", services.GetBillItemsByBillID)

}

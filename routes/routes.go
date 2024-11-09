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
	app.Get("/users/tenant", services.GetUsersWithTenant)
	app.Get("/contracts", services.GetAllContracts)
	app.Get("/contracts/:username", services.GetAllContractsByUsername)
	app.Get("/contracts/:contract_number/:contract_year", services.GetContractDetails)
	app.Put("/contracts/:contract_room_number/:status", services.UpdateContractStatus)
	app.Get("/contracts/check-active/:contract_room_number", services.CheckContractActive)
	app.Post("/contracts/create", services.CreateContract)
	app.Get("/bills", services.GetAllBills)
	app.Get("/bills/:tenant_username", services.GetBillsByTenantUsername)
	app.Get("/bills/:bill_id/items", services.GetBillItemsByBillID)
	app.Get("/transactions/:bill_id", services.GetTransactionsByBillID)
	app.Post("/transactions", services.CreateTransaction)
	app.Put("/transactions/:transaction_id/status/:status", services.UpdateTransactionStatus)
	app.Put("/bills/:bill_id/status/:status", services.UpdateBillStatus)
	// Route สำหรับการสร้างการจอง
	app.Post("/reservations/create", services.CreateReservation)

	// Route สำหรับดึงข้อมูลการจองทั้งหมด
	app.Get("/reservations", services.GetReservations)

	// Route สำหรับดึงข้อมูลการจองตาม tenant_username
	app.Get("/reservations/tenant/:tenant_username", services.GetReservationsByUsername)

	// Route สำหรับการอัปเดตสถานะของการจอง
	app.Put("/reservations/:reservation_id/status", services.UpdateReservationStatus)

	app.Get("/reservations/:reservation_id", services.GetReservationByID) // เพิ่มเส้นทางนี้
}

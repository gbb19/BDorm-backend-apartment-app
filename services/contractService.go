package services

import (
	"log"
	"onez19/datasources"

	"github.com/gofiber/fiber/v2"
)

func GetAllContractsByUsername(c *fiber.Ctx) error {
	username := c.Params("username") // รับค่า username จาก path parameter

	// ดึงข้อมูลหมายเลขห้องที่มีสัญญา active
	contracts, err := datasources.GetAllContractsByUsername(username)
	if err != nil {
		log.Println("Error fetching contracts:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch contracts"})
	}

	// หากไม่มีสัญญา active
	if contracts == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No active contracts found"})
	}

	// ส่งข้อมูลสัญญากลับไปยัง client
	return c.JSON(fiber.Map{
		"contracts": contracts,
	})
}

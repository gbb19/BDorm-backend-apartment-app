package services

import (
	"log"
	"onez19/datasources"
	"strconv"

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

func GetContractDetails(c *fiber.Ctx) error {
	// รับค่า contract_number และ contract_year จาก path parameter
	contractNumber, err := strconv.Atoi(c.Params("contract_number"))
	if err != nil {
		log.Println("Invalid contract number:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contract number"})
	}

	contractYear, err := strconv.Atoi(c.Params("contract_year"))
	if err != nil {
		log.Println("Invalid contract year:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid contract year"})
	}

	// ดึงรายละเอียดสัญญาจากฐานข้อมูล
	contract, err := datasources.GetContractDetails(contractNumber, contractYear)
	if err != nil {
		log.Println("Error fetching contract details:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch contract details"})
	}

	// หากไม่พบสัญญา
	if contract == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contract not found"})
	}

	// ส่งรายละเอียดสัญญากลับไปยัง client
	return c.JSON(fiber.Map{
		"contract": contract,
	})
}

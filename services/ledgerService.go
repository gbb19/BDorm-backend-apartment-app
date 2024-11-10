package services

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"onez19/datasources"
	"onez19/entities"
	"strconv"
)

func CreateLedger(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request entities.Ledger

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่อแทรกข้อมูลในตาราง ledger
	err := datasources.CreateLedger(request)
	if err != nil {
		log.Println("Error creating ledger:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create ledger"})
	}

	// ส่งการตอบกลับว่าเพิ่มข้อมูลสำเร็จ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ledger created successfully",
	})
}

func CreateLedgerItem(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request entities.LedgerItem

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่อแทรกข้อมูลในตาราง ledger_item
	err := datasources.CreateLedgerItem(request)
	if err != nil {
		log.Println("Error creating ledger_item:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create ledger_item"})
	}

	// ส่งการตอบกลับว่าเพิ่มข้อมูลสำเร็จ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ledger item created successfully",
	})
}

func UpdateLedgerItem(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request entities.LedgerItemUpdate

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่ออัปเดตข้อมูลในตาราง ledger_item
	err := datasources.UpdateLedgerItem(request.WaterUnit, request.ElectricityUnit, request.LedgerMonth, request.LedgerYear, request.LedgerItemRoomNumber)
	if err != nil {
		log.Println("Error updating ledger_item:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update ledger_item"})
	}

	// ส่งการตอบกลับว่าอัปเดตสำเร็จ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ledger item updated successfully",
	})
}

func UpdateLedgerItemStatus(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request struct {
		LedgerMonth          int `json:"ledger_month"`
		LedgerYear           int `json:"ledger_year"`
		LedgerItemRoomNumber int `json:"ledger_item_room_number"`
	}

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่ออัปเดตข้อมูลในตาราง ledger_item
	err := datasources.UpdateLedgerItemStatus(request.LedgerMonth, request.LedgerYear, request.LedgerItemRoomNumber)
	if err != nil {
		log.Println("Error updating ledger_item_status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update ledger_item_status"})
	}

	// ส่งการตอบกลับว่าอัปเดตสำเร็จ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Ledger item status updated successfully",
	})
}

func GetLedgerItemsByMonthAndYear(c *fiber.Ctx) error {
	month, err := strconv.Atoi(c.Params("month"))
	if err != nil {
		log.Println("Invalid month:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid month"})
	}

	year, err := strconv.Atoi(c.Params("year"))
	if err != nil {
		log.Println("Invalid year:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid year"})
	}

	ledgerItems, err := datasources.GetLedgerItemsByMonthAndYear(month, year)
	if err != nil {
		log.Println("Error fetching ledger items:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch ledger items"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"ledger_items": ledgerItems, // จะแสดง IsActive สำหรับแต่ละห้อง
	})
}

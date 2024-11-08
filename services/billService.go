package services

import (
	"log"
	"onez19/datasources"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBillsByTenantUsername(c *fiber.Ctx) error {
	// รับค่า tenant_username จาก path parameter
	tenantUsername := c.Params("tenant_username")

	// ดึงข้อมูลใบแจ้งหนี้จากฐานข้อมูล
	bills, err := datasources.GetBillsByTenantUsername(tenantUsername)
	if err != nil {
		log.Println("Error fetching bills:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch bills"})
	}

	// หากไม่พบใบแจ้งหนี้
	if len(bills) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No bills found"})
	}

	// ส่งข้อมูลใบแจ้งหนี้กลับไปยัง client
	return c.JSON(fiber.Map{
		"bills": bills,
	})
}

func GetBillItemsByBillID(c *fiber.Ctx) error {
	// รับค่า bill_id จาก path parameter และแปลงเป็น integer
	billID, err := strconv.Atoi(c.Params("bill_id"))
	if err != nil {
		log.Println("Invalid bill_id:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bill ID"})
	}

	// ดึงข้อมูล bill items จากฐานข้อมูล
	billItems, err := datasources.GetBillItemsByBillID(billID)
	if err != nil {
		log.Println("Error fetching bill items:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch bill items"})
	}

	// กรณีไม่มีข้อมูล bill items ให้ส่ง response แบบ 200 OK และ array ว่าง
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bill_items": billItems, // จะเป็น [] ถ้าไม่มีข้อมูล
	})
}

// ฟังก์ชันเพื่อดึงข้อมูล transaction ตาม bill_id
func GetTransactionsByBillID(c *fiber.Ctx) error {
	// รับค่า bill_id จาก path parameter และแปลงเป็น integer
	billID, err := strconv.Atoi(c.Params("bill_id"))
	if err != nil {
		log.Println("Invalid bill_id:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid bill ID"})
	}

	// ดึงข้อมูล transaction จากฐานข้อมูล
	transactions, err := datasources.GetTransactionsByBillID(billID)
	if err != nil {
		log.Println("Error fetching transactions:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch transactions"})
	}

	// กรณีไม่มีข้อมูล transaction ให้ส่ง response แบบ 200 OK และ array ว่าง
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transactions": transactions, // จะเป็น [] ถ้าไม่มีข้อมูล
	})
}

func CreateTransaction(c *fiber.Ctx) error {
	// รับค่า bill_id จาก body ของ request
	var request struct {
		BillID int `json:"bill_id"`
	}

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่อสร้าง transaction ใหม่
	transactionID, err := datasources.InsertTransaction(request.BillID)
	if err != nil {
		log.Println("Error creating transaction:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction"})
	}

	// คืนค่า transaction_id ที่ถูกสร้างไปยัง client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"transaction_id": transactionID,
	})
}

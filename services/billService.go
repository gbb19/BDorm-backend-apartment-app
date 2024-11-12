package services

import (
	"log"
	"onez19/datasources"
	"onez19/entities"
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

	// ส่งข้อมูลใบแจ้งหนี้กลับไปยัง client
	// ถ้าไม่มีข้อมูล bills, จะส่ง array ว่าง
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bills": bills, // ถ้าไม่มีข้อมูลจะเป็น [] เลย
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

func GetAllBills(c *fiber.Ctx) error {
	// ดึงข้อมูลทั้งหมดของบิลจากฐานข้อมูล
	bills, err := datasources.GetAllBills()
	if err != nil {
		log.Println("Error fetching all bills:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch bills"})
	}

	// ส่งข้อมูลใบแจ้งหนี้กลับไปยัง client
	// ถ้าไม่มีข้อมูล bills, จะส่ง array ว่าง
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bills": bills, // ถ้าไม่มีข้อมูลจะเป็น [] เลย
	})
}

func UpdateTransactionStatus(c *fiber.Ctx) error {
	// รับ transaction_id, status, และ username จาก request params
	transactionIDParam := c.Params("transaction_id")
	statusParam := c.Params("status")
	username := c.Params("username") // ดึง username จาก params

	// แปลง transaction_id และ status เป็น int
	transactionID, err := strconv.Atoi(transactionIDParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid transaction ID"})
	}

	status, err := strconv.Atoi(statusParam)
	if err != nil || (status != 1 && status != 2) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid status value"})
	}

	// ตรวจสอบว่า username มีค่าหรือไม่
	if username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Username is required"})
	}

	// อัปเดตสถานะของ transaction และ username
	err = datasources.UpdateTransactionStatus(transactionID, status, username)
	if err != nil {
		log.Println("Error updating transaction status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update transaction status"})
	}

	// ส่ง response กลับ
	return c.JSON(fiber.Map{"message": "Transaction status updated successfully"})
}

func UpdateBillStatus(c *fiber.Ctx) error {
	// รับ bill_id และ status จาก path parameter
	billIDStr := c.Params("bill_id")
	status := c.Params("status")

	// แปลง billID จาก string เป็น int
	billID, err := strconv.Atoi(billIDStr)
	if err != nil {
		log.Println("Invalid bill_id:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid bill_id. Must be an integer",
		})
	}

	// ตรวจสอบค่า status ว่าต้องเป็น 1 (Paid) หรือ 2 (Verified) เท่านั้น
	if status != "1" && status != "2" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status. Allowed values are 1 (Paid) or 2 (Verified)",
		})
	}

	// เรียกฟังก์ชันใน datasource เพื่ออัปเดตสถานะของใบแจ้งหนี้
	err = datasources.UpdateBillStatus(billID, status)
	if err != nil {
		log.Println("Error updating bill status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update bill status"})
	}

	// ค้นหา reservation ที่เชื่อมโยงกับ bill_id นี้
	reservationID, err := datasources.GetReservationByBillID(billID)
	if err != nil {
		log.Println("Error retrieving reservation ID:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve reservation"})
	}

	// หากไม่มี reservation ที่ตรงกับ bill_id ก็ไม่ต้องทำการอัปเดต reservation status
	if reservationID == 0 {
		return c.JSON(fiber.Map{
			"message": "Bill status updated successfully. No associated reservation found.",
		})
	}

	// ตรวจสอบค่า status ของ bill เพื่ออัปเดต status ของ reservation
	var newReservationStatus int
	if status == "1" {
		newReservationStatus = 3 // ถ้า bill_status เป็น 1 ให้อัปเดต reservation_status เป็น 3
	} else if status == "2" {
		newReservationStatus = 4 // ถ้า bill_status เป็น 2 ให้อัปเดต reservation_status เป็น 4
	}

	// อัปเดต reservation status
	err = datasources.UpdateReservationStatus(reservationID, newReservationStatus)
	if err != nil {
		log.Println("Error updating reservation status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update reservation status"})
	}

	// ส่งการตอบกลับว่าอัปเดตสำเร็จ
	return c.JSON(fiber.Map{
		"message": "Bill and reservation statuses updated successfully",
	})
}

func CreateBill(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request entities.BillCreate

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่อแทรกข้อมูลในตาราง bill
	billID, err := datasources.CreateBill(request)
	if err != nil {
		log.Println("Error creating bill:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create bill"})
	}

	// คืนค่า billID ที่ถูกสร้างไปยัง client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"bill_id": billID,
	})
}

func CreateBillItem(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request entities.BillItem

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่อแทรกข้อมูลในตาราง bill_item
	err := datasources.CreateBillItem(request)
	if err != nil {
		log.Println("Error creating bill item:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create bill item"})
	}

	// คืนค่าคำตอบที่ถูกต้องไปยัง client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Bill item created successfully",
	})
}

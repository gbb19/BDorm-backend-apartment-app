package services

import (
	"log"
	"onez19/datasources"
	"onez19/entities"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateReservation(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body ที่มีโครงสร้างเป็น ReservationCreate
	var request entities.ReservationCreate

	// อ่านข้อมูลจาก request body
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// เรียก datasource เพื่อแทรก reservation ใหม่
	reservationID, err := datasources.InsertReservation(request.MoveInDateTime, request.ReservationRoomNumber, request.TenantUsername, request.ManagerUsername, request.BillID)
	if err != nil {
		log.Println("Error creating reservation:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create reservation"})
	}

	// คืนค่า reservation_id ที่ถูกสร้างไปยัง client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"reservation_id": reservationID,
	})
}

func GetReservationsByUsername(c *fiber.Ctx) error {
	// รับค่า tenant_username จาก path parameter
	tenantUsername := c.Params("tenant_username")

	// ดึงข้อมูล reservations จากฐานข้อมูล
	reservations, err := datasources.GetReservationsByUsername(tenantUsername)
	if err != nil {
		log.Println("Error fetching reservations by username:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch reservations"})
	}

	// ส่งข้อมูล reservations กลับไปยัง client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"reservations": reservations, // ถ้าไม่มีข้อมูลจะเป็น [] เลย
	})
}

func UpdateReservationStatus(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request entities.UpdateReservationStatus

	// แปลงข้อมูลจาก request body เป็น struct
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// ตรวจสอบค่าที่ได้หลังจากแปลง JSON
	log.Printf("Parsed request - ReservationID: %d, ReservationStatus: %d", request.ReservationID, request.ReservationStatus)

	// ตรวจสอบว่าฟิลด์ ReservationID ไม่เป็นค่า default (เช่น 0)
	if request.ReservationID == 0 {
		log.Println("Missing reservationID in request body")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing reservationID in request body"})
	}

	// เรียกใช้ datasource เพื่ออัปเดตสถานะของการจอง
	err := datasources.UpdateReservationStatus(request.ReservationID, request.ReservationStatus)
	if err != nil {
		log.Println("Error updating reservation status:", err)
		if err.Error() == "No rows updated" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update reservation status"})
	}

	// ส่งข้อความตอบกลับเมื่ออัปเดตสำเร็จ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reservation status updated successfully"})
}

func GetReservations(c *fiber.Ctx) error {
	// เรียก datasource เพื่อดึงข้อมูล reservations ทั้งหมด
	reservations, err := datasources.GetReservations()
	if err != nil {
		log.Println("Error fetching reservations:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch reservations"})
	}

	// ถ้าไม่มีข้อมูล reservations ให้คืนค่าเป็นอาร์เรย์ว่าง
	if len(reservations) == 0 {
		return c.JSON(fiber.Map{"reservations": []interface{}{}})
	}

	// ส่งข้อมูล reservations กลับไปยัง client
	return c.JSON(fiber.Map{
		"reservations": reservations,
	})
}

func GetReservationByID(c *fiber.Ctx) error {
	// รับค่า reservation_id จาก path parameter
	reservationIDStr := c.Params("reservation_id")
	reservationID, err := strconv.Atoi(reservationIDStr) // แปลงจาก string เป็น int
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid reservation ID"})
	}

	// ดึงข้อมูล reservation จากฐานข้อมูล
	reservation, err := datasources.GetReservationByID(reservationID)
	if err != nil {
		log.Println("Error fetching reservation by ID:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch reservation"})
	}

	// ถ้าไม่พบข้อมูล reservation
	if reservation == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Reservation not found"})
	}

	// ส่งข้อมูล reservation กลับไปยัง client
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"reservation": reservation,
	})
}

func UpdateReservationDetails(c *fiber.Ctx) error {
	// รับข้อมูลจาก request body
	var request entities.UpdateReservationDetails

	// แปลงข้อมูลจาก request body เป็น struct
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	log.Printf("Request received with reservationID: %d, billID: %d, managerUsername: %s", 
		request.ReservationID, request.BillID, request.ManagerUsername)

	// เรียกใช้ datasource เพื่ออัปเดตข้อมูลในฐานข้อมูล
	err := datasources.UpdateReservationDetails(request.ReservationID, request.BillID, request.ManagerUsername)
	if err != nil {
		log.Println("Error updating reservation details:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update reservation details"})
	}

	// ส่งข้อความตอบกลับเมื่ออัปเดตสำเร็จ
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Reservation details updated successfully"})
}

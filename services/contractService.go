package services

import (
	"log"
	"onez19/datasources"
	"onez19/entities"
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

	// ส่งข้อมูลสัญญากลับไปยัง client
	// ถ้าไม่มีข้อมูล contracts, จะส่ง array ว่าง
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"contracts": contracts, // ถ้าไม่มีข้อมูลจะเป็น [] เลย
	})
}

func GetAllContracts(c *fiber.Ctx) error {
	// ดึงข้อมูลทั้งหมดของสัญญา active
	contracts, err := datasources.GetAllContracts()
	if err != nil {
		log.Println("Error fetching contracts:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch contracts"})
	}

	// ส่งข้อมูลสัญญากลับไปยัง client
	// ถ้าไม่มีข้อมูล contracts, จะส่ง array ว่าง
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"contracts": contracts, // ถ้าไม่มีข้อมูลจะเป็น [] เลย
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

func UpdateContractStatus(c *fiber.Ctx) error {
	// รับ contract_room_number และ status จาก path parameters
	contractRoomNumberStr := c.Params("contract_room_number")
	statusStr := c.Params("status")

	// แปลง contract_room_number จาก string เป็น int
	contractRoomNumber, err := strconv.Atoi(contractRoomNumberStr)
	if err != nil {
		log.Println("Invalid contract room number:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid contract room number",
		})
	}

	// แปลง status จาก string เป็น int
	status, err := strconv.Atoi(statusStr)
	if err != nil || (status != 0 && status != 1) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status. Allowed values are 0 (Active) or 1 (Cancelled)",
		})
	}

	// เรียกฟังก์ชันใน datasource เพื่ออัปเดตสถานะของ contract
	err = datasources.UpdateContractStatus(contractRoomNumber, status)
	if err != nil {
		log.Println("Error updating contract status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update contract status",
		})
	}

	// ส่งการตอบกลับว่าอัปเดตสำเร็จ
	return c.JSON(fiber.Map{
		"message": "Contract status updated successfully",
	})
}

func CheckContractActive(c *fiber.Ctx) error {
	// รับ contract_room_number จาก path parameter
	contractRoomNumberStr := c.Params("contract_room_number")

	// แปลง contract_room_number จาก string เป็น int
	contractRoomNumber, err := strconv.Atoi(contractRoomNumberStr)
	if err != nil {
		log.Println("Error checking contract status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to check contract status",
		})
	}

	// เรียกฟังก์ชันใน datasource เพื่อเช็คสถานะของ contract
	isActive, err := datasources.CheckContractActive(contractRoomNumber)
	if err != nil {
		log.Println("Error checking contract status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to check contract status",
		})
	}

	// ส่งการตอบกลับว่า contract นี้ยัง active หรือไม่
	if isActive {
		return c.JSON(fiber.Map{
			"message": "Contract is active",
		})
	} else {
		return c.JSON(fiber.Map{
			"message": "Contract is not active",
		})
	}
}

func CreateContract(c *fiber.Ctx) error {
	// สร้างตัวแปร contractData จาก struct ContractCreate เพื่อรับข้อมูลจาก request body
	var contractData entities.ContractCreate

	// อ่านข้อมูลจาก body ของ request และจับข้อมูลเข้าไปใน contractData
	if err := c.BodyParser(&contractData); err != nil {
		log.Println("Error parsing contract data:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	// ตรวจสอบว่า contract ที่มี contract_room_number เดียวกันนั้น active หรือไม่
	isActive, err := datasources.CheckContractActive(contractData.ContractRoomNumber)
	if err != nil {
		log.Println("Error checking contract status:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to check contract status",
		})
	}

	// หาก contract อยู่ในสถานะ active, ทำการอัปเดต status เป็น 1 (Cancelled)
	if isActive {
		err := datasources.UpdateContractStatus(contractData.ContractRoomNumber, 1) // 1 = Cancelled
		if err != nil {
			log.Println("Error updating contract status:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to update contract status",
			})
		}
	}

	// สร้าง contract ใหม่
	err = datasources.CreateContract(
		contractData.ContractYear,
		contractData.ContractRoomNumber,
		contractData.RentalPrice,
		contractData.WaterRate,
		contractData.ElectricityRate,
		contractData.InternetServiceFee,
		contractData.Username,
	)

	if err != nil {
		log.Println("Error creating contract:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create contract",
		})
	}

	// ส่งการตอบกลับว่า contract ได้ถูกสร้างสำเร็จ
	return c.JSON(fiber.Map{
		"message": "Contract created successfully",
	})
}
